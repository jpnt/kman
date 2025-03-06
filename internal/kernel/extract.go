package kernel

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/jpnt/kman/pkg/logger"
	"github.com/jpnt/kman/pkg/utils"
)

type ExtractStep struct {
	logger *logger.Logger
	ctx    *KernelContext
}

// Ensure struct implements interface
var _ IStep = (*ExtractStep)(nil)

func (s *ExtractStep) String() string {
	return "[ExtractStep]"
}

func (s *ExtractStep) Execute() error {
	archivePath := s.ctx.archivePath
	if archivePath == "" {
		return fmt.Errorf("cannot extract kernel: archive path is empty")
	}
	
	extractedPath := strings.TrimSuffix(archivePath, filepath.Ext(archivePath))
	extractedPath = strings.TrimSuffix(extractedPath, filepath.Ext(extractedPath))
	
	s.logger.Info("Extracting Linux kernel archive: %s ...", archivePath)
	err := utils.UncompressFile(archivePath, filepath.Dir(archivePath))
	if  err != nil {
		return fmt.Errorf("failed to uncompress archive: %w", err)
	}

	if !utils.FileExists(extractedPath) {
		return fmt.Errorf("expected extracted kernel path does not exist: %s", extractedPath)
	}

	s.ctx.directory, err = filepath.Abs(extractedPath)
	if err != nil {
		return fmt.Errorf("failed to resolve absolute path: %w", err)
	}
	s.logger.Info("Extracted Linux kernel to: %s", s.ctx.directory)

	if err := utils.RemoveFile(archivePath); err != nil {
		return fmt.Errorf("failed to remove archive file: %w", err)
	}
	s.logger.Info("Removed kernel archive file: %s", archivePath)

	return nil
}
