package cmd

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	regions = []string{"us", "cn", "eu", "jp"}
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
		width:   80,
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
	doc.WriteString(GetStyledTitle("RegistryHub") + "\n")

	// Menu options
	for i, choice := range m.choices {
		doc.WriteString(GetStyledOption(choice, m.cursor == i) + "\n")
	}

	// Quit instruction
	doc.WriteString("\n" + GetStyledQuitText())

	return borderedBox(doc.String())
}
func Run() {
	RegisterCommand(mainMenuName, "Main Menu", newMainMenuModel())
	p := tea.NewProgram(newMainMenuModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, errorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		os.Exit(1)
	}
}
