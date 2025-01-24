package kernel

import (
	"time"
	"fmt"

	"github.com/jpnt/kman/pkg/logger"
)

type IKernelFacade interface {
	Run() error
}

type KernelFacade struct {
	logger *logger.Logger
	cm     *CommandManager
	ctx    *KernelContext
}

type KernelContext struct {
	tarballURL    string
	downloadPath  string
	archivePath   string
	directory     string
	signatureURL  string
	configOptions []string
	oldConfigPath string
}

// Ensure struct implements interface
var _ IKernelFacade = (*KernelFacade)(nil)

func NewKernelFacade(l *logger.Logger, cm *CommandManager, kc *KernelContext) *KernelFacade {
	return &KernelFacade{
		logger: l,
		cm:     cm,
		ctx:    kc,
	}
}

func (kf *KernelFacade) Run() error {
	kf.logger.Info("Executing all given commands ...")
	startTime := time.Now()

	if err := kf.cm.ExecuteAll(); err != nil {
		kf.logger.Error("Command execution failed: %s", err.Error())
		return fmt.Errorf("execution failed: %w", err)
	}

	duration := time.Since(startTime).Seconds()
	kf.logger.Info("Done. Execution time: %.2f seconds", duration)
	return nil
}
