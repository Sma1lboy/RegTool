package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type changeAllModel struct{}

func (m changeAllModel) Init() tea.Cmd {
	return nil
}

func (m changeAllModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return GetCommand("mainMenu")
		case "r":
			fmt.Println("Changed all to region")
			return GetCommand("mainMenu")
		}
	}
	return m, nil
}

func (m changeAllModel) View() string {
	return "Change All to Region\n\nPress 'r' to change all to region, 'q' to go back.\n"
}

func init() {
	RegisterCommand("changeAll", "Change All to Region", changeAllModel{})
}
