package internal

import (
	"context"
	"fmt"

	"github.com/operator-framework/deppy/pkg/deppy"
	"github.com/operator-framework/deppy/pkg/deppy/solver"
	"github.com/perdasilva/deconflictor/internal/samples"
	"github.com/perdasilva/deconflictor/internal/stylizer"
)

type Runner struct {
	errorStylizer     stylizer.ErrorStylizer
	dummyEntitySource DummyEntitySource
	samples           []samples.VariableSource
}

func NewRunner(errorStylizer stylizer.ErrorStylizer, samples ...samples.VariableSource) *Runner {
	return &Runner{
		errorStylizer:     errorStylizer,
		samples:           samples,
		dummyEntitySource: DummyEntitySource{},
	}
}

func (r *Runner) Run() {
	for _, sample := range r.samples {
		fmt.Printf("Sample: %s\n", sample.Description())
		fmt.Println("Output:")
		satSolver, err := solver.NewDeppySolver(r.dummyEntitySource, NewVariableSource(sample))
		if err != nil {
			fmt.Printf("error creating solver for sample (%s): %s\n", sample.Description(), err)
			continue
		}

		_, err = satSolver.Solve(context.Background())
		if err == nil {
			fmt.Printf("sample did not generate error (%s)\n", sample.Description())
			continue
		}

		switch e := err.(type) {
		case deppy.NotSatisfiable:
			fmt.Printf(r.errorStylizer.Stylize(e))
		default:
			fmt.Printf("sample produced error other than NotSatisfiable: %s (%s)\n", err, sample.Description())
		}
	}
}
