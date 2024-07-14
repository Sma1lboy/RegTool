package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type listAllRegistryModel struct{}

func (m listAllRegistryModel) Init() tea.Cmd {
	return nil
}

func (m listAllRegistryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return GetCommand("mainMenu")
		case "l":
			fmt.Println("Listing all remote registries")
			return GetCommand("mainMenu")
		}
	}
	return m, nil
}

func (m listAllRegistryModel) View() string {
	return "List All Registry\n\nPress 'l' to list all remote registries, 'q' to go back.\n"
}

func init() {
	RegisterCommand("listAllRegistry", "List All Registry", listAllRegistryModel{})
}
