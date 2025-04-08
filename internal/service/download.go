package service

import (
	"github.com/jpnt/kman/internal/core"
	"github.com/jpnt/kman/pkg/logger"
	"github.com/jpnt/kman/pkg/progress"
	"github.com/jpnt/kman/pkg/utils"
)

type DownloadStep struct {
	log *logger.Logger
	ctx *core.KernelContext
}

// Ensure struct implements interface
var _ core.IStep = (*DownloadStep)(nil)

func (s *DownloadStep) Name() string {
	return "download"
}

func (s *DownloadStep) Execute() error {
	s.log.Info("Downloading Linux kernel tarball from URL: %s ...", s.ctx.TarballURL)

	p := &progress.WriteCounter{}
	kernelPath, err := utils.DownloadFile(s.ctx.TarballURL, s.ctx.DownloadPath, p)
	if err != nil {
		return err
	}

	s.ctx.ArchivePath = kernelPath
	s.log.Info("Downloaded Linux kernel tarball to: %s", kernelPath)
	return nil
}
