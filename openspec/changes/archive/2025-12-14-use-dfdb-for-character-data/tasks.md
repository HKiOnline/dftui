# Tasks for Use DFDB for Character Data

## Implementation Tasks

### 1. Backend Integration
- [x] Update backend.go to use dfdb for character loading
- [x] Implement GetUserCharacters function using dfdb.FSProvider
- [x] Load characters from db/characters directory
- [x] Filter characters by player name (for PCs) and include all NPCs

### 2. UI Integration
- [x] Update ui/model.go to use new backend character loading
- [x] Display PCs in a "Player Characters" section
- [x] Display NPCs in a "Non-Player Characters" section
- [x] Show character details when selected

### 3. Character Detail View
- [x] Create organized detail view showing:
  - Basic information (name, spirit type)
  - Spirit-specific information (blood potency for vampires)
  - Aspects
  - Skills
  - Disciplines (for vampires)
  - Stunts
  - Stress tracks
  - Consequences

### 4. Testing
- [x] Add unit tests for character loading from dfdb
- [x] Test filtering by player name
- [x] Test empty character list handling
- [x] Verify UI rendering with real character data

## Validation Tasks

### 5. Manual Testing
- [x] Test with sample characters in db/characters
- [x] Verify PCs are shown for matching usernames
- [x] Verify NPCs are always visible
- [x] Check detail view displays all information correctly

### 6. Documentation
- [x] Update README with new character data location
- [x] Document db/characters directory structure
- [x] Add examples of valid character JSON files
