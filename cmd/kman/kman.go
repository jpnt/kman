package main

import (
	"os"

	"github.com/jpnt/kman/internal/core"
	"github.com/jpnt/kman/internal/service"
	"github.com/jpnt/kman/pkg/logger"
)

func main() {
	l := logger.NewLogger(logger.InfoLevel)

	ctx := core.NewKernelContext()
	p := core.NewPipeline(ctx)
	f := service.NewStepFactory()

	b := core.NewPipelineBuilder(l, p, f)

	b = b.WithDefault()
	// TODO: dynamic argument builder configuration

	err := p.Run()
	if err != nil {
		l.Error("Error: %s", err.Error())
		os.Exit(1)
	}
}
