package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type updateModel struct{}

func (m updateModel) Init() tea.Cmd {
	return nil
}

func (m updateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return GetCommand("mainMenu")
		case "u":
			fmt.Println("Checking all apps and querying local software")
			return GetCommand("mainMenu")
		}
	}
	return m, nil
}

func (m updateModel) View() string {
	return "Update (Check All Apps)\n\nPress 'u' to check all apps and query local software, 'q' to go back.\n"
}

func init() {
	RegisterCommand("update", "Update (Check All Apps)", updateModel{})
}
