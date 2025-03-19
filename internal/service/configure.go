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

var defaultOptions = []string{"defconfig"}
var validOptions = []string{"defconfig", "menuconfig", "nconfig", "oldconfig"}

func (s *ConfigureStep) Name() string {
	return "configure"
}

func (s *ConfigureStep) Execute() error {
	if s.ctx.ConfigOptions == nil {
		s.log.Info("No config options were detected, setting up default options ...")
		s.ctx.ConfigOptions = defaultOptions
	}

	// TODO: better option than doing this
	configOptions := s.ctx.ConfigOptions
	for _, opt := range configOptions {
		if !isValidOption(opt) {
			return fmt.Errorf("invalid configuration option: %s", opt)
		}
	}

	if err := s.copyOldConfig(); err != nil {
		return fmt.Errorf("failed to copy .config: %w", err)
	}

	s.log.Info("Configuring Linux kernel with: %v ...", configOptions)

	dir := s.ctx.Directory
	if err := os.Chdir(dir); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", dir, err)
	}

	for _, opt := range configOptions {
		cmd := exec.Command("make", opt)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("kernel configuration failed: %w", err)
		}
	}

	s.log.Info("Configured Linux kernel")
	return nil
}

func (s *ConfigureStep) copyOldConfig() error {
	if s.ctx.OldConfigPath == "" {
		s.log.Warn("Skipping copy of old .config file")
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

func isValidOption(option string) bool {
	for _, validOption := range validOptions {
		if option == validOption {
			return true
		}
	}
	return false
}
