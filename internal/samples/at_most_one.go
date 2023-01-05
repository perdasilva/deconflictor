package samples

import (
	"github.com/operator-framework/deppy/pkg/deppy"
	"github.com/perdasilva/deconflictor/internal/typed_constraints/constraints"
)

var _ VariableSource = &AtMostOne{}

type AtMostOne struct{}

func (s *AtMostOne) Description() string {
	return "at most 1 can be chosen"
}

func (s *AtMostOne) Variables() []deppy.Variable {
	varA := NewVariable("A", constraints.Mandatory())
	varB := NewVariable("B", constraints.Mandatory())
	varC := NewVariable("C", constraints.Mandatory(), constraints.AtMost(1, "A", "B"))
	return []deppy.Variable{varA, varB, varC}
}
