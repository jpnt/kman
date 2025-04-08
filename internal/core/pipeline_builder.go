package core

// https://refactoring.guru/design-patterns/builder

import (
	"github.com/jpnt/kman/pkg/logger"
)

type IPipelineBuilder interface {
	WithStep(stepName string) IPipelineBuilder
	WithDefault() IPipelineBuilder
	Build() IPipeline
}

type PipelineBuilder struct {
	logger  logger.ILogger
	factory IStepFactory
	ctx     IKernelContext
	steps   []IStep
}

// Ensure struct implements interface
var _ IPipelineBuilder = (*PipelineBuilder)(nil)

func NewPipelineBuilder(l logger.ILogger, f IStepFactory, c IKernelContext) IPipelineBuilder {
	return &PipelineBuilder{logger: l, factory: f, ctx: c}
}

func (b *PipelineBuilder) WithStep(stepName string) IPipelineBuilder {
	step, err := b.factory.CreateStep(stepName, b.logger, b.ctx)
	if err != nil {
		b.logger.Warn("Could not create step %q: %s", stepName, err)
		return b
	}
	b.steps = append(b.steps, step)

	return b
}

func (b *PipelineBuilder) WithDefault() IPipelineBuilder {
	availableSteps := b.factory.AvailableSteps()

	for _, stepName := range availableSteps {
		b.WithStep(stepName)
	}

	return b
}

func (b *PipelineBuilder) Build() IPipeline {
	return &Pipeline{log: b.logger, ctx: b.ctx, steps: b.steps}
}
