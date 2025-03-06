package kernel

import (
	"fmt"
)

type IKernelFacade interface {
	Run() error
}

type KernelFacade struct {
	pl     *Pipeline
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

func NewKernelFacade(pl *Pipeline, kc *KernelContext) *KernelFacade {
	return &KernelFacade{
		pl:     pl,
		ctx:    kc,
	}
}

func (kf *KernelFacade) Run() error {
	if err := kf.pl.ExecuteAll(); err != nil {
		return fmt.Errorf("execution failed: %w", err)
	}
	return nil
}
