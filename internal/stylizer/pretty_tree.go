package stylizer

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"sort"

	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/operator-framework/deppy/pkg/deppy"
	"github.com/perdasilva/deconflictor/internal/typed_constraints"
	"github.com/perdasilva/deconflictor/internal/typed_constraints/constraints"
)

var _ ErrorStylizer = &PrettyTree{}

type nodeMap map[deppy.Identifier]*node

func (n nodeMap) GetOrNew(id deppy.Identifier) *node {
	if _, ok := n[id]; !ok {
		n[id] = &node{
			id:       id,
			refCount: 0,
		}
	}
	return n[id]
}

func (n nodeMap) Prune() {
	ids := make([]deppy.Identifier, 0, len(n))
	for id, _ := range n {
		ids = append(ids, id)
	}
	for _, id := range ids {
		node := n[id]
		if len(node.eigenConstraints) == 0 && len(node.edges) == 0 {
			delete(n, id)
		}
	}
}

type constraintEdge struct {
	toNode     *node
	constraint deppy.Constraint
}
type node struct {
	id               deppy.Identifier
	edges            []*constraintEdge
	eigenConstraints []deppy.Constraint
	refCount         int
}

type PrettyTree struct {
}

func (p *PrettyTree) Stylize(error deppy.NotSatisfiable) string {
	// group constraints by identifier
	nodeMap := buildNodeTree(error)

	var nodes []*node
	for _, node := range nodeMap {
		nodes = append(nodes, node)
	}
	sort.SliceStable(nodes, func(i, j int) bool {
		if nodes[i].refCount == nodes[j].refCount {
			return nodes[i].id < nodes[j].id
		}
		return nodes[i].refCount < nodes[j].refCount
	})
	l := buildTree(nodes)
	return l.Render()
}

func buildNodeTree(appliedConstraints []deppy.AppliedConstraint) nodeMap {
	nodeMap := nodeMap{}

	// build out node tree
	for _, ap := range appliedConstraints {
		variableID := ap.Variable.Identifier()
		n := nodeMap.GetOrNew(variableID)
		switch c := ap.Constraint.(type) {
		case typed_constraints.Mandatory:
			n.eigenConstraints = append(n.eigenConstraints, c)
		case typed_constraints.Prohibited:
			n.eigenConstraints = append(n.eigenConstraints, c)
		case typed_constraints.Conflict:
			conflictingID := deppy.Identifier(c)
			conflictingNode := nodeMap.GetOrNew(conflictingID)
			conflictingNode.refCount = conflictingNode.refCount + 1
			n.edges = append(n.edges, &constraintEdge{
				toNode:     conflictingNode,
				constraint: c,
			})
			conflictingNode.edges = append(conflictingNode.edges, &constraintEdge{
				toNode:     n,
				constraint: constraints.Conflict(n.id),
			})
		case typed_constraints.Dependency:
			for _, id := range c.Order() {
				toNode := nodeMap.GetOrNew(id)
				toNode.refCount = toNode.refCount + 1
				n.edges = append(n.edges, &constraintEdge{
					toNode:     toNode,
					constraint: c,
				})
			}
		case typed_constraints.AtMost:
			for _, id := range c.IDs() {
				toNode := nodeMap.GetOrNew(id)
				toNode.refCount = toNode.refCount + 1
				n.edges = append(n.edges, &constraintEdge{
					toNode:     toNode,
					constraint: c,
				})
			}
		}
	}

	// prune nodes without edges of eigenConstraints
	nodeMap.Prune()

	return nodeMap
}

func buildTree(nodes []*node) list.Writer {
	l := list.NewWriter()
	l.SetStyle(list.StyleConnectedRounded)
	visited := map[deppy.Identifier]struct{}{}

	for _, node := range nodes {
		if _, ok := visited[node.id]; !ok {
			renderTree(node, &visited, &l, true)
		}
	}
	return l
}

func renderTree(node *node, visited *map[deppy.Identifier]struct{}, l *list.Writer, suppressNodeID bool) {
	if _, ok := (*visited)[node.id]; ok {
		return
	}
	(*visited)[node.id] = struct{}{}

	if !suppressNodeID {
		(*l).AppendItem(node.id)
	}
	for _, c := range node.eigenConstraints {
		(*l).Indent()
		(*l).AppendItem(c.String(node.id))
		(*l).UnIndent()
	}

	visitedConstraints := map[string]struct{}{}
	for _, e := range node.edges {
		if _, ok := (*visited)[e.toNode.id]; ok {
			continue
		}

		if _, ok := visitedConstraints[hashConstraint(e.constraint)]; !ok {
			(*l).Indent()
			(*l).AppendItem(e.constraint.String(node.id))
			(*l).Indent()
			visitedConstraints[hashConstraint(e.constraint)] = struct{}{}
		}
		renderTree(e.toNode, visited, l, true)
		(*l).UnIndent()
		(*l).UnIndent()
	}
	return
}

func hashConstraint(c deppy.Constraint) string {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(c)
	return hex.EncodeToString(b.Bytes())
}
