package main

import (
	"kman/internal/kernel"
	"kman/pkg/logger"
)

func main() {
	l := logger.NewLogger(logger.InfoLevel)
	b := kernel.NewKernelBuilder(l)

	// TODO: dynamic argument builder configuration
	// if args.len()
	b = b.WithDefault()
	// else:
	// b = b.WithArguments(args)
	f := b.Build()

	f.Run()
}
