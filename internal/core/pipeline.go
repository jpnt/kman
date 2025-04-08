package core

import (
	"fmt"
	"time"

	"github.com/jpnt/kman/pkg/logger"
)

type IPipeline interface {
	Run() error
}

type Pipeline struct {
	log   logger.ILogger
	steps []IStep
	ctx   IKernelContext
}

// Ensure struct implements interface
var _ IPipeline = (*Pipeline)(nil)

func (pl *Pipeline) Run() error {
	numSteps := len(pl.steps)
	var stepNames []string

	if numSteps == 0 {
		return fmt.Errorf("no steps were configured")
	}

	for _, step := range pl.steps {
		stepNames = append(stepNames, step.Name())
	}

	pl.log.Info("Executing %d steps: %v", numSteps, stepNames)

	for i, step := range pl.steps {
		start := time.Now()
		pl.log.Info("==> Starting step [%d/%d]: %s", i+1, numSteps, step.Name())

		if err := pl.ctx.Validate(step.Name()); err != nil {
			return fmt.Errorf("validation failed for step %q: %w", step.Name(), err)
		}

		if err := step.Execute(); err != nil {
			return fmt.Errorf("execution failed for step %q: %w", step.Name(), err)
		}

		duration := time.Since(start)
		pl.log.Info("<== Completed step [%d/%d]: %s in %s", i+1, numSteps, step.Name(), duration)
	}
	return nil
}
