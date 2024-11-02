package kernel

import (
	"kman/internal/pkg"
	"kman/pkg/logger"
)

type IKernelBuilder interface {
	WithDownload() *KernelBuilder
	WithDefault() *KernelBuilder
	Build() *KernelFacade
}
 
type KernelBuilder struct {
	logger	*logger.Logger
	cmds	*pkg.CommandManager
}

func NewKernelBuilder(l *logger.Logger) *KernelBuilder {
	return &KernelBuilder{
		logger: l,
		cmds: pkg.NewCommandManager(),
	}
}

func (kb *KernelBuilder) WithDownload() *KernelBuilder {
	cmd := &DownloadCommand{}
	kb.cmds.AddCommand(cmd)
	return kb
}

func (kb *KernelBuilder) WithDefault() *KernelBuilder {
	kb.WithDownload()
	return kb
}

func (kb *KernelBuilder) Build() *KernelFacade {
	return NewKernelFacade(kb.logger, kb.cmds)
}
