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
	c.logger.Info("Extracting Linux kernel archive: %s", c.ctx.archivePath)

	if c.ctx.archivePath == "" {
		return fmt.Errorf("cannot extract kernel: kernel archive not found")
	}

	extractPath, err := extractKernel(c.ctx.archivePath)
	if err != nil {
		return err
	}

	c.logger.Info("Extracted Linux kernel to: %s", extractPath)
	c.ctx.directory = extractPath

	return nil
}

func extractKernel(kernelPath string) (string, error) {
	kernelDirPath := filepath.Dir(kernelPath)

	// Expected extracted kernel directory from filename
	extractKernelPath := strings.TrimSuffix(kernelPath, filepath.Ext(kernelPath))
	extractKernelPath = strings.TrimSuffix(extractKernelPath, filepath.Ext(extractKernelPath))

	if _, err := os.Stat(extractKernelPath); err == nil {
		fmt.Printf("Kernel already extracted at: %s\n", extractKernelPath)
		return extractKernelPath, nil
	}

	// Start spinner in separate goroutine
	done := make(chan bool)
	go utils.ShowSpinner(done)

	err := utils.UncompressFile(kernelPath, kernelDirPath)
	if err != nil {
		done <- true
		return "", err
	}

	done <- true
	return extractKernelPath, nil
}
