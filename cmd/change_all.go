package cmd

import (
	"fmt"
	"regtool/source"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
)

type changeAllModel struct {
	cursor     int
	region     string
	stage      int
	successMsg string
}

func (m changeAllModel) Init() tea.Cmd {
	return nil
}

func (m changeAllModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return GetCommand("mainMenu")
		case "enter":
			if m.stage == 0 {
				m.region = regions[m.cursor]
				m.stage = 1
				return m, nil
			} else if m.stage == 1 {
				// Execute the change all operation
				err := source.ChangeAllRegistry(m.region, nil)
				if err != nil {
					m.successMsg = fmt.Sprintf("Error: Failed to change all apps to the %s region\n%s", m.region, err)
				} else {
					m.successMsg = fmt.Sprintf("Successfully changed all apps to region %s", m.region)
				}
				m.stage = 2
				return m, nil
			} else if m.stage == 2 {
				return GetCommand("mainMenu")
			}
		case "up", "k":
			if m.stage == 0 && m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.stage == 0 && m.cursor < len(regions)-1 {
				m.cursor++
			}
		}
	}
	return m, nil
}

func (m changeAllModel) View() string {
	switch m.stage {
	case 0:
		s := "Change All to Region\n\nSelect a region:\n\n"
		for i, region := range regions {
			cursor := " "
			line := region
			if m.cursor == i {
				cursor = ">"
				boldBlue := color.New(color.FgBlue).Add(color.Bold)
				line = boldBlue.Sprintf(region)
			}
			s += fmt.Sprintf("%s %s\n", cursor, line)
		}
		s += "\nPress 'enter' to confirm, 'q' to go back.\n"
		return s
	case 1:
		return fmt.Sprintf(
			"Change All to Region\n\nSelected Region: %s\n\nPress 'enter' to apply change to all apps, 'q' to go back.\n",
			m.region,
		)
	case 2:
		return fmt.Sprintf(
			"%s\n\nPress 'enter' to return to main menu.\n",
			m.successMsg,
		)
	default:
		return "Unexpected stage."
	}
}

func init() {
	RegisterCommand("changeAll", "Change All to Region", changeAllModel{})
}
