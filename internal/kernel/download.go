package kernel

import (
	"github.com/jpnt/kman/pkg/logger"
	"github.com/jpnt/kman/pkg/progress"
	"github.com/jpnt/kman/pkg/utils"
)

type DownloadCommand struct {
	logger *logger.Logger
	ctx    *KernelContext
}

// Ensure struct implements interface
var _ ICommand = (*DownloadCommand)(nil)

func (c *DownloadCommand) String() string {
	return "[DownloadCommand]"
}

func (c *DownloadCommand) Execute() error {
	c.logger.Info("Downloading Linux kernel tarball from URL: %s ...", c.ctx.tarballURL)

	p := &progress.WriteCounter{}
	kernelPath, err := utils.DownloadFile(c.ctx.tarballURL, c.ctx.downloadPath, p)

	if err != nil {
		return err
	}

	c.logger.Info("Downloaded Linux kernel tarball to: %s", kernelPath)
	c.ctx.archivePath = kernelPath

	return nil
}
