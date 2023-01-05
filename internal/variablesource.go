package internal

import (
	"context"

	"github.com/operator-framework/deppy/pkg/deppy"
	"github.com/operator-framework/deppy/pkg/deppy/input"
	"github.com/perdasilva/deconflictor/internal/samples"
)

var _ input.VariableSource = &SimpleVariableSource{}

type SimpleVariableSource struct {
	varSource samples.VariableSource
}

func NewVariableSource(sampleVariableSource samples.VariableSource) *SimpleVariableSource {
	return &SimpleVariableSource{
		varSource: sampleVariableSource,
	}
}

func (s *SimpleVariableSource) GetVariables(_ context.Context, _ input.EntitySource) ([]deppy.Variable, error) {
	return s.varSource.Variables(), nil
}
