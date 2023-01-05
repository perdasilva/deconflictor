package samples

import (
	"github.com/operator-framework/deppy/pkg/deppy"
	"github.com/perdasilva/deconflictor/internal/typed_constraints/constraints"
)

var _ VariableSource = &SimpleConflict{}

type SimpleConflict struct{}

func (s *SimpleConflict) Description() string {
	return "simple conflict between two variables"
}

func (s *SimpleConflict) Variables() []deppy.Variable {
	varA := NewVariable("A", constraints.Mandatory(), constraints.Conflict("B"))
	varB := NewVariable("B", constraints.Mandatory())
	return []deppy.Variable{varA, varB}
}
