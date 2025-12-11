package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/hkionline/dftui/models"
)

// renderCharactersTab renders the Characters tab content
func (m Model) renderCharactersTab() string {
	// Switch between list and detail view based on current view mode
	if m.characterViewMode == CharacterViewDetail {
		return m.renderCharacterDetail()
	}

	// List view
	if m.err != nil {
		return fmt.Sprintf("Error loading characters: %v", m.err)
	}

	if len(m.characters) == 0 {
		return "No characters found"
	}

	var lines []string
	lines = append(lines, lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("Characters for %s:", m.username)))
	lines = append(lines, "")

	// Render each character with PC/NPC indicator and selection highlight
	for i, char := range m.characters {
		line := renderCharacter(char, i == m.selectedCharacterIndex)
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}

// renderCharacterDetail renders the detailed view of a selected character
func (m Model) renderCharacterDetail() string {
	if m.selectedCharacter == nil {
		return "No character selected"
	}

	char := m.selectedCharacter

	// Create styles for the detail view
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("15")).
		MarginBottom(1)

	labelStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("12")).
		Width(15)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("15"))

	// Determine type label and color
	var typeDisplay string
	if char.Type == models.PC {
		typeDisplay = lipgloss.NewStyle().
			Foreground(lipgloss.Color("10")).
			Bold(true).
			Render("Player Character (PC)")
	} else {
		typeDisplay = lipgloss.NewStyle().
			Foreground(lipgloss.Color("11")).
			Bold(true).
			Render("Non-Player Character (NPC)")
	}

	// Build the detail view
	var lines []string
	lines = append(lines, titleStyle.Render("Character Details"))
	lines = append(lines, "")
	lines = append(lines, fmt.Sprintf("%s %s",
		labelStyle.Render("Name:"),
		valueStyle.Render(char.Name)))
	lines = append(lines, fmt.Sprintf("%s %s",
		labelStyle.Render("ID:"),
		valueStyle.Render(char.ID)))
	lines = append(lines, fmt.Sprintf("%s %s",
		labelStyle.Render("Type:"),
		typeDisplay))
	lines = append(lines, "")
	lines = append(lines, lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Render("Press ESC to return to character list"))

	return strings.Join(lines, "\n")
}

// renderCharacter renders a single character line with type indicator and optional selection highlight
func renderCharacter(char models.Character, isSelected bool) string {
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

	// Add selection indicator and highlight
	cursor := "  "
	if isSelected {
		cursor = "> "
		// Highlight selected character with background color
		nameStyle = nameStyle.Background(lipgloss.Color("237"))
	}

	return fmt.Sprintf("%s%s %s",
		cursor,
		typeStyle.Render(typeLabel),
		nameStyle.Render(char.Name))
}
