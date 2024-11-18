package kernel

import (
	"errors"
	"fmt"
	"kman/pkg/logger"
	"kman/pkg/utils"
	"os"
	"os/exec"
	"path/filepath"
)

type ConfigureCommand struct {
	logger *logger.Logger
	ctx    *KernelContext
}

var _ ICommand = (*ConfigureCommand)(nil)

var defaultOption = "defconfig"
var options = []string{"defconfig", "menuconfig", "nconfig", "oldconfig"}

func (c *ConfigureCommand) Execute() error {
	if c.ctx.configOption == "" {
		c.ctx.configOption = defaultOption
	}

	if !isValidOption(c.ctx.configOption) {
		return fmt.Errorf("invalid configuration option: %s", c.ctx.configOption)
	}

	if err := c.copyOldConfig(); err != nil {
		return fmt.Errorf("failed to copy .config: %w", err)
	}

	c.logger.Info("Configuring Linux kernel with: make %s", c.ctx.configOption)

	if err := os.Chdir(c.ctx.directory); err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", c.ctx.directory, err)
	}

	cmd := exec.Command("make", c.ctx.configOption)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("kernel configuration failed: %w", err)
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
	for _, validOption := range options {
		if option == validOption {
			return true
		}
	}
	return false
}
