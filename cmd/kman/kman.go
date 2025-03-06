package main

import (
	"os"

	"github.com/jpnt/kman/internal/kernel"
	"github.com/jpnt/kman/pkg/logger"
)

func main() {
	l := logger.NewLogger(logger.InfoLevel)
	b := kernel.NewKernelBuilder(l)

	// fmt.Println(args.len())

	// TODO: dynamic argument builder configuration
	// if args.len()
	b = b.WithDefault()
	// else:
	// b = b.WithArguments(args)


	f, err := b.Build()
	if err != nil {
		l.Error("Error: %s", err.Error())
	}

	if f.Run() != nil {
		l.Error("Error: %s", err.Error())
		os.Exit(1)
	}
}
