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
	if len(pl.steps) == 0 {
		return fmt.Errorf("no steps were configured")
	}

	pl.log.Info("The following steps will be executed:")
	for _, step := range pl.steps {
		pl.log.Info("- %s", step.Name())
	}

	pl.log.Info("Starting execution ...")

	for _, step := range pl.steps {
		start := time.Now()
		pl.log.Info("==> Starting step: %s ...", step.Name())

		if err := pl.ctx.Validate(step.Name()); err != nil {
			return fmt.Errorf("validation failed for step %q: %w", step.Name(), err)
		}

		if err := step.Execute(); err != nil {
			return fmt.Errorf("execution failed for step %q: %w", step.Name(), err)
		}

		duration := time.Since(start)
		pl.log.Info("<== Completed step: %s in %s", step.Name(), duration)
	}
	return nil
}
