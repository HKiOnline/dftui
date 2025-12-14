# character-list Specification Update

## Purpose
Update the character list specification to reflect loading real character data from dfdb instead of using mock/stub data.

## Requirements

### MODIFIED Requirement: Backend Integration
The system SHALL fetch user characters from the dfdb filesystem provider instead of using stub/mock data.

#### Scenario: Fetch characters from dfdb
- **WHEN** the Characters tab requests character data for a user
- **THEN** the backend SHALL use dfdb.FSProvider to load JSON files from db/characters
- **AND** the backend SHALL filter characters by player name (for PCs)
- **AND** the backend SHALL include all NPCs in the results
- **AND** the TODO comment for future implementation SHALL be removed

#### Scenario: Character data matches JSON format
- **WHEN** characters are loaded from db/characters
- **THEN** each character SHALL conform to the format specified in docs/characters_json_format.md
- **AND** vampire characters SHALL include bloodPotency, disciplines, and hunger stress tracks
- **AND** ghoul characters SHALL have appropriate aspects for their spirit type
- **AND** human characters SHALL use basic Fate Condensed rules

### MODIFIED Requirement: Character List Display
The system SHALL display PCs and NPCs in separate sections with clear visual distinction.

#### Scenario: Separate PC and NPC sections
- **WHEN** the Characters tab is active
- **THEN** a "Player Characters" section SHALL be displayed first
- **AND** a "Non-Player Characters" section SHALL be displayed second
- **AND** PCs SHALL only show characters where player name matches logged-in user
- **AND** NPCs SHALL show all non-player characters from db/characters

#### Scenario: Character display includes spirit type
- **WHEN** a character is displayed in the list
- **THEN** the spirit type (vampire, ghoul, human) SHALL be shown
- **AND** the display SHALL use appropriate icons or text indicators for each type

### ADDED Requirement: Character Data Validation
The system SHALL validate loaded character data and handle errors gracefully.

#### Scenario: Invalid character JSON
- **WHEN** a character file contains invalid JSON
- **THEN** the system SHALL log an error
- **AND** the invalid file SHALL be skipped
- **AND** valid characters SHALL still be displayed

#### Scenario: Missing required fields
- **WHEN** a character is missing required fields (id, name, group)
- **THEN** the system SHALL log a warning
- **AND** the incomplete character SHALL be skipped
- **AND** other characters SHALL still be displayed

### MODIFIED Requirement: Character Detail View
The system SHALL display complete character information in an organized detail view.

#### Scenario: Display basic character information
- **WHEN** a character is selected and detail view is active
- **THEN** the detail view SHALL show:
  - Character name
  - Spirit type (vampire, ghoul, human)
  - Player/owner name
  - Group (PC/NPC indicator)
  - Embrace year (for vampires/ghouls) or setting year (for humans)
- **AND** the information SHALL be clearly labeled and formatted

#### Scenario: Display spirit-specific information
- **WHEN** a vampire character is displayed in detail view
- **THEN** the detail view SHALL show:
  - Blood potency
  - Disciplines with ratings
  - Hunger stress track
- **WHEN** a ghoul character is displayed in detail view
- **THEN** the detail view SHALL show:
  - Master relationship aspect
  - Appropriate stunts for ghouls
- **WHEN** a human character is displayed in detail view
- **THEN** the detail view SHALL show:
  - Standard Fate Condensed aspects
  - Basic stress tracks (physical and mental)

#### Scenario: Display aspects organized by type
- **WHEN** character aspects are displayed in detail view
- **THEN** they SHALL be grouped by type:
  - High concept
  - Trouble
  - Relationship
  - Free
  - Clan (vampire only)
  - Covenant (vampire/ghoul only)
- **AND** each aspect SHALL show title and description if available

#### Scenario: Display skills organized by type
- **WHEN** character skills are displayed in detail view
- **THEN** they SHALL be grouped by category:
  - Mental skills
  - Physical skills
  - Social skills
- **AND** each skill SHALL show title and rating
- **AND** skills with rating > 0 SHALL be highlighted or emphasized

#### Scenario: Display stunts and consequences
- **WHEN** character stunts are displayed in detail view
- **THEN** each stunt SHALL show:
  - Title
  - Description
- **WHEN** character consequences are displayed in detail view
- **THEN** each consequence SHALL show:
  - Level (2, 4, or 6)
  - Active status
  - Title if active
