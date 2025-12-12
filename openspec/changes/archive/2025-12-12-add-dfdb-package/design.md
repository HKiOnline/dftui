# Design: dfdb Character Database Package

## Context

The Dark Fate TUI requires a backend library for managing character data. This library (`dflib`) needs to be designed for future extraction into a separate Go module that can be reused across multiple Dark Fate projects. The `dfdb` package is the first component, providing character storage with a provider-based interface.

The design follows the existing `promptsdb` pattern from `samples/promptsdb/` which demonstrates:
- A `Provider` interface for CRUD operations
- A filesystem provider implementation using YAML files
- In-memory caching with mutex-protected access
- Factory function for provider instantiation

## Goals

- Create an isolated `dflib` package structure ready for future module extraction
- Implement provider interface pattern for database abstraction
- Use JSON files for character storage (per documentation spec)
- Support full Dark Fate character data model (vampire, ghoul, human spirits)
- Thread-safe operations for concurrent TUI access

## Non-Goals

- Database backends other than filesystem (deferred to future needs)
- Character validation logic (belongs in application layer)
- Migration tooling from other formats
- Network-based storage providers

## Package Structure

```
dflib/
  dfdb/                    # Character database package
    provider.go            # Provider interface definition
    fs_provider.go         # Filesystem JSON provider implementation
    fs_provider_test.go    # Provider tests
    new.go                 # Factory function
  dfm/                     # Models package
    character.go           # Full character model with JSON/YAML tags
    aspect.go              # Aspect model
    skill.go               # Skill model
    stunt.go               # Stunt model
    discipline.go          # Discipline model (vampire-only)
    consequence.go         # Consequence model
    query.go               # Query types for filtering
```

## Decisions

### Decision 1: Provider Interface Pattern
**What**: Use interface-based abstraction for storage backends.
**Why**: Enables future backends (SQLite, PostgreSQL, cloud storage) without changing consumer code. Proven pattern from `promptsdb`.

```go
type Provider interface {
    Create(character Character) error
    Read(characterID string) (Character, error)
    Update(character Character) error
    Delete(characterID string) error
    List(query CharacterQuery) ([]Character, error)
}
```

### Decision 2: JSON File Storage
**What**: Store characters as individual JSON files in a configurable directory (default: `db/characters/`).
**Why**: 
- Matches the documented character JSON format in `docs/characters_json_format.md`
- Human-readable filenames include character name for easy identification
- Git-friendly
- Simple deployment without database dependencies

**File naming pattern**: `{character_name}_{uuid}.json`
- Character name lowercased with spaces replaced by underscores
- Followed by underscore and UUID v4
- Example: `john_smith_550e8400-e29b-41d4-a716-446655440000.json`

**Character name validation**: Character names containing non-alphanumeric characters (except spaces) SHALL be rejected. Only letters (a-z, A-Z), numbers (0-9), and spaces are allowed in character names.

**File identification**: Provider identifies files by matching UUID v4 pattern in filename suffix: `_([0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12})\.json$`. The character name prefix is ignored during file identification.

**Directory behavior**: If the configured directory does not exist, the provider SHALL create it on initialization.

### Decision 3: In-Memory Cache
**What**: Load all characters into memory on provider initialization.
**Why**:
- Fast reads for TUI responsiveness
- Character counts are manageable (dozens to hundreds, not thousands)
- Follows `promptsdb` pattern

**Trade-off**: Higher memory usage, but acceptable for expected data volumes.

### Decision 4: Separate Models Package (dfm)
**What**: Place character models in `dflib/dfm/` separate from `dfdb`.
**Why**:
- Models can be used independently of storage
- Future `dice` package may need model types
- Clean separation of concerns

### Decision 5: Full Character Model
**What**: Implement complete character model per `docs/characters_json_format.md`.
**Why**: Supports all character types (vampire, ghoul, human) with their specific attributes.

**Model includes**:
- Core attributes (id, name, player, spirit, group, etc.)
- Aspects (type-specific: high concept, trouble, clan, covenant, relationship, free)
- Skills (20 skills with title, group, rating)
- Stunts (title, description)
- Disciplines (vampire-only, 16 types with ratings)
- Consequences (level, isActive, title)
- Stress tracks (physical, mental, hunger for vampires)

### Decision 6: Character Query Filtering
**What**: CharacterQuery supports filtering by spirit, player, and group.
**Why**: Common access patterns in the TUI require filtering characters by these attributes.

```go
type CharacterQuery struct {
    Spirit string // "vampire", "ghoul", "human", or empty for all
    Player string // username or empty for all
    Group  string // "pc", "npc", or empty for all
}
```

### Decision 7: File Rename on Name Change
**What**: When a character's name changes via Update, rename the JSON file to reflect the new name.
**Why**: Keeps filenames human-readable and consistent with character data.

**Implementation**: Use write-new-then-delete-old approach for atomicity and safety.

## Alternatives Considered

### SQLite Instead of JSON Files
**Rejected**: Adds deployment complexity, overkill for current scale. Can add as alternative provider later.

### Single Character File (All in One)
**Rejected**: Poor Git merge behavior, doesn't scale with character count.

### Shared models/ Package
**Rejected**: Couples dflib to dftui structure, prevents clean module extraction.

### Simple os.Rename for File Rename
**Rejected**: Write-new-then-delete-old is safer and prevents data loss on failure.

## Risks / Trade-offs

| Risk | Impact | Mitigation |
|------|--------|------------|
| File I/O performance | Low for expected character counts | In-memory cache handles reads |
| Concurrent file writes | Data corruption | Mutex-protected operations |
| JSON schema drift | Incompatible files | Version field in model (future) |
| File rename race condition | Data loss | Write-new-then-delete-old pattern |

## Migration Plan

1. Create `dflib/dfdb/` and `dflib/dfm/` packages
2. Implement filesystem provider with tests
3. Update `services/backend.go` to use dfdb provider
4. Update `ui/characters.go` to import from `dflib/dfm/`
5. Remove `models/` directory (replaced by `dflib/dfm/`)
6. Create `db/characters/` as default storage location

No data migration needed - this is a new capability.
