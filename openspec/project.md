# Project Context

## Purpose
Dark Fate Terminal UI (dftui) is a terminal-based front-end for the Dark Fate RPG manager, accessible via SSH. The project aims to provide an intuitive and responsive terminal user interface for managing Dark Fate RPG sessions and data.

## Tech Stack
- **Language**: Go
- **UI Framework**: Charm Wish (SSH-based TUI framework)
- **Protocol**: SSH for remote access
- Additional libraries: [To be added as project develops]

## Project Conventions

### Code Style
- Follow standard Go formatting (`gofmt`, `goimports`)
- Use Go naming conventions:
  - Exported identifiers: PascalCase
  - Unexported identifiers: camelCase
  - Package names: lowercase, single word when possible
- Comment all exported functions, types, and packages
- Keep functions focused and testable
- Error handling: always check and handle errors explicitly

### Architecture Patterns
- **Tabbed TUI Interface**: Main navigation via tabs (Characters, Sessions, Chronicles, Campaigns, Fate Tracker)
- **MVC/Component-based TUI**: Separate UI components, business logic, and data models
- **SSH Session Management**: Handle multiple concurrent SSH connections
- **State Management**:
  - User authentication and session state
  - Active game session tracking (for session-mode dice roller)
  - Selected Chronicle/Campaign context
- **Multi-user Real-time Updates**: Shared game session data (dice rolls, turn tracking) visible to all participants
- **Data Rendering**:
  - Markdown renderer for READMEs and session notes
  - JSON-based interactive UI for Fate tracker
  - File download handling for resource blobs
- **Backend Integration**: Connect to Dark Fate RPG backend for:
  - User authentication
  - Character data (PCs and NPCs)
  - Chronicle/Campaign/Session data
  - File storage (READMEs, notes, resources)

### Testing Strategy
- Write unit tests for business logic and data models
- Use Go's built-in testing framework (`testing` package)
- Aim for reasonable test coverage on critical paths
- Integration tests for SSH connectivity and session handling
- [Consider table-driven tests for Go best practices]

### Git Workflow
- **Branching**: [Define your strategy - e.g., main/develop, feature branches]
- **Commits**: Use clear, descriptive commit messages
  - Format: `<type>: <description>` (e.g., "feat: add character sheet view")
  - Types: feat, fix, refactor, docs, test, chore
- **Pull Requests**: [Define your PR process if using]

## Domain Context

### Data Hierarchy
The application follows a three-level hierarchy:
- **Chronicle** → **Campaign** → **Session**

**Chronicle**:
- Story setting that includes playable campaigns for specific characters
- More fixed/stable than campaigns
- Contains:
  - README (markdown) - explains the chronicle setting
  - Downloadable resource files (PDFs and other blobs)

**Campaign**:
- Part of a specific Chronicle
- Has specific characters, both:
  - Player Characters (PCs) - user's characters
  - Non-Player Characters (NPCs) - gamemaster's characters
- Contains:
  - README (markdown) - campaign information
  - Downloadable resource files (PDFs and other blobs)
  - List of gaming sessions

**Session**:
- Gaming session set in a specific campaign
- Contains:
  - Session notes (markdown format)
  - Fate session tracker and dice roller results (JSON format)
  - Records of who's turn it is and dice roll history

### UI Views and Features
The terminal UI has a tabbed interface with the following views:

1. **Characters Tab**
   - List of characters visible to the logged-in user
   - Each character marked as PC (player character) or NPC (non-player character)

2. **Sessions Tab**
   - List of active game sessions currently available
   - Access to session-specific features

3. **Chronicles Tab**
   - List of available Chronicles
   - Access to selected Chronicle's README (markdown display)
   - List of resource files available for download

4. **Campaigns Tab**
   - List of campaigns
   - Access to selected campaign's README (markdown display)
   - List of game sessions for the selected campaign
   - Session notes display (markdown)
   - Fate session tracker and dice roller results
   - List of campaign resource files available for download

5. **Fate Session Tracker & Dice Roller**
   - **Single-mode**: Rolls not stored (standalone use)
   - **Session-mode**: Rolls visible to anyone in the same game session (requires session selection)
   - Tracks whose turn it is in the game
   - Records and displays dice roll results

### Data Storage Formats
- **READMEs**: Markdown files, displayed as formatted text in TUI
- **Session notes**: Markdown files, displayed as formatted text in TUI
- **Fate tracker data**: JSON files, rendered as interactive TUI view
- **Resource files**: Binary blobs (PDFs, etc.), downloadable but not displayed in TUI

### Terminal UI Paradigm
- Users interact via SSH, expecting responsive keyboard-driven navigation
- Session Management: Multiple users may connect simultaneously
- Real-time updates for shared game sessions (dice rolls, turn tracking)

## Important Constraints
- Must run efficiently in terminal environments with limited rendering capabilities
- SSH protocol constraints (latency, connection handling, security)
- Cross-platform compatibility (Linux, macOS, Windows terminals)
- Performance: responsive UI even over slow network connections
- [Add any specific performance or security requirements]

## External Dependencies
- **Charm Wish**: SSH framework for building terminal applications
- **Dark Fate RPG Backend**: [Document API endpoints, authentication, data formats]
- [Add other Charm libraries if using - Bubble Tea, Lip Gloss, etc.]
- [Add any database or storage systems]
