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
	pipeline 
}

// Ensure struct implements interface
var _ IPipelineBuilder = (*PipelineBuilder)(nil)

func NewPipelineBuilder(l logger.ILogger, c IKernelContext, f IStepFactory) IPipelineBuilder {
	return &PipelineBuilder{logger: l, context: c, factory: f}
}

func (b *PipelineBuilder) WithStep(stepName string) IPipelineBuilder {
	step, err := b.factory.CreateStep(stepName, b.logger, b.pl.Ctx())
	if err != nil {
		b.logger.Warn("Unrecognized step: %q", stepName)
		return b
	}
	b.pl.AddStep(step)
	return b
}

func (b *PipelineBuilder) WithDefault() IPipelineBuilder {
	return b.
		WithStep("list").
		WithStep("download").
		WithStep("verify").
		WithStep("extract").
		WithStep("patch").
		WithStep("configure").
		WithStep("compile").
		WithStep("install").
		WithStep("initramfs").
		WithStep("bootloader")
}

func (b *PipelineBuilder) Build() IPipeline {
	return &Pipeline{ctx}
}
