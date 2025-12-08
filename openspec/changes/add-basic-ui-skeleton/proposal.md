# Change: Add Basic UI Skeleton with SSH Server and Tabbed Interface

## Why
This is the foundational change that establishes the core UI framework for dftui. We need a working SSH server that can accept connections, identify users, and display a basic tabbed interface. Starting with the Characters tab provides a concrete foundation for the full application while keeping the scope manageable.

## What Changes
- Initialize Go module and project structure
- Set up SSH server using Charm Wish framework
- Implement user identification from SSH login
- Create tabbed UI framework with navigation
- Implement Characters tab with placeholder character list
- Add stub functions for future backend integration

## Impact
- Affected specs: `ssh-server`, `tabbed-ui`, `character-list` (all new)
- Affected code: New Go project structure
  - `main.go` - entry point and SSH server setup
  - `ui/tabs.go` - tabbed interface framework
  - `ui/characters.go` - characters tab view
  - `models/character.go` - character data model
  - `services/backend.go` - stub backend functions
