package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/hkionline/dftui/dflib/dfm"
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

	// Separate PCs and NPCs
	var pcs []dfm.Character
	var npcs []dfm.Character
	for _, char := range m.characters {
		if char.Group == string(dfm.PC) {
			pcs = append(pcs, char)
		} else if char.Group == string(dfm.NPC) {
			npcs = append(npcs, char)
		}
	}

	// Render Player Characters section
	if len(pcs) > 0 {
		lines = append(lines, lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10")).Render("Player Characters:"))
		for i, char := range pcs {
			line := renderCharacter(char, m.selectedCharacterIndex == i)
			lines = append(lines, line)
		}
		lines = append(lines, "")
	}

	// Render Non-Player Characters section
	if len(npcs) > 0 {
		lines = append(lines, lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("11")).Render("Non-Player Characters:"))
		startIndex := len(pcs)
		for i, char := range npcs {
			line := renderCharacter(char, m.selectedCharacterIndex == startIndex+i)
			lines = append(lines, line)
		}
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
	if char.Group == string(dfm.PC) {
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

	// Basic information
	lines = append(lines, lipgloss.NewStyle().Bold(true).Render("Basic Information:"))
	lines = append(lines, fmt.Sprintf("%s %s",
		labelStyle.Render("Name:"),
		valueStyle.Render(char.Name)))

	if char.Player != "" {
		lines = append(lines, fmt.Sprintf("%s %s",
			labelStyle.Render("Player:"),
			valueStyle.Render(char.Player)))
	}

	lines = append(lines, fmt.Sprintf("%s %s",
		labelStyle.Render("Type:"),
		typeDisplay))

	// Spirit type and player information
	spiritDisplay := char.Spirit
	if spiritDisplay == string(dfm.SpiritVampire) {
		spiritDisplay = "Vampire"
	} else if spiritDisplay == string(dfm.SpiritGhoul) {
		spiritDisplay = "Ghoul"
	} else if spiritDisplay == string(dfm.SpiritHuman) {
		spiritDisplay = "Human"
	}
	lines = append(lines, fmt.Sprintf("%s %s",
		labelStyle.Render("Spirit:"),
		valueStyle.Render(spiritDisplay)))

	// Year information (embrace or setting year)
	if char.EmbraceYear != 0 {
		lines = append(lines, fmt.Sprintf("%s %s",
			labelStyle.Render("Embrace:"),
			valueStyle.Render(fmt.Sprintf("%d", char.EmbraceYear))))
	}
	if char.SettingYear != 0 && char.SettingYear != char.EmbraceYear {
		lines = append(lines, fmt.Sprintf("%s %s",
			labelStyle.Render("Setting:"),
			valueStyle.Render(fmt.Sprintf("%d", char.SettingYear))))
	}

	// Fate section
	lines = append(lines, "")
	lines = append(lines, lipgloss.NewStyle().Bold(true).Render("Fate:"))
	lines = append(lines, fmt.Sprintf("%s %s",
		labelStyle.Render("Refresh:"),
		valueStyle.Render(fmt.Sprintf("%d", char.Refresh))))
	lines = append(lines, fmt.Sprintf("%s %s",
		labelStyle.Render("Points:"),
		valueStyle.Render(fmt.Sprintf("%d", char.FatePoint))))

	// Aspects
	if len(char.Aspects) > 0 {
		lines = append(lines, "")
		lines = append(lines, lipgloss.NewStyle().Bold(true).Render("Aspects:"))
		for _, aspect := range char.Aspects {
			aspectType := aspect.Type
			if aspectType == "high concept" {
				aspectType = "High Concept"
			} else if aspectType == "trouble" {
				aspectType = "Trouble"
			} else if aspectType == "clan" {
				aspectType = "Clan"
			} else if aspectType == "covenant" {
				aspectType = "Covenant"
			}
			lines = append(lines, fmt.Sprintf("  %s: %s", aspectType, aspect.Title))
			if aspect.Description != "" {
				lines = append(lines, fmt.Sprintf("    %s", aspect.Description))
			}
		}
	}

	// Beast and Blood section (vampire only) - moved after Aspects
	if char.Spirit == string(dfm.SpiritVampire) {
		lines = append(lines, "")
		lines = append(lines, lipgloss.NewStyle().Bold(true).Render("Beast and Blood:"))
		if char.BloodPotency > 0 {
			lines = append(lines, fmt.Sprintf("Blood Potency %d", char.BloodPotency))
		}
		// Disciplines moved here from Stunts section - now on one line
		if len(char.Disciplines) > 0 {
			var disciplineLines []string
			for _, disc := range char.Disciplines {
				if disc.Rating > 0 {
					// Capitalize discipline names
					capitalizedTitle := strings.Title(disc.Title)
					disciplineLines = append(disciplineLines, fmt.Sprintf("%s %d", capitalizedTitle, disc.Rating))
				}
			}
			if len(disciplineLines) > 0 {
				lines = append(lines, fmt.Sprintf("Disciplines: %s", strings.Join(disciplineLines, ", ")))
			}
		}
	}

	// Skills (grouped by category)
	if len(char.Skills) > 0 {
		lines = append(lines, "")
		lines = append(lines, lipgloss.NewStyle().Bold(true).Render("Skills:"))

		mental := []string{}
		physical := []string{}
		social := []string{}
		for _, skill := range char.Skills {
			if skill.Rating > 0 {
				// Capitalize skill name and add space before rating
				capitalizedTitle := strings.Title(skill.Title)
				line := fmt.Sprintf("%s %d", capitalizedTitle, skill.Rating)
				if skill.Group == "mental" {
					mental = append(mental, line)
				} else if skill.Group == "physical" {
					physical = append(physical, line)
				} else if skill.Group == "social" {
					social = append(social, line)
				}
			}
		}

		if len(mental) > 0 {
			lines = append(lines, fmt.Sprintf("  Mental: %s", strings.Join(mental, ", ")))
		}
		if len(physical) > 0 {
			lines = append(lines, fmt.Sprintf("  Physical: %s", strings.Join(physical, ", ")))
		}
		if len(social) > 0 {
			lines = append(lines, fmt.Sprintf("  Social: %s", strings.Join(social, ", ")))
		}
	}

	// Stunts
	if len(char.Stunts) > 0 {
		lines = append(lines, "")
		lines = append(lines, lipgloss.NewStyle().Bold(true).Render("Stunts:"))
		for _, stunt := range char.Stunts {
			if stunt.Title != "" {
				lines = append(lines, fmt.Sprintf("  %s", stunt.Title))
			}
			if stunt.Description != "" {
				lines = append(lines, fmt.Sprintf("    %s", stunt.Description))
			}
		}
	}

	// Stress tracks
	lines = append(lines, "")
	lines = append(lines, lipgloss.NewStyle().Bold(true).Render("Stress Tracks:"))
	lines = append(lines, fmt.Sprintf("  Physical: %d/%d", char.PhysicalStressCurrent, char.PhysicalStressLimit))
	lines = append(lines, fmt.Sprintf("  Mental: %d/%d", char.MentalStressCurrent, char.MentalStressLimit))
	if char.Spirit == string(dfm.SpiritVampire) && char.HungerStressLimit > 0 {
		lines = append(lines, fmt.Sprintf("  Hunger: %d/%d", char.HungerStressCurrent, char.HungerStressLimit))
	}

	// Consequences
	if len(char.Consequences) > 0 {
		lines = append(lines, "")
		lines = append(lines, lipgloss.NewStyle().Bold(true).Render("Consequences:"))
		for _, cons := range char.Consequences {
			status := "Inactive"
			if cons.IsActive {
				status = "Active"
			}
			lines = append(lines, fmt.Sprintf("  Level %d: %s", cons.Level, status))
			if cons.Title != "" && cons.IsActive {
				lines = append(lines, fmt.Sprintf("    %s", cons.Title))
			}
		}
	}

	// Notes section
	if char.Notes != "" {
		lines = append(lines, "")
		lines = append(lines, lipgloss.NewStyle().Bold(true).Render("Notes:"))
		lines = append(lines, valueStyle.Render(char.Notes))
	}

	lines = append(lines, "")
	lines = append(lines, lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Render("Press ESC to return to character list"))

	return strings.Join(lines, "\n")
}

// renderCharacter renders a single character line with description and optional selection highlight
func renderCharacter(char dfm.Character, isSelected bool) string {
	nameStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("15"))

	// Add selection indicator and highlight
	cursor := "  "
	if isSelected {
		cursor = "> "
		// Highlight selected character with background color
		nameStyle = nameStyle.Background(lipgloss.Color("237"))
	}

	// Prepare description - truncate to 50 characters if needed
	description := char.Description
	if len(description) > 50 {
		description = description[:50] + "..."
	}

	// Format: cursor + name + " - " + description
	var result string
	if description != "" {
		result = fmt.Sprintf("%s%s - %s",
			cursor,
			nameStyle.Render(char.Name),
			lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render(description))
	} else {
		result = fmt.Sprintf("%s%s",
			cursor,
			nameStyle.Render(char.Name))
	}

	return result
}
