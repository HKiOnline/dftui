package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hkionline/dftui/dflib/dfm"
	"github.com/hkionline/dftui/services"
)

// Tab represents a tab in the UI
type Tab int

const (
	TabCharacters Tab = iota
	TabSessions
	TabChronicles
	TabCampaigns
	TabFateTracker
)

// CharacterViewMode represents the current view mode in the Characters tab
type CharacterViewMode int

const (
	CharacterViewList CharacterViewMode = iota
	CharacterViewDetail
)

// TabInfo holds display information for tabs
type TabInfo struct {
	Name string
	Tab  Tab
}

var tabs = []TabInfo{
	{Name: "Characters", Tab: TabCharacters},
	{Name: "Sessions", Tab: TabSessions},
	{Name: "Chronicles", Tab: TabChronicles},
	{Name: "Campaigns", Tab: TabCampaigns},
	{Name: "Fate Tracker", Tab: TabFateTracker},
}

// Model is the main Bubble Tea model for the application
// It follows the Elm architecture: Model -> Update -> View
// See: https://github.com/charmbracelet/bubbletea
type Model struct {
	username               string
	activeTab              Tab
	characters             []dfm.Character
	backend                services.Backend
	err                    error
	width                  int
	height                 int
	selectedCharacterIndex int               // Index of currently selected character in list (0-based, -1 if none)
	characterViewMode      CharacterViewMode // Current view mode in Characters tab (list or detail)
	selectedCharacter      *dfm.Character    // Currently selected character for detail view
}

// NewModel creates a new UI model
func NewModel(username string, backend services.Backend) Model {
	return Model{
		username:               username,
		activeTab:              TabCharacters,
		backend:                backend,
		selectedCharacterIndex: 0,                 // Start with first character selected
		characterViewMode:      CharacterViewList, // Start in list view
		selectedCharacter:      nil,               // No character selected initially
	}
}

// Init initializes the model (Bubble Tea lifecycle method)
func (m Model) Init() tea.Cmd {
	// Load user's characters
	return loadCharacters(m.username, m.backend)
}

// Update handles messages and updates the model (Bubble Tea lifecycle method)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			// Quit the application
			return m, tea.Quit

		case "tab", "right":
			// Navigate to next tab
			m.activeTab = (m.activeTab + 1) % Tab(len(tabs))
			return m, nil

		case "shift+tab", "left":
			// Navigate to previous tab
			m.activeTab = (m.activeTab - 1 + Tab(len(tabs))) % Tab(len(tabs))
			return m, nil

		case "1":
			m.activeTab = TabCharacters
			return m, nil
		case "2":
			m.activeTab = TabSessions
			return m, nil
		case "3":
			m.activeTab = TabChronicles
			return m, nil
		case "4":
			m.activeTab = TabCampaigns
			return m, nil
		case "5":
			m.activeTab = TabFateTracker
			return m, nil

		case "up":
			// Navigate up in character list (only in Characters tab, list view)
			if m.activeTab == TabCharacters && m.characterViewMode == CharacterViewList && len(m.characters) > 0 {
				if m.selectedCharacterIndex > 0 {
					m.selectedCharacterIndex--
				}
			}
			return m, nil

		case "down":
			// Navigate down in character list (only in Characters tab, list view)
			if m.activeTab == TabCharacters && m.characterViewMode == CharacterViewList && len(m.characters) > 0 {
				if m.selectedCharacterIndex < len(m.characters)-1 {
					m.selectedCharacterIndex++
				}
			}
			return m, nil

		case "enter":
			// Select character and switch to detail view (only in Characters tab, list view)
			if m.activeTab == TabCharacters && m.characterViewMode == CharacterViewList && len(m.characters) > 0 {
				if m.selectedCharacterIndex >= 0 && m.selectedCharacterIndex < len(m.characters) {
					m.selectedCharacter = &m.characters[m.selectedCharacterIndex]
					m.characterViewMode = CharacterViewDetail
				}
			}
			return m, nil

		case "esc":
			// Return to list view from detail view (only in Characters tab, detail view)
			if m.activeTab == TabCharacters && m.characterViewMode == CharacterViewDetail {
				m.characterViewMode = CharacterViewList
				m.selectedCharacter = nil
			}
			return m, nil
		}

	case tea.WindowSizeMsg:
		// Handle terminal resize
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case charactersLoadedMsg:
		// Characters loaded from backend
		m.characters = msg.characters
		m.err = msg.err
		// Select first character if list is not empty
		if len(m.characters) > 0 {
			m.selectedCharacterIndex = 0
		} else {
			m.selectedCharacterIndex = -1
		}
		return m, nil
	}

	return m, nil
}

// View renders the UI (Bubble Tea lifecycle method)
func (m Model) View() string {
	if m.width == 0 {
		// Terminal size not yet known
		return "Loading..."
	}

	// Render tab bar
	tabBar := m.renderTabBar()

	// Render content based on active tab
	content := m.renderContent()

	// Render help text
	help := m.renderHelp()

	return fmt.Sprintf("%s\n\n%s\n\n%s", tabBar, content, help)
}

// renderTabBar renders the tab navigation bar
func (m Model) renderTabBar() string {
	var renderedTabs []string

	activeStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("15")).
		Background(lipgloss.Color("63")).
		Padding(0, 2)

	inactiveStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("245")).
		Padding(0, 2)

	for _, tab := range tabs {
		if tab.Tab == m.activeTab {
			renderedTabs = append(renderedTabs, activeStyle.Render(tab.Name))
		} else {
			renderedTabs = append(renderedTabs, inactiveStyle.Render(tab.Name))
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
}

// renderContent renders the content area based on the active tab
func (m Model) renderContent() string {
	contentStyle := lipgloss.NewStyle().
		Width(m.width-4).
		Height(m.height-10).
		Padding(1, 2)

	var content string

	switch m.activeTab {
	case TabCharacters:
		content = m.renderCharactersTab()
	case TabSessions:
		content = "Sessions tab - Not yet implemented"
	case TabChronicles:
		content = "Chronicles tab - Not yet implemented"
	case TabCampaigns:
		content = "Campaigns tab - Not yet implemented"
	case TabFateTracker:
		content = "Fate Tracker tab - Not yet implemented"
	}

	return contentStyle.Render(content)
}

// renderHelp renders the help text at the bottom
func (m Model) renderHelp() string {
	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Padding(0, 2)

	var help string
	// Context-sensitive help based on current tab and view mode
	if m.activeTab == TabCharacters {
		if m.characterViewMode == CharacterViewList {
			help = "↑/↓: Navigate | Enter: View Details | Tab/→: Next | Shift+Tab/←: Previous | 1-5: Jump to tab | q: Quit"
		} else if m.characterViewMode == CharacterViewDetail {
			help = "ESC: Back to List | Tab/→: Next | Shift+Tab/←: Previous | 1-5: Jump to tab | q: Quit"
		}
	} else {
		help = "Tab/→: Next | Shift+Tab/←: Previous | 1-5: Jump to tab | q: Quit"
	}

	return helpStyle.Render(help)
}

// charactersLoadedMsg is sent when characters are loaded from backend
type charactersLoadedMsg struct {
	characters []dfm.Character
	err        error
}

// loadCharacters loads characters from the backend
func loadCharacters(username string, backend services.Backend) tea.Cmd {
	return func() tea.Msg {
		characters, err := backend.GetUserCharacters(username)
		return charactersLoadedMsg{
			characters: characters,
			err:        err,
		}
	}
}
