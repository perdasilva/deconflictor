package typed_constraints

import (
	"fmt"
	"strings"

	"github.com/go-air/gini/z"

	"github.com/operator-framework/deppy/pkg/deppy"
)

type Mandatory struct{}

func (constraint Mandatory) String(subject deppy.Identifier) string {
	return fmt.Sprintf("%s is Mandatory", subject)
}

func (constraint Mandatory) Apply(lm deppy.LitMapping, subject deppy.Identifier) z.Lit {
	return lm.LitOf(subject)
}

func (constraint Mandatory) Order() []deppy.Identifier {
	return nil
}

func (constraint Mandatory) Anchor() bool {
	return true
}

type Prohibited struct{}

func (constraint Prohibited) String(subject deppy.Identifier) string {
	return fmt.Sprintf("%s is Prohibited", subject)
}

func (constraint Prohibited) Apply(lm deppy.LitMapping, subject deppy.Identifier) z.Lit {
	return lm.LitOf(subject).Not()
}

func (constraint Prohibited) Order() []deppy.Identifier {
	return nil
}

func (constraint Prohibited) Anchor() bool {
	return false
}

type Dependency []deppy.Identifier

func (constraint Dependency) String(subject deppy.Identifier) string {
	if len(constraint) == 0 {
		return fmt.Sprintf("%s has a Dependency without any candidates to satisfy it", subject)
	}
	s := make([]string, len(constraint))
	for i, each := range constraint {
		s[i] = string(each)
	}
	return fmt.Sprintf("%s requires at least one of %s", subject, strings.Join(s, ", "))
}

func (constraint Dependency) Apply(lm deppy.LitMapping, subject deppy.Identifier) z.Lit {
	m := lm.LitOf(subject).Not()
	for _, each := range constraint {
		m = lm.LogicCircuit().Or(m, lm.LitOf(each))
	}
	return m
}

func (constraint Dependency) Order() []deppy.Identifier {
	return constraint
}

func (constraint Dependency) Anchor() bool {
	return false
}

type Conflict deppy.Identifier

func (constraint Conflict) String(subject deppy.Identifier) string {
	return fmt.Sprintf("%s conflicts with %s", subject, constraint)
}

func (constraint Conflict) Apply(lm deppy.LitMapping, subject deppy.Identifier) z.Lit {
	return lm.LogicCircuit().Or(lm.LitOf(subject).Not(), lm.LitOf(deppy.Identifier(constraint)).Not())
}

func (constraint Conflict) Order() []deppy.Identifier {
	return nil
}

func (constraint Conflict) Anchor() bool {
	return false
}

type AtMost struct {
	ids []deppy.Identifier
	n   int
}

func NewAtMost(n int, ids ...deppy.Identifier) deppy.Constraint {
	return AtMost{
		ids: ids,
		n:   n,
	}
}

func (constraint AtMost) String(subject deppy.Identifier) string {
	s := make([]string, len(constraint.ids))
	for i, each := range constraint.ids {
		s[i] = string(each)
	}
	return fmt.Sprintf("%s permits at most %d of %s", subject, constraint.n, strings.Join(s, ", "))
}

func (constraint AtMost) Apply(lm deppy.LitMapping, subject deppy.Identifier) z.Lit {
	ms := make([]z.Lit, len(constraint.ids))
	for i, each := range constraint.ids {
		ms[i] = lm.LitOf(each)
	}
	return lm.LogicCircuit().CardSort(ms).Leq(constraint.n)
}

func (constraint AtMost) Order() []deppy.Identifier {
	return nil
}

func (constraint AtMost) Anchor() bool {
	return false
}

func (constraint AtMost) IDs() []deppy.Identifier {
	return constraint.ids
}
