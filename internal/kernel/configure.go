package kernel

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/jpnt/kman/pkg/logger"
	"github.com/jpnt/kman/pkg/utils"
)

type ConfigureCommand struct {
	logger *logger.Logger
	ctx    *KernelContext
}

var _ ICommand = (*ConfigureCommand)(nil)

var defaultOptions = []string{"defconfig"}
var validOptions = []string{"defconfig", "menuconfig", "nconfig", "oldconfig"}

func (c *ConfigureCommand) String() string {
	return "Configure"
}

func (c *ConfigureCommand) Execute() error {
	if c.ctx.configOptions == nil {
		c.logger.Info("No config options were detected, setting up default options ...")
		c.ctx.configOptions = defaultOptions
	}

	configOptions := c.ctx.configOptions
	for _, opt := range configOptions {
		if !isValidOption(opt) {
			return fmt.Errorf("invalid configuration option: %s", opt)
		}
	}

	if err := c.copyOldConfig(); err != nil {
		return fmt.Errorf("failed to copy .config: %w", err)
	}

	c.logger.Info("Configuring Linux kernel with: %v ...", configOptions)

	dir := c.ctx.directory
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

	c.logger.Info("Configured Linux kernel")
	return nil
}

func (c *ConfigureCommand) copyOldConfig() error {
	if c.ctx.oldConfigPath == "" {
		c.logger.Warn("Skipping copy of old .config file")
		return nil
	}

	oldConfigPath := filepath.Join(c.ctx.oldConfigPath)
	newConfigPath := filepath.Join(c.ctx.directory, ".config")

	_, err := os.Stat(oldConfigPath)
	if errors.Is(err, os.ErrNotExist) || c.ctx.oldConfigPath == "" {
		c.logger.Warn(".config file not found in %s, skipping copy", oldConfigPath)
		return nil
	}

	c.logger.Info("Copying .config from %s to %s", oldConfigPath, newConfigPath)
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
