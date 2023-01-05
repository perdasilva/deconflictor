package stylizer

import (
	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/operator-framework/deppy/pkg/deppy"
)

var _ ErrorStylizer = &PrettyList{}

type PrettyList struct {
}

func (p *PrettyList) Stylize(error deppy.NotSatisfiable) string {
	// group constraints by identifier
	type node struct {
		id                 deppy.Identifier
		inOrderConstraints []deppy.Constraint
	}
	nodeMap := map[deppy.Identifier]*node{}

	for _, appliedConstraint := range error {
		variableID := appliedConstraint.Variable.Identifier()
		if _, ok := nodeMap[variableID]; !ok {
			nodeMap[variableID] = &node{
				id: variableID,
			}
		}
		n := nodeMap[variableID]
		n.inOrderConstraints = append(n.inOrderConstraints, appliedConstraint.Constraint)
	}

	l := list.NewWriter()
	l.SetStyle(list.StyleConnectedRounded)
	for _, n := range nodeMap {
		l.AppendItem(n.id)
		for _, c := range n.inOrderConstraints {
			l.Indent()
			l.AppendItem(c.String(n.id))
			l.UnIndent()
		}
	}
	return l.Render()
}
