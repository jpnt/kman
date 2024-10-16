package kernel

import "kman/pkg/logger"

type KernelManager interface {
	ListAvailableVersions() ([]Kernel, error)
	DownloadKernel(sourceURL, destPath string) (string, error)
	VerifyKernelSignature(signatureURL, kernelPath string) error
	UncompressKernel(kernelPath string) (string, error)
	CopyConfigToKernel(sourcePath, destPath string) (string, error)
	ConfigureKernel(kernelPath string) error
	ApplyPatches(patchPaths []string) error
	CompileKernel(kernelPath string, numJobs int) (string, error)
	CopyKernelToBoot(kernelImagePath, destPath string) (string, error)
	RemoveKernel(kernelVersion string) error
}

type KernelManagerImpl struct {
	logger		*logger.Logger
	downloader	Downloader
	configurer	Configurer
	compiler	Compiler
	installer	Installer
}

func NewKernelManager(
	logger *logger.Logger,
	downloader Downloader,
	configurer Configurer,
	compiler Compiler,
	installer Installer,
) KernelManager {
	return &KernelManagerImpl{
		logger:		logger,
		downloader:	downloader,
		configurer:	configurer,
		installer:	installer,
	}
}

func (km *KernelManagerImpl) ListAvailableVersions() ([]Kernel, error) {
	ks.logger.Info("Listing available versions ...")
	kernels, err := km.downloader.ListAvailableVersions()
	return kernels, err
}

func (km *KernelManagerImpl) DownloadKernel(sourceURL, destPath string) error {
	ks.logger.Info("Downloading source: %s ...", sourceURL)
	return km.Downloader.DownloadKernel(sourceURL, destPath)
}
