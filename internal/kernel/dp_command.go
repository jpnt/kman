package kernel

import "fmt"

type ICommand interface {
	fmt.Stringer  // enforce String() implementation
	Execute() error
}

type ICommandManager interface {
	AddCommand(ICommand)
	ExecuteAll() error
}

type CommandManager struct {
	commands []ICommand
}

// Ensure struct implements interface
var _ ICommandManager = (*CommandManager)(nil)

func NewCommandManager() *CommandManager {
	return &CommandManager{}
}

func (cm *CommandManager) AddCommand(cmd ICommand) {
	cm.commands = append(cm.commands, cmd)
}

func (cm *CommandManager) ExecuteAll() error {
	for _, cmd := range cm.commands {
		if err := cmd.Execute(); err != nil {
			return fmt.Errorf("%s: %w", cmd, err)
		}
	}
	return nil
}
