package kernel

import (
	"fmt"
	"kman/pkg/logger"
	"os"
	"os/exec"
)

type ConfigureCommand struct {
	logger *logger.Logger
	ctx    *KernelContext
}

var _ ICommand = (*ConfigureCommand)(nil)

var defaultOption = "defconfig"
var options = []string{"defconfig", "menuconfig", "nconfig", "oldconfig"}

func (c *ConfigureCommand) Execute() error {
	// TODO
	o := defaultOption
	c.logger.Info("Configuring Linux kernel with: make %s", o)

	cmd := exec.Command("make", o)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("kernel configuration failed: %w", err)
	}

	c.logger.Info("Configured Linux kernel")

	return nil
}
