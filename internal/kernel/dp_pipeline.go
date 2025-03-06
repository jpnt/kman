package kernel

import (
	"fmt"
)

type IStep interface {
	fmt.Stringer  // enforce String() implementation
	Execute() error
}

type IPipeline interface {
	AddStep(IStep)
	ExecuteAll() error
}

type Pipeline struct {
	steps []IStep
}

// Ensure struct implements interface
var _ IPipeline = (*Pipeline)(nil)

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (pl *Pipeline) AddStep(step IStep) {
	pl.steps = append(pl.steps, step)
}

func (pl *Pipeline) ExecuteAll() error {
	for _, step := range pl.steps {
		if err := step.Execute(); err != nil {
			return fmt.Errorf("step %s failed: %w", step, err)
		}
	}
	return nil
}
