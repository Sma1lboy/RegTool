package cmd

import (
	"sort"

	tea "github.com/charmbracelet/bubbletea"
)

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

var commandRegistry = map[string]CommandInfo{}

const mainMenuName = "mainMenu"

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
	keys := make([]string, 0, len(commandRegistry))
	for key := range commandRegistry {
		if key != mainMenuName {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	descriptions := make([]string, len(keys))
	for i, key := range keys {
		descriptions[i] = commandRegistry[key].Description
	}
	return descriptions
}

func ListCommandNames() []string {
	keys := make([]string, 0, len(commandRegistry))
	for key := range commandRegistry {
		if key != mainMenuName {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	names := make([]string, len(keys))
	for i, key := range keys {
		names[i] = commandRegistry[key].Name
	}
	return names
}
