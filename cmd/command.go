package cmd

import tea "github.com/charmbracelet/bubbletea"

type Command interface {
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string
}

var commandRegistry = make(map[string]Command)

func RegisterCommand(name string, cmd Command) {
	commandRegistry[name] = cmd
}

func GetCommand(name string) (tea.Model, tea.Cmd) {
	cmd, exists := commandRegistry[name]
	if !exists {
		return nil, nil
	}
	return cmd, cmd.Init()
}

func ListCommands() []string {
	commands := make([]string, 0, len(commandRegistry))
	for name := range commandRegistry {
		commands = append(commands, name)
	}
	return commands
}
