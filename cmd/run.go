package cmd

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	primaryColor    = lipgloss.Color("#61AFEF")
	bgColor         = lipgloss.Color("#282C34")
	textColor       = lipgloss.Color("#ABB2BF")
	selectedColor   = lipgloss.Color("#98C379")
	unselectedColor = lipgloss.Color("#3E4451")
	accentColor     = lipgloss.Color("#C678DD")

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#E5C07B")).
			Bold(true).
			MarginBottom(1)

	optionStyle = lipgloss.NewStyle().
			Foreground(textColor)

	selectedOptionStyle = lipgloss.NewStyle().
				Foreground(selectedColor).
				Bold(true)

	quitTextStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Italic(true)
)

type mainMenuModel struct {
	cursor  int
	choices []string
	names   []string
	width   int
}

func newMainMenuModel() mainMenuModel {
	return mainMenuModel{
		choices: ListCommandDescriptions(),
		names:   ListCommandNames(),
		width:   80, // Default width
	}
}

func (m mainMenuModel) Init() tea.Cmd {
	return nil
}

func (m mainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
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
	doc := strings.Builder{}

	// Title
	doc.WriteString(titleStyle.Render("🔧 RegistryHub") + "\n")

	// Menu options
	for i, choice := range m.choices {
		cursor := "  "
		style := optionStyle
		if m.cursor == i {
			cursor = "▶ "
			style = selectedOptionStyle
		}
		doc.WriteString(style.Render(cursor+choice) + "\n")
	}

	// Quit instruction
	doc.WriteString("\n" + quitTextStyle.Render("Press 'q' to quit"))

	return doc.String()
}

func Run() {
	RegisterCommand(mainMenuName, "Main Menu", newMainMenuModel())
	p := tea.NewProgram(newMainMenuModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
