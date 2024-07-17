package cmd

import (
	"regtool/source"
	"strings"
	"time"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
)

type listAllRegistryModel struct {
	output     []string
	done       bool
	messagesCh chan string
	scroll     int
	height     int
	width      int
}

func (m listAllRegistryModel) Init() tea.Cmd {
	return tea.Batch(
		m.startListingRegistries(),
		tickEvery(),
	)
}

func (m listAllRegistryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return GetCommand(mainMenuName)
		case "up", "k":
			if m.scroll > 0 {
				m.scroll--
			}
		case "down", "j":
			if m.scroll < len(m.output)-m.height {
				m.scroll++
			}
		case "r":
			if m.done {
				m.output = nil
				m.done = false
				m.messagesCh = make(chan string)
				m.scroll = 0
				return m, tea.Batch(
					m.startListingRegistries(),
					tickEvery(),
				)
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
	return m, nil
}

func (m listAllRegistryModel) View() string {
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
		}
	}

	var finalBuilder strings.Builder
	finalBuilder.WriteString(GetStyledTitle("Registry List") + "\n\n")
	finalBuilder.WriteString(contentBuilder.String())
	finalBuilder.WriteString("\n")

	if m.done {
		finalBuilder.WriteString(GetStyledQuitText() + "\n")
		finalBuilder.WriteString(GetInfoText("Press 'r' to reload, use up/down arrows to scroll") + "\n")
	} else {
		finalBuilder.WriteString(GetInfoText("Loading... Use j/k arrows to scroll") + "\n")
	}

	return finalBuilder.String()
}

func (m listAllRegistryModel) formatApp(appLine string, allLines []string) string {
	var builder strings.Builder
	builder.WriteString(GetStyledTitle(appLine) + "\n")
	inCurrentApp := false
	for _, line := range allLines {
		if line == appLine {
			inCurrentApp = true
			continue
		}
		if inCurrentApp {
			if strings.HasPrefix(line, "APP: ") {
				break
			}
			if strings.HasPrefix(line, "  REGION: ") {
				parts := strings.SplitN(line, ", URL: ", 2)
				builder.WriteString(m.wrapText(GetInfoText(parts[0]), m.width) + "\n")
				if len(parts) > 1 {
					builder.WriteString(m.wrapText(GetSuccessText("URL: "+parts[1]), m.width) + "\n")
				}
			} else if strings.HasPrefix(line, "ERROR: ") {
				builder.WriteString(m.wrapText(GetErrorText(line), m.width) + "\n")
			}
		}
	}
	builder.WriteString("\n")
	return builder.String()
}

func (m listAllRegistryModel) wrapText(text string, width int) string {
	words := strings.Fields(removeColorAttributes(text))
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
		styledLines[i] = applyOriginalStyle(text, line)
	}

	return strings.Join(styledLines, "\n    ")
}

func removeColorAttributes(s string) string {
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

func applyOriginalStyle(original, wrapped string) string {
	if strings.HasPrefix(original, GetInfoText("")) {
		return GetInfoText(wrapped)
	} else if strings.HasPrefix(original, GetSuccessText("")) {
		return GetSuccessText(wrapped)
	} else if strings.HasPrefix(original, GetErrorText("")) {
		return GetErrorText(wrapped)
	}
	return wrapped
}

func (m listAllRegistryModel) startListingRegistries() tea.Cmd {
	return func() tea.Msg {
		go source.ListAllRegistry(m.messagesCh)
		return nil
	}
}

func tickEvery() tea.Cmd {
	return tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func NewListAllRegistryModel() listAllRegistryModel {
	return listAllRegistryModel{
		messagesCh: make(chan string),
		height:     20,
		width:      80,
	}
}

func init() {
	RegisterCommand("listAllRegistry", "List All Registry", NewListAllRegistryModel())
}
