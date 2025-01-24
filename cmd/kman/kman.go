package main

import (
	"os"

	"github.com/jpnt/kman/internal/kernel"
	"github.com/jpnt/kman/pkg/logger"
)

func main() {
	l := logger.NewLogger(logger.InfoLevel)
	b := kernel.NewKernelBuilder(l)

	// TODO: dynamic argument builder configuration
	// if args.len()
	b = b.WithDefault()
	// else:
	// b = b.WithArguments(args)
	f, _ := b.Build()

	if f.Run() != nil {
		os.Exit(1)
	}
}
