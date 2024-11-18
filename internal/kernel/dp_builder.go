package kernel

import (
	"kman/pkg/logger"
)

type IKernelBuilder interface {
	WithList() *KernelBuilder
	WithDownload() *KernelBuilder
	WithVerify() *KernelBuilder
	WithExtract() *KernelBuilder
	WithConfigure() *KernelBuilder
	// WithPatch() *KernelBuilder
	// WithCompile() *KernelBuilder
	// WithInstall() *KernelBuilder
	WithDefault() *KernelBuilder
	Build() *KernelFacade
}

type KernelBuilder struct {
	logger *logger.Logger
	cmds   *CommandManager
	ctx    *KernelContext
}

// Ensure struct implements interface
var _ IKernelBuilder = (*KernelBuilder)(nil)

func NewKernelBuilder(l *logger.Logger) *KernelBuilder {
	sharedCtx := &KernelContext{}
	return &KernelBuilder{
		logger: l,
		cmds:   NewCommandManager(),
		ctx:    sharedCtx,
	}
}

func (kb *KernelBuilder) WithList() *KernelBuilder {
	// how can I inject into each command a logger and the context of facade while im in the builder?
	cmd := &ListCommand{
		logger: kb.logger,
		ctx:    kb.ctx,
	}
	kb.cmds.AddCommand(cmd)
	return kb
}

func (kb *KernelBuilder) WithDownload() *KernelBuilder {
	cmd := &DownloadCommand{
		logger: kb.logger,
		ctx:    kb.ctx,
	}
	kb.cmds.AddCommand(cmd)
	return kb
}

func (kb *KernelBuilder) WithVerify() *KernelBuilder {
	cmd := &VerifyCommand{
		logger: kb.logger,
		ctx:    kb.ctx,
	}
	kb.cmds.AddCommand(cmd)
	return kb
}

func (kb *KernelBuilder) WithExtract() *KernelBuilder {
	cmd := &ExtractCommand{
		logger: kb.logger,
		ctx:    kb.ctx,
	}
	kb.cmds.AddCommand(cmd)
	return kb
}

func (kb *KernelBuilder) WithConfigure() *KernelBuilder {
	cmd := &ConfigureCommand{
		logger: kb.logger,
		ctx:    kb.ctx,
	}
	kb.cmds.AddCommand(cmd)
	return kb
}

func (kb *KernelBuilder) WithDefault() *KernelBuilder {
	kb.WithList()
	kb.WithDownload()
	kb.WithVerify()
	kb.WithExtract()
	kb.WithConfigure()
	return kb
}

func (kb *KernelBuilder) Build() *KernelFacade {
	return NewKernelFacade(kb.logger, kb.cmds, kb.ctx)
}
