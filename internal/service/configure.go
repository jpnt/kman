package service

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jpnt/kman/internal/core"
	"github.com/jpnt/kman/pkg/logger"
	"github.com/jpnt/kman/pkg/utils"
)

type ConfigureStep struct {
	log *logger.Logger
	ctx *core.KernelContext
}

var _ core.IStep = (*ConfigureStep)(nil)

func (s *ConfigureStep) Name() string {
	return "configure"
}

func (s *ConfigureStep) Execute() error {
	if s.ctx.ConfigOptions == nil {
		defaultConfigOptions := []string{"tinyconfig"}
		s.log.Warn("Kernel config options not provided, defaulting to %s", defaultConfigOptions)
		s.ctx.ConfigOptions = defaultConfigOptions
	}

	if err := s.copyOldConfig(); err != nil {
		return fmt.Errorf("failed to copy .config: %w", err)
	}

	dir := s.ctx.Directory
	if err := os.Chdir(dir); err != nil {
		return fmt.Errorf("failed to change directory to %q: %w", dir, err)
	}

	for _, option := range s.ctx.ConfigOptions {
		s.log.Info("Configuring Linux kernel with: 'make %s' ...", option)

		cmd := exec.Command("make", option)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("kernel configuration failed: %w", err)
		}
	}

	return nil
}

func (s *ConfigureStep) copyOldConfig() error {
	if s.ctx.OldConfigPath == "" {
		s.log.Warn("Old .config path not provided, skipping ...")
		return nil
	}

	oldConfigPath := filepath.Join(s.ctx.OldConfigPath)
	newConfigPath := filepath.Join(s.ctx.Directory, ".config")

	_, err := os.Stat(oldConfigPath)
	if errors.Is(err, os.ErrNotExist) || s.ctx.OldConfigPath == "" {
		s.log.Warn(".config file not found in %s, skipping copy", oldConfigPath)
		return nil
	}

	s.log.Info("Copying .config from %s to %s", oldConfigPath, newConfigPath)
	if err := utils.CopyFile(oldConfigPath, newConfigPath); err != nil {
		return fmt.Errorf("error copying .config: %w", err)
	}
	return nil
}
