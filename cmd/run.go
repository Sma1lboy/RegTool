package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	primaryColor   = lipgloss.Color("#4A90E2")
	secondaryColor = lipgloss.Color("#6FCF97")
	bgColor        = lipgloss.Color("#F0F4F8")
	textColor      = lipgloss.Color("#333333")

	titleStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Background(bgColor).
			Bold(true).
			Padding(0, 1).
			MarginBottom(1).
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(primaryColor)
	optionStyle = lipgloss.NewStyle().
			Padding(0, 1)
	selectedOptionStyle = optionStyle.Copy().
				Foreground(primaryColor).
				Background(lipgloss.Color("#E6F0FF")).
				Bold(true)
	quitTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Italic(true)
	frameStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(1).
			Background(bgColor)
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
	s := titleStyle.Render("RegistryHub") + "\n"

	for i, choice := range m.choices {
		if m.cursor == i {
			s += selectedOptionStyle.Render(fmt.Sprintf(" • %s", choice)) + "\n"
		} else {
			s += optionStyle.Render(fmt.Sprintf("   %s", choice)) + "\n"
		}
	}

	s += "\n" + quitTextStyle.Render("Press 'q' to quit.")
	return frameStyle.Render(s)
}

func Run() {
	RegisterCommand(mainMenuName, "Main Menu", newMainMenuModel())

	p := tea.NewProgram(newMainMenuModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
