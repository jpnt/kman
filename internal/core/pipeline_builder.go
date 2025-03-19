package core

import (
	"github.com/jpnt/kman/pkg/logger"
)

type IPipelineBuilder interface {
	WithStep(stepName string) IPipelineBuilder
	WithDefault() IPipelineBuilder
}

type PipelineBuilder struct {
	logger  logger.ILogger
	pl      IPipeline
	factory IStepFactory
}

// Ensure struct implements interface
var _ IPipelineBuilder = (*PipelineBuilder)(nil)

func NewPipelineBuilder(l logger.ILogger, p IPipeline, f IStepFactory) IPipelineBuilder {
	return &PipelineBuilder{logger: l, pl: p, factory: f}
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
