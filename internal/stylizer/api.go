package stylizer

import "github.com/operator-framework/deppy/pkg/deppy"

type ErrorStylizer interface {
	Stylize(error deppy.NotSatisfiable) string
}
