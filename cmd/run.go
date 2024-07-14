package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type mainMenuModel struct {
	cursor  int
	choices []string
}

func newMainMenuModel() mainMenuModel {
	return mainMenuModel{
		choices: ListCommands(),
	}
}

func (m mainMenuModel) Init() tea.Cmd {
	return nil
}

func (m mainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			cmd, initCmd := GetCommand(m.choices[m.cursor])
			if cmd != nil {
				return cmd, initCmd
			}
		}
	}
	return m, nil
}

func (m mainMenuModel) View() string {
	s := "RegistryHub\n\n"
	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nPress 'q' to quit.\n"
	return s
}

func Run() {
	RegisterCommand("mainMenu", newMainMenuModel())

	p := tea.NewProgram(newMainMenuModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
