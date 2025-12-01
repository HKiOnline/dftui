<!-- OPENSPEC:START -->
# OpenSpec Instructions

These instructions are for AI assistants working in this project.

Always open `@/openspec/AGENTS.md` when the request:
- Mentions planning or proposals (words like proposal, spec, change, plan)
- Introduces new capabilities, breaking changes, architecture shifts, or big performance/security work
- Sounds ambiguous and you need the authoritative spec before coding

Use `@/openspec/AGENTS.md` to learn:
- How to create and apply change proposals
- Spec format and conventions
- Project structure and guidelines

Keep this managed block so 'openspec update' can refresh the instructions.

<!-- OPENSPEC:END -->

# Dark Fate Terminal UI - Project Context

## Overview
Dark Fate Terminal UI (dftui) is an SSH-accessible terminal user interface for the Dark Fate RPG manager. Built with Go and the Charm Wish framework, it provides a tabbed interface for managing RPG chronicles, campaigns, sessions, and characters.

## Data Model Hierarchy
Understanding the data hierarchy is crucial for working on this project:

```
Chronicle (story setting)
  ├── README.md
  ├── Resource files (PDFs, etc.)
  └── Campaigns
        ├── README.md
        ├── Resource files
        ├── Characters (PCs and NPCs)
        └── Sessions
              ├── Session notes (markdown)
              └── Fate tracker data (JSON)
```

**Key Relationships:**
- Chronicles contain Campaigns
- Campaigns contain Sessions
- Campaigns have associated Characters (both PC and NPC)
- Each level has READMEs (markdown) and downloadable resources
- Sessions store notes (markdown) and game state (JSON)

## UI Structure
The TUI has five main tabs:

1. **Characters** - List of PCs and NPCs visible to user
2. **Sessions** - Active game sessions
3. **Chronicles** - Browse chronicles, view READMEs, download resources
4. **Campaigns** - Browse campaigns, sessions, notes, and resources
5. **Fate Tracker** - Dice roller with single-mode and session-mode

## File Formats
- **Markdown**: READMEs, session notes (display as formatted text)
- **JSON**: Fate tracker data (render as interactive TUI)
- **Binary blobs**: Resource files (download, not display)

## Technical Stack
- **Language**: Go
- **TUI Framework**: Charm Wish (SSH-based)
- **Protocol**: SSH for remote access
- **Multi-user**: Concurrent connections, shared session state

## When Working on Features
- Always consider the Chronicle → Campaign → Session hierarchy
- Multi-user access patterns (shared game sessions)
- Markdown rendering for text content
- JSON-based state for game mechanics
- File download capabilities for resources
- Real-time updates for shared session data