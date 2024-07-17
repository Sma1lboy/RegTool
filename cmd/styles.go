package cmd

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	primaryColor    = lipgloss.Color("#61AFEF")
	bgColor         = lipgloss.Color("#282C34")
	textColor       = lipgloss.Color("#ABB2BF")
	selectedColor   = lipgloss.Color("#98C379")
	unselectedColor = lipgloss.Color("#3E4451")
	accentColor     = lipgloss.Color("#C678DD")

	// Styles
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

	// Additional shared styles
	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B6B")).
			Bold(true)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4CAF50")).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#56B6C2"))

	// Helper function to create a bordered box
	borderedBox = func(s string) string {
		return lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(1).
			Render(s)
	}
)

// GetStyledTitle returns a styled title for a given command
func GetStyledTitle(title string) string {
	return titleStyle.Render("🔧 " + title)
}

// GetStyledOption returns a styled option for menus
func GetStyledOption(option string, selected bool) string {
	style := optionStyle
	cursor := "  "
	if selected {
		style = selectedOptionStyle
		cursor = "▶ "
	}
	return style.Render(cursor + option)
}

// GetStyledQuitText returns the styled quit instruction
func GetStyledQuitText() string {
	return quitTextStyle.Render("Press 'q' to quit")
}

// GetErrorText returns styled error text
func GetErrorText(text string) string {
	return errorStyle.Render(text)
}

// GetSuccessText returns styled success text
func GetSuccessText(text string) string {
	return successStyle.Render(text)
}

// GetInfoText returns styled info text
func GetInfoText(text string) string {
	return infoStyle.Render(text)
}

// GetBorderedBox returns a bordered box containing the given text
func GetBorderedBox(text string) string {
	return borderedBox(text)
}
