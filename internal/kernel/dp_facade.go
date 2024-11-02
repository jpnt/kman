package kernel

import (
	"kman/internal/pkg"
	"kman/pkg/logger"
)

type IKernelFacade interface {
	Run()
}

type KernelFacade struct {
	logger	*logger.Logger
	cmds	*pkg.CommandManager
}

func NewKernelFacade(l *logger.Logger, cm *pkg.CommandManager) *KernelFacade {
	return &KernelFacade{
		logger: l,
		cmds: cm,
	}
}

func (kf *KernelFacade) Run() {
	kf.logger.Info("Executing all given commands ...")
	if err := kf.cmds.ExecuteAll(); err != nil {
		kf.logger.Error("failed to execute commands: %s", err.Error())
	}
}
