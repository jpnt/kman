package kernel

import (
	"kman/pkg/logger"
	"os"
)

type IKernelFacade interface {
	Run()
}

type KernelFacade struct {
	logger *logger.Logger
	cmds   *CommandManager
	ctx    *KernelContext
}

type KernelContext struct {
	tarballURL   string
	downloadPath string
	archivePath  string
	directory    string
	signatureURL string
	skipVerify   bool
}

// Ensure struct implements interface
var _ IKernelFacade = (*KernelFacade)(nil)

func NewKernelFacade(l *logger.Logger, cm *CommandManager) *KernelFacade {
	return &KernelFacade{
		logger: l,
		cmds:   cm,
	}
}

func (kf *KernelFacade) Run() {
	kf.logger.Info("Executing all given commands ...")
	if err := kf.cmds.ExecuteAll(); err != nil {
		kf.logger.Error("failed to execute commands: %s", err.Error())
		os.Exit(1)
	}
}
