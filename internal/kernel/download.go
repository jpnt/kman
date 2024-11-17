package kernel

import (
	"kman/pkg/logger"
	"kman/pkg/progress"
	"kman/pkg/utils"
)

type DownloadCommand struct {
	logger *logger.Logger
	ctx    *KernelContext
}

// Ensure struct implements interface
var _ ICommand = (*DownloadCommand)(nil)

func (c *DownloadCommand) Execute() error {
	p := &progress.WriteCounter{}

	kernelPath, err := utils.DownloadFile(c.ctx.sourceURL, c.ctx.downloadPath, p)

	if err != nil {
		return err
	}

	c.logger.Info("Kernel archive path: %s\n", kernelPath)
	c.ctx.archivePath = kernelPath

	return nil
}
