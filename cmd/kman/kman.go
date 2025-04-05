package main

import (
	"os"

	"github.com/jpnt/kman/internal/core"
	"github.com/jpnt/kman/internal/service"
	"github.com/jpnt/kman/pkg/logger"
)

func main() {
	log := logger.NewLogger(logger.InfoLevel)
	ctx := core.NewKernelContext()
	// pipeline := core.NewPipeline(ctx)
	fac := service.NewStepFactory()
	plb := core.NewPipelineBuilder(log, ctx, fac)

	pipeline = plb.WithDefault()

	err := pipeline.Run()
	if err != nil {
		log.Error("%s", err.Error())
		os.Exit(1)
	}
}
