# Implementation Tasks

## 1. Update UI Model
- [x] 1.1 Add `selectedCharacterIndex` field to Model struct to track cursor position in list
- [x] 1.2 Add `characterViewMode` field to Model struct (enum: list/detail)
- [x] 1.3 Add `selectedCharacter` field to Model struct to store currently selected character

## 2. Implement Character List Selection
- [x] 2.1 Add up/down arrow key handlers in Update() to navigate character list
- [x] 2.2 Add enter key handler to select character and switch to detail view
- [x] 2.3 Update renderCharactersTab() to highlight selected character in list
- [x] 2.4 Add visual cursor/highlight indicator for selected character

## 3. Implement Character Detail View
- [x] 3.1 Create renderCharacterDetail() function to display selected character details
- [x] 3.2 Add ESC key handler to return from detail view to list
- [x] 3.3 Update renderCharactersTab() to switch between list and detail rendering based on view mode
- [x] 3.4 Display character fields (ID, Name, Type) in detail view with formatting

## 4. Testing and Validation
- [x] 4.1 Test navigation through character list with arrow keys
- [x] 4.2 Test selecting character with enter key
- [x] 4.3 Test returning to list from detail view with ESC
- [x] 4.4 Test edge cases (empty list, single character, first/last navigation)
- [x] 4.5 Test visual highlighting and rendering across different terminal sizes

## 5. Documentation
- [x] 5.1 Update help text in renderHelp() to show new navigation keys
- [x] 5.2 Add code comments for new state fields and functions
