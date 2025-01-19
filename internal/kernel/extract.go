package kernel

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"kman/pkg/logger"
	// "kman/pkg/spinner"
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
	archivePath = c.ctx.archivePath

	if c.ctx.archivePath == "" {
		return fmt.Errorf("cannot extract kernel: archive path is empty")
	}

	c.logger.Info("Extracting Linux kernel archive: %s ...", c.ctx.archivePath)

	// archivePath, err := filepath.Abs(c.ctx.archivePath)
	// if err != nil {
		// return fmt.Errorf("failed to resolve absolute path for archive: %w", err)
	// }

	if _, err := os.Stat(archivePath); os.IsNotExist(err) {
		return fmt.Errorf("archive file does not exist: %s", archivePath)
	}

	// TEMP: Guessing expected extracted kernel directory from filename
	extractedPath := strings.TrimSuffix(archivePath, filepath.Ext(archivePath))
	extractedPath = strings.TrimSuffix(extractedPath, filepath.Ext(extractedPath))

	// TODO: extractedPath should be result of UncompressFile instead of
	// guessing. Also instead of utils.ShowSpinner in a coroutine do it better.
	// extractedPath, err := utils.UncompressFile(archivePath, filepath.Dir(archivePath))
	_, err = utils.UncompressFile(archivePath, filepath.Dir(archivePath))
	if err != nil {
		return fmt.Errorf("failed to uncompress archive: %w", err)
	}
	c.logger.Info("Extracted Linux kernel: %s", extractedPath)
	c.ctx.directory = extractedPath

	if err := utils.RemoveFile(archivePath); err != nil {
		return fmt.Errorf("failed to remove archive file: %w", err)
	}
	c.logger.Info("Removed kernel archive file: %s", archivePath)

	return nil
}
