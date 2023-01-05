package samples

import (
	"github.com/operator-framework/deppy/pkg/deppy"
	"github.com/perdasilva/deconflictor/internal/typed_constraints/constraints"
)

var _ VariableSource = &DependencyConflict{}

type DependencyConflict struct{}

func (s *DependencyConflict) Description() string {
	return "dependency of dependency conflicts with other top level variable"
}

func (s *DependencyConflict) Variables() []deppy.Variable {
	varA := NewVariable("A", constraints.Mandatory(), constraints.Dependency("B"))
	varB := NewVariable("B", constraints.Dependency("C"))
	varC := NewVariable("C")
	varD := NewVariable("D", constraints.Mandatory(), constraints.Conflict("C"))
	return []deppy.Variable{varA, varB, varC, varD}
}
