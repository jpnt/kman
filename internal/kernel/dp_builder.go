package kernel

import (
	"fmt"

	"github.com/jpnt/kman/pkg/logger"
)

// TODO: ensure execution follows this order. why: useful for dynamic config later on
type IKernelBuilder interface {
	WithCommand(name string) IKernelBuilder
	WithDefault() IKernelBuilder
	Build() (IKernelFacade, error)
}

type KernelBuilder struct {
	logger   *logger.Logger
	pl       *Pipeline
	ctx      *KernelContext
	registry map[string]func(*logger.Logger, *KernelContext) IStep
}

// Ensure struct implements interface
var _ IKernelBuilder = (*KernelBuilder)(nil)

func NewKernelBuilder(l *logger.Logger) IKernelBuilder {
	sharedCtx := &KernelContext{}

	// Command registry
	registry := map[string]func(*logger.Logger, *KernelContext) IStep{
		"list":      func(l *logger.Logger, c *KernelContext) IStep { return &ListStep{logger: l, ctx: c} },
		"download":  func(l *logger.Logger, c *KernelContext) IStep { return &DownloadStep{logger: l, ctx: c} },
		"verify":    func(l *logger.Logger, c *KernelContext) IStep { return &VerifyStep{logger: l, ctx: c} },
		"extract":   func(l *logger.Logger, c *KernelContext) IStep { return &ExtractStep{logger: l, ctx: c} },
		"patch":     func(l *logger.Logger, c *KernelContext) IStep { return &PatchStep{logger: l, ctx: c} },
		"configure": func(l *logger.Logger, c *KernelContext) IStep { return &ConfigureStep{logger: l, ctx: c} },
		"compile":   func(l *logger.Logger, c *KernelContext) IStep { return &CompileStep{logger: l, ctx: c} },
		"install":   func(l *logger.Logger, c *KernelContext) IStep { return &InstallStep{logger: l, ctx: c} },
	}

	return &KernelBuilder{
		logger:   l,
		pl:       NewPipeline(),
		ctx:      sharedCtx,
		registry: registry,
	}
}

func (kb *KernelBuilder) WithCommand(name string) IKernelBuilder {
	if constructor, ok := kb.registry[name]; ok {
		kb.pl.AddStep(constructor(kb.logger, kb.ctx))
	} else {
		kb.logger.Warn("Command '%s' is not recognized", name)
	}
	return kb
}

func (kb *KernelBuilder) WithDefault() IKernelBuilder {
	return kb.
		WithCommand("list").
		WithCommand("download").
		WithCommand("verify").
		WithCommand("extract").
		WithCommand("patch").
		WithCommand("configure").
		WithCommand("compile").
		WithCommand("install")

}

func (kb *KernelBuilder) Build() (IKernelFacade, error) {
	if len(kb.pl.steps) == 0 {
		return nil, fmt.Errorf("no commands were configured in the KernelBuilder")
	}
	return NewKernelFacade(kb.pl, kb.ctx), nil
}
