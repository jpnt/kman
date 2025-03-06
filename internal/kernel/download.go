package kernel

import (
	"github.com/jpnt/kman/pkg/logger"
	"github.com/jpnt/kman/pkg/progress"
	"github.com/jpnt/kman/pkg/utils"
)

type DownloadStep struct {
	logger *logger.Logger
	ctx    *KernelContext
}

// Ensure struct implements interface
var _ IStep = (*DownloadStep)(nil)

func (s *DownloadStep) String() string {
	return "[DownloadStep]"
}

func (s *DownloadStep) Execute() error {
	s.logger.Info("Downloading Linux kernel tarball from URL: %s ...", s.ctx.tarballURL)

	p := &progress.WriteCounter{}
	kernelPath, err := utils.DownloadFile(s.ctx.tarballURL, s.ctx.downloadPath, p)

	if err != nil {
		return err
	}

	s.logger.Info("Downloaded Linux kernel tarball to: %s", kernelPath)
	s.ctx.archivePath = kernelPath

	return nil
}
