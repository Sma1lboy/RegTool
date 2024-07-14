package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
)

type mainMenuModel struct {
	cursor  int
	choices []string
	names   []string
}

func newMainMenuModel() mainMenuModel {
	return mainMenuModel{
		choices: ListCommandDescriptions(),
		names:   ListCommandNames(),
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
			cmd, initCmd := GetCommand(m.names[m.cursor])
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
		line := choice
		if m.cursor == i {
			cursor = ">" // cursor!
			// Use color package to bold and color the selected line
			bold := color.New(color.FgBlue).Add(color.Bold)
			line = bold.Sprintf(choice)
		}
		s += fmt.Sprintf("%s %s\n", cursor, line)
	}
	s += "\nPress 'q' to quit.\n"
	return s
}

func Run() {
	RegisterCommand(mainMenuName, "Main Menu", newMainMenuModel())

	p := tea.NewProgram(newMainMenuModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
