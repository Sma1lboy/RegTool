package cmd

import (
	"fmt"
	"log"
	"registryhub/source"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
)

var regions = []string{"us", "cn", "eu", "jp"}

type changeNameRegionModel struct {
	input      textinput.Model
	stage      int
	appName    string
	region     string
	cursor     int
	successMsg string
}

func newChangeNameRegionModel() changeNameRegionModel {
	ti := textinput.New()
	ti.Placeholder = "Enter app name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return changeNameRegionModel{input: ti, stage: 0, cursor: 0}
}

func (m changeNameRegionModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m changeNameRegionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return GetCommand(mainMenuName)
		case "enter":
			if m.stage == 0 {
				if m.input.Value() == "" {
					return m, nil
				}
				m.appName = m.input.Value()
				m.input.Reset()
				m.input.Blur()
				m.stage = 1
				return m, nil
			} else if m.stage == 1 {
				m.region = regions[m.cursor]
				m.stage = 2
				return m, nil
			} else if m.stage == 2 {
				log.Println("Executing update operation...")
				err2 := source.UpdateRegistry(m.region, m.appName)
				if err2 != nil {
					log.Printf("Update failed: %v", err2)
					m.successMsg = fmt.Sprintf("Error: invalid app name to update %s to the %s registry\n %s", m.appName, m.region, err2)
				} else {
					log.Println("Update succeeded")
					m.successMsg = fmt.Sprintf("Changed %s to region %s successfully", m.appName, m.region)
				}
				m.stage = 3
				return m, nil
			} else if m.stage == 3 {
				return GetCommand(mainMenuName)
			}
		case "up", "k":
			if m.stage == 1 && m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.stage == 1 && m.cursor < len(regions)-1 {
				m.cursor++
			}
		}
	}

	if m.stage == 0 {
		m.input, cmd = m.input.Update(msg)
	}

	return m, cmd
}

func (m changeNameRegionModel) View() string {
	switch m.stage {
	case 0:
		return fmt.Sprintf(
			"Change Name to Region\n\n%s\n\nPress 'enter' to confirm, 'q' to go back.\n",
			m.input.View(),
		)
	case 1:
		s := fmt.Sprintf("Change Name to Region\n\nApp Name: %s\n\nSelect a region:\n\n", m.appName)
		for i, region := range regions {
			cursor := " " // no cursor
			line := region
			if m.cursor == i {
				cursor = ">" // cursor!
				boldBlue := color.New(color.FgBlue).Add(color.Bold)
				line = boldBlue.Sprintf(region)
			}
			s += fmt.Sprintf("%s %s\n", cursor, line)
		}
		s += "\nPress 'enter' to confirm, 'q' to go back.\n"
		return s
	case 2:
		return fmt.Sprintf(
			"Change Name to Region\n\nApp Name: %s\nRegion: %s\n\nPress 'enter' to apply change, 'q' to go back.\n",
			m.appName,
			m.region,
		)
	case 3:
		return fmt.Sprintf(
			"%s\n\nPress 'enter' to return to main menu.\n",
			m.successMsg,
		)
	default:
		return "Unexpected stage."
	}
}

func init() {
	RegisterCommand("changeNameRegion", "Change a App to Region", newChangeNameRegionModel())
}
