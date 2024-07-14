package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type changeNameRegionModel struct{}

func (m changeNameRegionModel) Init() tea.Cmd {
	return nil
}

func (m changeNameRegionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return GetCommand("mainMenu")
		case "n":
			fmt.Println("Changed name to region")
			return GetCommand("mainMenu")
		}
	}
	return m, nil
}

func (m changeNameRegionModel) View() string {
	return "Change Name to Region\n\nPress 'n' to change name to region, 'q' to go back.\n"
}

func init() {
	RegisterCommand("changeNameRegion", "Change Name to Region", changeNameRegionModel{})
}
