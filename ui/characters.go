package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/hkionline/dftui/models"
)

// renderCharactersTab renders the Characters tab content
func (m Model) renderCharactersTab() string {
	if m.err != nil {
		return fmt.Sprintf("Error loading characters: %v", m.err)
	}

	if len(m.characters) == 0 {
		return "No characters found"
	}

	var lines []string
	lines = append(lines, lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("Characters for %s:", m.username)))
	lines = append(lines, "")

	// Render each character with PC/NPC indicator
	for _, char := range m.characters {
		line := renderCharacter(char)
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}

// renderCharacter renders a single character line with type indicator
func renderCharacter(char models.Character) string {
	var typeStyle lipgloss.Style
	var typeLabel string

	if char.Type == models.PC {
		// Player Character - highlight in green
		typeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("10")).
			Bold(true)
		typeLabel = "[PC]"
	} else {
		// Non-Player Character - show in yellow
		typeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("11")).
			Bold(true)
		typeLabel = "[NPC]"
	}

	nameStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("15"))

	return fmt.Sprintf("  %s %s",
		typeStyle.Render(typeLabel),
		nameStyle.Render(char.Name))
}
