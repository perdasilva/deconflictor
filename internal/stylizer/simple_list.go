package stylizer

import (
	"fmt"
	"strings"

	"github.com/operator-framework/deppy/pkg/deppy"
)

var _ ErrorStylizer = &SimpleList{}

type SimpleList struct{}

func (p *SimpleList) Stylize(error deppy.NotSatisfiable) string {
	sb := strings.Builder{}
	for _, appliedConstraint := range error {
		sb.WriteString(fmt.Sprintf("%s\n", appliedConstraint.String()))
	}
	return sb.String()
}
