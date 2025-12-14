# Database Directory Structure

## Overview
The `db` directory contains all character data for the Dark Fate Terminal UI application. Character files are stored as JSON and automatically loaded by the application at startup.

## Directory Layout

```
db/
├── characters/          # Character JSON files
│   ├── {name}_{uuid}.json  # Individual character files
│   └── ...                # More character files
└── users.json            # User configuration (reserved for future use)
```

## Characters Directory

The `db/characters` directory stores all character data as JSON files. Each file represents a single character and must follow the naming convention:

**Filename Format:** `{name}_{uuid}.json`

Where:
- `{name}` is the character's name converted to lowercase with spaces replaced by underscores
- `{uuid}` is a UUID v4 identifier (e.g., `550e8400-e29b-41d4-a716-446655440000`)

### Examples:
- `victor_joki_550e8400-e29b-41d4-a716-446655440000.json`
- `nathan_quincy_550e8400-e29b-41d4-a716-446655440001.json`

## Character JSON Format

Each character file must conform to the format specified in [characters_json_format.md](characters_json_format.md). The application validates that each character has:

- `id`: UUID v4 identifier (must match filename)
- `name`: Character name
- `group`: Either "pc" (Player Character) or "npc" (Non-Player Character)
- `player`: Username of the player who owns this character (for PCs) or gamemaster (for NPCs)
- `spirit`: Type of character - "vampire", "ghoul", or "human"

## Loading Behavior

The application:
1. Scans all `.json` files in `db/characters`
2. Validates each file contains proper JSON and required fields
3. Skips files with invalid JSON or missing required fields (logs warning)
4. Loads valid characters into memory
5. Filters characters based on logged-in user:
   - Shows all PCs where `player` field matches the username
   - Shows all NPCs regardless of player field
6. Displays PCs first, then NPCs in separate sections

## Best Practices

1. **Use UUIDs**: Always use valid UUID v4 identifiers for character IDs
2. **Valid filenames**: Ensure filenames match the pattern `{name}_{uuid}.json`
3. **Required fields**: Include all required fields (id, name, group, player, spirit)
4. **Character types**: Use appropriate spirit type for each character:
   - Vampire characters: include `bloodPotency`, `disciplines`, and `hungerStress*` fields
   - Ghoul characters: include appropriate aspects for ghouls
   - Human characters: use basic Fate Condensed rules
5. **Testing**: Test with the sample characters in this directory before creating new ones

## Examples

See the following example character files:
- [vampire_character.json](vampire_character.json) - Vampire PC example
- [ghoul_character.json](ghoul_character.json) - Ghoul NPC example  
- [human_character.json](human_character.json) - Human PC example
