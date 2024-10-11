package pkg

type Command interface {
	Execute() error
}

type CommandManager struct {
	commands []Command
}

func (cm *CommandManager) ExecuteAll() {
	for _, cmd := range cm.commands {
		cmd.Execute()
	}
}
