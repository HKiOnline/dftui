package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hkionline/dftui/models"
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
	username   string
	activeTab  Tab
	characters []models.Character
	backend    services.Backend
	err        error
	width      int
	height     int
}

// NewModel creates a new UI model
func NewModel(username string, backend services.Backend) Model {
	return Model{
		username:  username,
		activeTab: TabCharacters,
		backend:   backend,
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

	help := "Tab/→: Next | Shift+Tab/←: Previous | 1-5: Jump to tab | q: Quit"
	return helpStyle.Render(help)
}

// charactersLoadedMsg is sent when characters are loaded from backend
type charactersLoadedMsg struct {
	characters []models.Character
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
