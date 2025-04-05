package core

import (
	"fmt"
)

type IPipeline interface {
	Ctx() IKernelContext
	AddStep(step IStep)
	Run() error
}

type Pipeline struct {
	steps []IStep
	ctx   IKernelContext
}

// Ensure struct implements interface
var _ IPipeline = (*Pipeline)(nil)

// func NewPipeline(c IKernelContext) *Pipeline {
	// return &Pipeline{ctx: c}
// }

func (pl *Pipeline) Ctx() IKernelContext {
	return pl.ctx
}

func (pl *Pipeline) AddStep(step IStep) {
	pl.steps = append(pl.steps, step)
}

func (pl *Pipeline) Run() error {
	if len(pl.steps) == 0 {
		return fmt.Errorf("no steps were configured")
	}

	for _, step := range pl.steps {
		if err := pl.ctx.Validate(step.Name()); err != nil {
			return fmt.Errorf("validation failed for step %q: %w", step.Name(), err)
		}

		if err := step.Execute(); err != nil {
			return fmt.Errorf("execution failed for step %q: %w", step.Name(), err)
		}
	}
	return nil
}
