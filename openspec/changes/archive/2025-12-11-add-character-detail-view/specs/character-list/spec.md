# character-list Specification Delta

## ADDED Requirements

### Requirement: Character Selection Navigation
The system SHALL allow users to navigate and select characters in the character list using keyboard controls.

#### Scenario: Navigate character list with arrow keys
- **WHEN** the Characters tab is active and displaying a list of characters
- **AND** the user presses the up or down arrow key
- **THEN** the selection cursor SHALL move to the previous or next character respectively
- **AND** the currently selected character SHALL be visually highlighted

#### Scenario: Select character with Enter key
- **WHEN** a character is highlighted in the character list
- **AND** the user presses the Enter key
- **THEN** the system SHALL switch to character detail view
- **AND** the detail view SHALL display the selected character's information

#### Scenario: Navigate at list boundaries
- **WHEN** the first character is selected
- **AND** the user presses the up arrow key
- **THEN** the selection SHALL remain on the first character (no wrap-around)
- **WHEN** the last character is selected
- **AND** the user presses the down arrow key
- **THEN** the selection SHALL remain on the last character (no wrap-around)

#### Scenario: Empty list selection handling
- **WHEN** the character list is empty (no characters)
- **THEN** arrow keys SHALL have no effect
- **AND** pressing Enter SHALL have no effect
- **AND** the "No characters found" message SHALL be displayed

### Requirement: Character Detail View
The system SHALL provide a detailed view for a selected character with navigation back to the list.

#### Scenario: Display character details
- **WHEN** a character is selected and detail view is active
- **THEN** the detail view SHALL display the character's ID
- **AND** the detail view SHALL display the character's Name
- **AND** the detail view SHALL display the character's Type (PC or NPC)
- **AND** the information SHALL be clearly labeled and formatted

#### Scenario: Return to character list
- **WHEN** the detail view is active
- **AND** the user presses the ESC key
- **THEN** the system SHALL return to the character list view
- **AND** the previously selected character SHALL remain highlighted

#### Scenario: Preserve selection across view changes
- **WHEN** the user switches from list to detail view and back
- **THEN** the same character SHALL remain selected in the list
- **AND** the cursor position SHALL be preserved

### Requirement: Visual Selection Indicator
The system SHALL provide clear visual feedback for the currently selected character in the list view.

#### Scenario: Highlight selected character
- **WHEN** a character is selected in the list
- **THEN** it SHALL be visually distinguished from other characters
- **AND** the visual distinction MAY include highlighting, cursor indicator, or color change
- **AND** only one character SHALL be highlighted at a time

#### Scenario: Selection visibility
- **WHEN** the character list is displayed
- **THEN** the selected character SHALL be clearly visible
- **AND** the selection indicator SHALL be distinguishable from the PC/NPC type indicator

### Requirement: Updated Keyboard Controls Help
The system SHALL update the help text to reflect the new navigation controls for character selection.

#### Scenario: Help text includes character navigation
- **WHEN** the Characters tab is active in list view
- **THEN** the help text SHALL indicate that up/down arrows navigate the list
- **AND** the help text SHALL indicate that Enter selects a character
- **WHEN** the Characters tab is active in detail view
- **THEN** the help text SHALL indicate that ESC returns to the list

## MODIFIED Requirements

### Requirement: Character List Display
The system SHALL display a list of characters associated with the logged-in user with selection capability.

#### Scenario: Characters tab shows user's characters
- **WHEN** the Characters tab is active
- **THEN** a list of characters SHALL be displayed
- **AND** each character SHALL show its name
- **AND** each character SHALL show a type indicator (PC or NPC)
- **AND** one character SHALL be selected by default (the first character if list is not empty)

#### Scenario: Character type indicators
- **WHEN** a character with type "PC" is displayed
- **THEN** it SHALL be marked as "[PC]" or similar visual indicator
- **WHEN** a character with type "NPC" is displayed
- **THEN** it SHALL be marked as "[NPC]" or similar visual indicator

#### Scenario: Empty character list
- **WHEN** the user has no characters
- **THEN** a message SHALL be displayed indicating "No characters found"
- **AND** no selection cursor SHALL be shown
