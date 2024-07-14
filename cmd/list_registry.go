package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type listRegistryModel struct {
	appName string
	stage   int
}

func (m listRegistryModel) Init() tea.Cmd {
	return nil
}

func (m listRegistryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return GetCommand("mainMenu")
		default:
			if m.stage == 0 {
				m.appName = msg.String()
				m.stage = 1
			} else {
				fmt.Printf("Listing registries for app: %s\n", m.appName)
				return GetCommand("mainMenu")
			}
		}
	}
	return m, nil
}

func (m listRegistryModel) View() string {
	if m.stage == 0 {
		return "List Registry by App Name\n\nEnter the app name:\n"
	}
	return fmt.Sprintf("App Name: %s\n\nPress any key to list registries for this app, 'q' to go back.\n", m.appName)
}

func init() {
	RegisterCommand("listRegistry", listRegistryModel{})
}
