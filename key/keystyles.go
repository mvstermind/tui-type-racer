package key

import "github.com/charmbracelet/lipgloss"

var (
	ValidChar = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA"))
	// CurrChar     = lipgloss.NewStyle().Background(lipgloss.Color("#7D56F4")).Foreground(lipgloss.Color("#FAFAFA"))
	NotTypedChar = lipgloss.NewStyle().Foreground(lipgloss.Color("#808080"))
)
