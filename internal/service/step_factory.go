package service

import (
	"fmt"

	"github.com/jpnt/kman/internal/core"
	"github.com/jpnt/kman/pkg/logger"
)

type StepFactory struct{}

var _ core.IStepFactory = (*StepFactory)(nil)

func NewStepFactory() core.IStepFactory {
	return &StepFactory{}
}

func (f *StepFactory) AvailableSteps() []string {
	return []string{
		"list", "download", "verify", "extract", "patch",
		"configure", "compile", "install", "initramfs", "bootloader",
	}
}

func (f *StepFactory) CreateStep(name string, log logger.ILogger, ctx core.IKernelContext) (core.IStep, error) {
	// Type assert and convert to concrete type
	l, ok := log.(*logger.Logger)
	if !ok {
		return nil, fmt.Errorf("invalid logger type")
	}
	c, ok := ctx.(*core.KernelContext)
	if !ok {
		return nil, fmt.Errorf("invalid context type")
	}

	switch name {
	case "list":
		return &ListStep{log: l, ctx: c}, nil
	case "download":
		return &DownloadStep{log: l, ctx: c}, nil
	case "verify":
		return &VerifyStep{log: l, ctx: c}, nil
	case "extract":
		return &ExtractStep{log: l, ctx: c}, nil
	case "patch":
		return &PatchStep{log: l, ctx: c}, nil
	case "configure":
		return &ConfigureStep{log: l, ctx: c}, nil
	case "compile":
		return &CompileStep{log: l, ctx: c}, nil
	case "install":
		return &InstallStep{log: l, ctx: c}, nil
	// TODO: not implemented yet
	// case "initramfs":
	// return &InitramfsStep{log: l, ctx: c}, nil
	// case "bootloader":
	// return &BootloaderStep{log: l, ctx: c}, nil
	default:
		return nil, fmt.Errorf("step name is not defined")
	}
}
