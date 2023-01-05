package constraints

import (
	"github.com/perdasilva/deconflictor/internal/typed_constraints"

	"github.com/operator-framework/deppy/pkg/deppy"
)

// Mandatory returns a Constraint that will permit only solutions that
// contain a particular Variable.
func Mandatory() deppy.Constraint {
	return typed_constraints.Mandatory{}
}

// Prohibited returns a Constraint that will reject any solution that
// contains a particular Variable. Callers may also decide to omit
// an Variable from input to Solve rather than Apply such a
// Constraint.
func Prohibited() deppy.Constraint {
	return typed_constraints.Prohibited{}
}

// Dependency returns a Constraint that will only permit solutions
// containing a given Variable on the condition that at least one
// of the Variables identified by the given Identifiers also
// appears in the solution. Identifiers appearing earlier in the
// argument list have higher preference than those appearing later.
func Dependency(ids ...deppy.Identifier) deppy.Constraint {
	return typed_constraints.Dependency(ids)
}

// Conflict returns a Constraint that will permit solutions containing
// either the constrained Variable, the Variable identified by
// the given Identifier, or neither, but not both.
func Conflict(id deppy.Identifier) deppy.Constraint {
	return typed_constraints.Conflict(id)
}

// AtMost returns a Constraint that forbids solutions that contain
// more than n of the Variables identified by the given
// Identifiers.
func AtMost(n int, ids ...deppy.Identifier) deppy.Constraint {
	return typed_constraints.NewAtMost(n, ids...)
}
