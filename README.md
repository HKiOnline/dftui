# Dark Fate Terminal UI

Dark Fate Terminal UI is a front-end for Dark Fate RPG manager accessible via SSH. It is built in Go language and uses the wonderful Charm Wish -framework.


## Features

- SSH server with user identification
- Tabbed interface with keyboard navigation
- Characters tab displaying PCs and NPCs (with mock data)
- Placeholder tabs for Sessions, Chronicles, Campaigns, and Fate Tracker

## Current Status

This is the initial UI skeleton implementation. The SSH server and tabbed interface are functional, but only the Characters tab has content (using stub backend data). Other tabs are placeholders awaiting implementation.

## TODO

- Starter terminal UI
    - Tab for characters (logged in users characters and NPCs visible to the user)
        - List of characters, marked either PC (player character) or NPC (non-player character)
    - Tab for sessions (active game sessions available)
        - List of active game sessions
    - Tab for Chronicles (game setting for a campaings)
        - Access to selected chronicle's README
        - List of resource files that can be downloaded
    - Tab for Campaings
        - List of campaings
            - Access to selected campaing's README
            - List of game sessions of the selected campaing
                - Session notes
                - Fate sssion tracker and dice roller results
            - List of campaing resource files that can be downloaded
    - Access to Fate session tracker and dice roller (tracks who's turn it is and allows dice rolls)
        - Single-mode (rolls not stored)
        - Session-mode (rolls visible to anyone in a same game session, session must be selected)


## Data and conceptual hierarchies

Chronicle -> Campaings -> Sessions.

Chronicle is a story setting which includes playaple campaings for specific characters. Campaings are played in sessions. Chronicle is more fixed. It can have a README to explain what it is and downloadable content such as PDF-files. Campaings are part of some Chronicle. 
Campaing have specific character, both player (user's characters) and non-player characters (admin i.e. gamemaster's characters). Campaing can have a README and downloadable content such as PDF-files.
Sessions are gaming sessions set in a specific campaing. They include notes from the session and stored record of the fate session tracker and dice roller associated with the gaming session.

README and session notes should be stored as markdown files and displayed as text. Fate session tracker and dice roller results should be stored as JSON-files. They are displayed as a fate session tracker and dice roller UI view built with Wish. Chronicle and campaing resource files are blob files downloaded and not displayed in the TUI.

## Building and Running

### Prerequisites

- Go 1.21 or later
- SSH client (for connecting to the server)

### Build

```bash
go build -o dftui .
```

### Run

Start the SSH server (default port 2222):

```bash
./dftui
```

Or specify a custom port:

```bash
./dftui --port 3000
```

The server will automatically generate an SSH host key on first run at `~/.dftui/id_rsa`.

### Connect

From another terminal:

```bash
ssh localhost -p 2222
```

Your SSH username will be used to identify you in the application.

### Keyboard Shortcuts

- **Tab** or **Right Arrow**: Navigate to next tab
- **Shift+Tab** or **Left Arrow**: Navigate to previous tab
- **1-5**: Jump directly to tabs (1=Characters, 2=Sessions, 3=Chronicles, 4=Campaigns, 5=Fate Tracker)
- **q** or **Ctrl+C**: Quit the application

## Development

### Project Structure

```
dftui/
├── main.go              # Entry point, SSH server setup
├── go.mod               # Go module dependencies
├── models/              # Data models
│   ├── character.go     # Character data structure
│   └── character_test.go
├── services/            # Backend services
│   └── backend.go       # Backend interface and stub implementation
└── ui/                  # User interface
    ├── model.go         # Main Bubble Tea model and tab logic
    └── characters.go    # Characters tab view
```

### Running Tests

```bash
go test ./...
```

### Backend Integration

The current implementation uses stub backend services (see `services/backend.go`). These return mock data and are marked with TODO comments for future implementation with the actual Dark Fate RPG backend API.

## Technology Stack

- **Language**: Go
- **SSH Framework**: [Charm Wish](https://github.com/charmbracelet/wish)
- **TUI Framework**: [Charm Bubble Tea](https://github.com/charmbracelet/bubbletea)
- **Styling**: [Charm Lip Gloss](https://github.com/charmbracelet/lipgloss)
