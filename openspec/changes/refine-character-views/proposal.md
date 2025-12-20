# Change: Refine Character Views and Data Format

## Why
The current character views and data format need refinement to better support the Dark Fate RPG gameplay. The changes will add missing fate point tracking, improve character display organization, and enhance the user experience by providing more relevant information in a better structured format.

## What Changes
- Add `fatePoint` field to character data model to track current fate points
- Update character JSON format documentation and sample files
- Refine character detail view layout and information organization
- Improve character list view formatting and information display
- Update dflib/dfm character model to include fatePoint field

## Impact
- Affected specs: character-list
- Affected code: dflib/dfm/character.go, docs/characters_json_format.md, sample character JSON files
- Affected UI: character detail view, character list view
- This is a non-breaking change that adds functionality without removing existing features