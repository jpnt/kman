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

	archivePath, err := filepath.Abs(c.ctx.archivePath)
	if err != nil {
		return fmt.Errorf("failed to resolve absolute path for archive: %w", err)
	}

	if _, err := os.Stat(archivePath); os.IsNotExist(err) {
		return fmt.Errorf("archive file does not exist: %s", archivePath)
	}

	uncompressDir := filepath.Dir(archivePath)

	// Guessing expected extracted kernel directory from filename
	extractedPath := strings.TrimSuffix(archivePath, filepath.Ext(archivePath))
	extractedPath = strings.TrimSuffix(extractedPath, filepath.Ext(extractedPath))

	// TEST: let the extract tool deal with it...
	// if _, err := os.Stat(extractedPath); err == nil {
		// c.logger.Info("Kernel already extracted: %s", extractedPath)
		// c.ctx.directory = extractedPath
		// return nil
	// }

	// done := make(chan bool)
	// go utils.ShowSpinner(done)
	// defer func() { done <- true }()

	// TODO: extractedPath should be result of UncompressFile instead of
	// guessing. Also instead of utils.ShowSpinner in a coroutine do it better.
	if err := utils.UncompressFile(archivePath, uncompressDir); err != nil {
		return fmt.Errorf("failed to uncompress archive: %w", err)
	}
	c.logger.Info("Extracted Linux kernel: %s", extractedPath)
	c.ctx.directory = extractedPath

	if err := utils.RemoveFile(archivePath); err != nil {
		return fmt.Errorf("failed to remove archive: %w", err)
	}
	c.logger.Info("Extracted Linux kernel: %s", extractedPath)

	return nil
}
