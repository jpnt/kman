package pkg

type Command interface {
	Execute() error
}

type CommandManager struct {
	commands []Command
}

func NewCommandManager() *CommandManager {
	return &CommandManager{}
}

func (cm *CommandManager) AddCommand(cmd Command) {
	cm.commands = append(cm.commands, cmd)
}

func (cm *CommandManager) ExecuteAll() error {
	for _, cmd := range cm.commands {
		if err := cmd.Execute(); err != nil {
			return err
		}
	}
	return nil
}
