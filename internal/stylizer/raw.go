package stylizer

import "github.com/operator-framework/deppy/pkg/deppy"

var _ ErrorStylizer = &RawStylizer{}

type RawStylizer struct{}

func (r *RawStylizer) Stylize(err deppy.NotSatisfiable) string {
	return err.Error()
}
