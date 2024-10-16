package kernel

import (
	"kman/pkg/logger"
	"kman/internal/pkg"
)

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

func (kb *KernelBuilder) Build() (*KernelFacade, error) {
	return NewKernelFacade(kb.logger, kb.cmds), nil
}
