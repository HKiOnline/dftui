# Change: Add Character Detail View

## Why
Users need to view detailed information about individual characters. Currently, the Characters tab only shows a list of character names and types (PC/NPC), but provides no way to see additional character details or interact with individual characters.

## What Changes
- Add selection capability to the character list (using up/down arrow keys and enter to select)
- Add a character detail view that displays when a character is selected
- Add navigation to return from detail view back to character list (using ESC or back key)
- Update UI state management to track selected character and view mode (list vs detail)

## Impact
- Affected specs: `character-list`
- Affected code:
  - `ui/model.go` - Add state for selected character and view mode
  - `ui/characters.go` - Add detail view rendering and list selection UI
  - May need to extend `models.Character` if additional detail fields are required (TBD based on backend data)
