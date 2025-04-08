package main

import (
	"os"

	"github.com/jpnt/kman/internal/core"
	"github.com/jpnt/kman/internal/service"
	"github.com/jpnt/kman/pkg/logger"
)

var Version string = "0.0.3"

func main() {
	log := logger.NewLogger(logger.InfoLevel)
	ctx := core.NewKernelContext()
	factory := service.NewStepFactory()
	builder := core.NewPipelineBuilder(log, factory, ctx)

	pipeline := builder.WithDefault().Build()
	log.Info("kman version %s", Version)

	err := pipeline.Run()
	if err != nil {
		log.Error("%s", err.Error())
		os.Exit(1)
	}
}
