package cmd

import (
	"regtool/source"
	"strings"
	"unicode"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type listRegistryModel struct {
	appNameInput textinput.Model
	appName      string
	stage        int
	output       []string
	done         bool
	messagesCh   chan string
	scroll       int
	height       int
	width        int
}

func (m listRegistryModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m listRegistryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return GetCommand("mainMenu")
		case "up", "k":
			if m.scroll > 0 {
				m.scroll--
			}
		case "down", "j":
			if m.scroll < len(m.output)-m.height {
				m.scroll++
			}
		default:
			if m.stage == 0 {
				var cmd tea.Cmd
				m.appNameInput, cmd = m.appNameInput.Update(msg)
				cmds = append(cmds, cmd)

				if msg.Type == tea.KeyEnter {
					m.appName = m.appNameInput.Value()
					m.stage = 1
					m.output = nil
					m.done = false
					m.messagesCh = make(chan string)
					m.scroll = 0
					return m, tea.Batch(
						m.startListingRegistryByAppName(m.appName),
						tickEvery(),
					)
				}
			} else {
				return GetCommand("mainMenu")
			}
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height - 6
		m.width = msg.Width - 4
	case tickMsg:
		select {
		case message, ok := <-m.messagesCh:
			if !ok {
				m.done = true
				return m, nil
			}
			m.output = append(m.output, message)
			if len(m.output) > m.height && m.scroll == len(m.output)-m.height-1 {
				m.scroll++
			}
		default:
			// No message available, do nothing
		}
		return m, tickEvery()
	}
	return m, tea.Batch(cmds...)
}

func (m listRegistryModel) View() string {
	if m.stage == 0 {
		return "List Registry by App Name\n\nEnter the app name:\n" + m.appNameInput.View()
	}

	var contentBuilder strings.Builder

	visibleOutput := m.output
	if len(m.output) > m.height {
		start := m.scroll
		end := m.scroll + m.height
		if end > len(m.output) {
			end = len(m.output)
		}
		visibleOutput = m.output[start:end]
	}

	for _, line := range visibleOutput {
		if strings.HasPrefix(line, "APP: ") {
			contentBuilder.WriteString("\n" + m.formatApp(line, m.output))
		} else {
			contentBuilder.WriteString(line + "\n")
		}
	}

	var finalBuilder strings.Builder
	finalBuilder.WriteString(GetStyledTitle("Registry List") + "\n\n")
	finalBuilder.WriteString(contentBuilder.String())
	finalBuilder.WriteString("\n")

	if m.done {
		finalBuilder.WriteString(GetStyledQuitText() + "\n")
		finalBuilder.WriteString(GetInfoText("Press 'q' to go back, use up/down arrows to scroll") + "\n")
	} else {
		finalBuilder.WriteString(GetInfoText("Loading... Use j/k arrows to scroll") + "\n")
	}

	return finalBuilder.String()
}
func (m listRegistryModel) formatApp(appLine string, allLines []string) string {
	var builder strings.Builder
	inCurrentApp := false

	for _, line := range allLines {
		if strings.HasPrefix(line, "APP: ") {
			if inCurrentApp {
				break
			}
			inCurrentApp = true
			builder.WriteString(GetStyledTitle(line) + "\n")
			continue
		}
	}
	builder.WriteString("\n")
	return builder.String()
}

func (m listRegistryModel) wrapText(text string, width int) string {
	words := strings.Fields(removeSingleColorAttributes(text))
	if len(words) == 0 {
		return ""
	}

	var lines []string
	var currentLine string

	for _, word := range words {
		if len(currentLine)+len(word)+1 > width {
			lines = append(lines, strings.TrimSpace(currentLine))
			currentLine = word
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}

	if currentLine != "" {
		lines = append(lines, strings.TrimSpace(currentLine))
	}

	styledLines := make([]string, len(lines))
	for i, line := range lines {
		styledLines[i] = applySingleOriginalStyle(text, line)
	}

	return strings.Join(styledLines, "\n    ")
}

func (m listRegistryModel) startListingRegistryByAppName(appName string) tea.Cmd {
	return func() tea.Msg {
		go source.ListRegistryByAppName(appName, m.messagesCh)
		return nil
	}
}

func NewListRegistryModel() listRegistryModel {
	ti := textinput.New()
	ti.Placeholder = "App Name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return listRegistryModel{
		appNameInput: ti,
		messagesCh:   make(chan string),
		height:       20,
		width:        80,
	}
}

func init() {
	RegisterCommand("listRegistry", "List Registry by App Name", NewListRegistryModel())
}

func removeSingleColorAttributes(s string) string {
	var result strings.Builder
	inEscapeSeq := false
	for _, r := range s {
		if r == '\x1b' {
			inEscapeSeq = true
		} else if inEscapeSeq {
			if unicode.IsLetter(r) {
				inEscapeSeq = false
			}
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func applySingleOriginalStyle(original, wrapped string) string {
	if strings.HasPrefix(original, GetInfoText("")) {
		return GetInfoText(wrapped)
	} else if strings.HasPrefix(original, GetSuccessText("")) {
		return GetSuccessText(wrapped)
	} else if strings.HasPrefix(original, GetErrorText("")) {
		return GetErrorText(wrapped)
	}
	return wrapped
}
