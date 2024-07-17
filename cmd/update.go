package cmd

import (
	"fmt"
	"regtool/source"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// updateModel defines the structure for the update model
type updateModel struct {
	confirming     bool
	completed      bool
	progress       progress.Model
	progressVal    float64
	spinner        spinner.Model
	msg            string
	updateMessages []string    // Slice to store all update messages
	updateChan     chan string // Channel to receive update messages
}

// Init initializes the update model
func (m updateModel) Init() tea.Cmd {
	return m.spinner.Tick
}

// Update handles the messages for the update model
func (m updateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if !m.confirming {
				m.confirming = true
				m.progressVal = 0

				updateResult := make(chan error)
				m.updateChan = make(chan string)
				go func() {
					err := source.Update(m.updateChan)
					updateResult <- err
					close(m.updateChan)
				}()

				return m, tea.Batch(animateProgress(), waitForUpdate(updateResult), waitForUpdateMessages(m.updateChan))
			}
		case "q", "esc":
			return GetCommand(mainMenuName)
		}
	case tickMsg:
		m.progressVal += 0.01
		cmd := m.progress.IncrPercent(0.01)
		if m.progressVal >= 1.0 {
			if m.completed {
				return m, tea.Batch(cmd, func() tea.Msg { return updateCompleteMsg{} })
			}
			m.progressVal = 1.0
		}
		var spinCmd tea.Cmd
		m.spinner, spinCmd = m.spinner.Update(msg)
		return m, tea.Batch(animateProgress(), cmd, spinCmd)

	case updateCompleteMsg:
		return GetCommand(mainMenuName)

	case updateResultMsg:
		if msg.err != nil {
			m.msg = msg.err.Error()
			m.confirming = false
		} else {
			m.progress.SetPercent(1.0)
			m.progressVal = .85
			m.completed = true
			m.msg = "Update completed successfully!"
		}
		return m, nil

	case updateMessageMsg:
		m.updateMessages = append(m.updateMessages, msg.message) // Store the message
		return m, tea.Batch(waitForUpdateMessages(m.updateChan)) // Continue listening for more messages
	}

	var cmd tea.Cmd
	var model tea.Model
	model, cmd = m.progress.Update(msg)
	m.progress = model.(progress.Model)
	var spinCmd tea.Cmd
	m.spinner, spinCmd = m.spinner.Update(msg)
	return m, tea.Batch(cmd, spinCmd)
}

// View renders the view for the update model
func (m updateModel) View() string {
	doc := strings.Builder{}
	// Title
	doc.WriteString(titleStyle.Render("🔧 Init/Update All Apps Recording") + "\n")

	if !m.confirming {
		doc.WriteString("Are you sure you want to check all apps and query local software? Press 'Enter' to confirm, 'q' or 'esc' to go back.\n")
		return doc.String()
	}

	doc.WriteString(m.spinner.View() + "Updating... Press 'q' or 'esc' to cancel.\n\n")
	doc.WriteString(fmt.Sprintf("%s\n\n", m.progress.View()))
	doc.WriteString("Update Messages:\n")
	for _, msg := range m.updateMessages {
		doc.WriteString(fmt.Sprintf(" - %s\n", msg))
	}
	doc.WriteString("\n" + m.msg)
	return doc.String()
}

// init initializes the update command with its model
func init() {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Align(lipgloss.Center)

	RegisterCommand("update", "Init/Update All Apps Recording", updateModel{
		progress: progress.New(progress.WithDefaultGradient()),
		spinner:  s,
	})
}

// updateCompleteMsg is used to signal that the update is complete
type updateCompleteMsg struct{}

// tickMsg is used to signal a tick for updating progress
type tickMsg time.Time

// updateResultMsg is used to signal the result of the update
type updateResultMsg struct {
	err error
}

// updateMessageMsg is used to send update messages
type updateMessageMsg struct {
	message string
}

// animateProgress sends a tick message at a fixed interval
func animateProgress() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// waitForUpdate waits for the update to complete and sends the result
func waitForUpdate(updateResult chan error) tea.Cmd {
	return func() tea.Msg {
		err := <-updateResult
		return updateResultMsg{err}
	}
}

// waitForUpdateMessages waits for update messages and sends them as tea messages
func waitForUpdateMessages(updateChan chan string) tea.Cmd {
	return func() tea.Msg {
		for msg := range updateChan {
			return updateMessageMsg{message: msg}
		}
		return nil
	}
}
