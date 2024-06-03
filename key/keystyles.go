package key

import "github.com/charmbracelet/lipgloss"

var (
	validChar = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA"))
	// CurrChar     = lipgloss.NewStyle().Background(lipgloss.Color("#7D56F4")).Foreground(lipgloss.Color("#FAFAFA"))
	notTypedChar = lipgloss.NewStyle().Foreground(lipgloss.Color("#808080"))
	mistakes     = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0070"))
)
