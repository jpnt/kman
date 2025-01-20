package kernel

import (
	"fmt"
	"path/filepath"
	"strings"

	"kman/pkg/logger"
	"kman/pkg/utils"
)

type ExtractCommand struct {
	logger *logger.Logger
	ctx    *KernelContext
}

// Ensure struct implements interface
var _ ICommand = (*ExtractCommand)(nil)

func (c *ExtractCommand) String() string {
	return "Extract"
}

func (c *ExtractCommand) Execute() error {
	archivePath := c.ctx.archivePath
	if archivePath == "" {
		return fmt.Errorf("cannot extract kernel: archive path is empty")
	}
	
	extractedPath := strings.TrimSuffix(archivePath, filepath.Ext(archivePath))
	extractedPath = strings.TrimSuffix(extractedPath, filepath.Ext(extractedPath))
	
	c.logger.Info("Extracting Linux kernel archive: %s ...", archivePath)
	_, err := utils.UncompressFile(archivePath, filepath.Dir(archivePath))
	if err != nil {
		return fmt.Errorf("failed to uncompress archive: %w", err)
	}

	if !utils.FileExists(extractedPath) {
		return fmt.Errorf("expected extracted kernel path does not exist: %s", extractedPath)
	}

	c.ctx.directory, err = filepath.Abs(extractedPath)
	if err != nil {
		return fmt.Errorf("failed to resolve absolute path: %w", err)
	}
	c.logger.Info("Extracted Linux kernel to: %s", c.ctx.directory)

	if err := utils.RemoveFile(archivePath); err != nil {
		return fmt.Errorf("failed to remove archive file: %w", err)
	}
	c.logger.Info("Removed kernel archive file: %s", archivePath)

	return nil
}
