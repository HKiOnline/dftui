# Character List Specification Updates

## ADDED Requirements

### Requirement: Character Fate Point Tracking
The system SHALL track current fate points for each character.

#### Scenario: Fate point field in character data
- **WHEN** a character is loaded or created
- **THEN** the character SHALL have a fatePoint field
- **AND** the fatePoint field SHALL be a numeric value
- **AND** the fatePoint field SHALL default to 0

#### Scenario: Fate point display in detail view
- **WHEN** a character detail view is displayed
- **THEN** the fatePoint value SHALL be shown in the Fate section
- **AND** it SHALL be labeled as "Points Available"

### Requirement: Character Description Field
The system SHALL support and display character descriptions.

#### Scenario: Description field in character data
- **WHEN** a character is loaded or created
- **THEN** the character SHALL have a description field
- **AND** the description field SHALL be a string value
- **AND** the description field SHALL default to empty string

#### Scenario: Description display in list view
- **WHEN** a character is displayed in the list view
- **THEN** the description SHALL be shown after the character name
- **AND** the description SHALL be truncated to 50 characters
- **AND** the description SHALL be preceded by a dash separator

### Requirement: Character Notes Display
The system SHALL display character notes in the detail view.

#### Scenario: Notes section in detail view
- **WHEN** a character detail view is displayed
- **AND** the character has notes content
- **THEN** a Notes section SHALL be shown at the bottom
- **AND** the notes content SHALL be displayed

#### Scenario: Empty notes handling
- **WHEN** a character detail view is displayed
- **AND** the character has no notes
- **THEN** the Notes section SHALL be omitted or shown as empty

## MODIFIED Requirements

### Requirement: Character List Display
The system SHALL display a list of characters associated with the logged-in user with improved formatting.

#### Scenario: Characters tab shows user's characters with descriptions
- **WHEN** the Characters tab is active
- **THEN** a list of characters SHALL be displayed
- **AND** each character SHALL show its name followed by a dash and truncated description
- **AND** the type indicator (PC or NPC) SHALL be omitted from the character name
- **AND** characters SHALL be grouped by PC and NPC sections
- **AND** one character SHALL be selected by default (the first character if list is not empty)

### Requirement: Character Detail View
The system SHALL provide a detailed view for a selected character with reorganized information.

#### Scenario: Display character details with improved layout
- **WHEN** a character is selected and detail view is active
- **THEN** the detail view SHALL display the character's Name prominently
- **AND** the detail view SHALL display the Player field just below the character name
- **AND** the detail view SHALL NOT display the character's ID
- **AND** the detail view SHALL include a Fate section after basic information
- **AND** the Fate section SHALL show "Refresh" and "Points Available" values
- **AND** the detail view SHALL include a "Beast and Blood" section for vampire characters
- **AND** the "Beast and Blood" section SHALL show Blood Potency and disciplines
- **AND** discipline names SHALL be capitalized
- **AND** skill names SHALL be capitalized with space before rank
- **AND** the detail view SHALL NOT include a separate Vampire-information section

### Requirement: Character Data Model
The system SHALL define an enhanced character data structure with fate point tracking.

#### Scenario: Character model structure with fate points
- **WHEN** a character is represented in the system
- **THEN** it SHALL have an ID (unique identifier)
- **AND** it SHALL have a Name (string)
- **AND** it SHALL have a Type (either "PC" or "NPC")
- **AND** it SHALL have a fatePoint field (numeric, default 0)
- **AND** it SHALL have a description field (string, default "")
- **AND** it SHALL have a notes field (string, default "")
- **AND** the fatePoint field SHALL be placed after the refresh field in the data structure

## REMOVED Requirements

### Requirement: Character Type Indicators in List
The system SHALL NO LONGER display type indicators next to character names in the list view.

#### Scenario: Type indicators removed from list view
- **WHEN** characters are displayed in the list view
- **THEN** character names SHALL NOT be prefixed with type indicators like "[PC]" or "[NPC]"
- **AND** character grouping by PC/NPC sections SHALL provide sufficient type distinction