package kernel

import (
	"fmt"
	"os"
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

func (c *ExtractCommand) Execute() error {
	if c.ctx.archivePath == "" {
		return fmt.Errorf("cannot extract kernel: archive path is empty")
	}

	c.logger.Info("Extracting Linux kernel archive: %s ...", c.ctx.archivePath)

	archivePath := c.ctx.archivePath
	uncompressDir := filepath.Dir(archivePath)
	// Expected extracted kernel directory from filename
	extractedPath := strings.TrimSuffix(archivePath, filepath.Ext(archivePath))
	extractedPath = strings.TrimSuffix(extractedPath, filepath.Ext(extractedPath))

	if _, err := os.Stat(extractedPath); err == nil {
		c.logger.Info("Kernel already extracted: %s", extractedPath)
		c.ctx.directory = extractedPath
		return nil
	}

	done := make(chan bool)
	go utils.ShowSpinner(done)
	defer func() { done <- true }()

	if err := utils.UncompressFile(archivePath, uncompressDir); err != nil {
		return fmt.Errorf("failed to uncompress archive: %w", err)
	}

	c.logger.Info("Extracted Linux kernel: %s", extractedPath)
	c.ctx.directory = extractedPath

	return nil
}
