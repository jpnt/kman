package kernel

import "kman/pkg/logger"

type ConfigureCommand struct {
	logger *logger.Logger
	ctx    *KernelContext
}

var _ ICommand = (*ConfigureCommand)(nil)

func (c *ConfigureCommand) Execute() error {
	// If args == "": run default command: make defconfig

	// Parse config options from the arguments passed at runtime

	return nil
}
