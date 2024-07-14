package cmd

import tea "github.com/charmbracelet/bubbletea"

type Command interface {
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string
}

type CommandInfo struct {
	Name        string
	Description string
	Command     Command
}

var commandRegistry = make(map[string]CommandInfo)
var mainMenuName = "mainMenu"

func RegisterCommand(name, description string, cmd Command) {
	commandRegistry[name] = CommandInfo{
		Name:        name,
		Description: description,
		Command:     cmd,
	}
}

func GetCommand(name string) (tea.Model, tea.Cmd) {
	info, exists := commandRegistry[name]
	if !exists {
		return nil, nil
	}
	return info.Command, info.Command.Init()
}

func ListCommandDescriptions() []string {
	descriptions := make([]string, 0, len(commandRegistry))
	for _, info := range commandRegistry {
		if info.Name != mainMenuName {
			descriptions = append(descriptions, info.Description)
		}
	}
	return descriptions
}

func ListCommandNames() []string {
	names := make([]string, 0, len(commandRegistry))
	for name := range commandRegistry {
		if name != mainMenuName {
			names = append(names, name)
		}
	}
	return names
}
