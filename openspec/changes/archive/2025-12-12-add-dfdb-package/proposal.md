# Change: Add dfdb Character Database Package

## Why

The Dark Fate TUI needs a reusable, isolated backend library (`dflib`) that can be shared across multiple projects. The first component of this library is `dfdb` - a database abstraction for storing and managing Dark Fate RPG characters. This enables future extraction into a separate Go module while providing a clean provider-based interface similar to the existing `promptsdb` pattern.

## What Changes

- **NEW**: Create `dflib/dfdb` package with provider interface for character storage
- **NEW**: Create `dflib/dfm` package for character models with JSON/YAML struct tags
- **NEW**: Implement filesystem provider using JSON files for character persistence
- **NEW**: Create default `db/characters/` storage directory
- **REMOVE**: Delete `models/` directory (replaced by `dflib/dfm/`)
- Update TUI to use the new dfdb provider for character data

## Impact

- Affected specs: character-list (data source changes)
- Affected code: `ui/characters.go`, `services/backend.go`
- New packages: `dflib/dfdb/`, `dflib/dfm/`
- Removed: `models/` directory
- **No breaking changes** to existing TUI functionality
