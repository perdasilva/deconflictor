package samples

import "github.com/operator-framework/deppy/pkg/deppy"

type VariableSource interface {
	Description() string
	Variables() []deppy.Variable
}

var _ deppy.Variable = &Variable{}

type Variable struct {
	id          deppy.Identifier
	constraints []deppy.Constraint
}

func NewVariable(id string, constraints ...deppy.Constraint) *Variable {
	return &Variable{
		id:          deppy.IdentifierFromString(id),
		constraints: constraints,
	}
}

func (v *Variable) Identifier() deppy.Identifier {
	return v.id
}

func (v *Variable) Constraints() []deppy.Constraint {
	return v.constraints
}

func (v *Variable) AddConstraints(constraint ...deppy.Constraint) {
	v.constraints = append(v.constraints, constraint...)
}
