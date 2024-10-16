package kernel

import (
	"fmt"

	"kman/pkg/logger"
	"kman/internal/pkg"
)

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

func (kf *KernelFacade) ManageKernel() error {
	kf.logger.Info("Executing all given commands ...")
	if err := kf.cmds.ExecuteAll(); err != nil {
		return fmt.Errorf("failed to execute commands: %w", err)
	}
	return nil
}
