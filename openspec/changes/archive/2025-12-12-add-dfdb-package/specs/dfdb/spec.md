# dfdb - Dark Fate Character Database

## ADDED Requirements

### Requirement: Provider Interface

The dfdb package SHALL define a Provider interface that abstracts character storage operations, enabling multiple backend implementations.

#### Scenario: Provider interface defines CRUD operations
- **WHEN** a consumer imports the dfdb package
- **THEN** the Provider interface SHALL expose Create, Read, Update, Delete, and List methods for character management

#### Scenario: Provider returns errors for invalid operations
- **WHEN** a provider method fails (e.g., character not found, write error)
- **THEN** the method SHALL return a descriptive error

---

### Requirement: Filesystem Provider

The dfdb package SHALL provide a filesystem-based Provider implementation that stores characters as JSON files in a configurable directory.

#### Scenario: Initialize filesystem provider with directory
- **WHEN** a filesystem provider is created with a directory path
- **THEN** the provider SHALL create the directory if it does not exist
- **AND** the provider SHALL load all existing JSON character files into an in-memory cache
- **AND** the provider SHALL identify files by matching UUID v4 pattern in filename suffix

#### Scenario: Create character persists to JSON file
- **WHEN** Create is called with a valid Character
- **THEN** the provider SHALL write the character to a JSON file named `{name}_{id}.json` where name is lowercased and spaces are replaced with underscores
- **AND** the provider SHALL add the character to the in-memory cache

#### Scenario: Create character with invalid name rejected
- **WHEN** Create is called with a Character whose name contains non-alphanumeric characters (except spaces)
- **THEN** the provider SHALL return an error indicating invalid character name
- **AND** no file SHALL be written

#### Scenario: Read character returns from cache
- **WHEN** Read is called with a valid character ID
- **THEN** the provider SHALL return the character from the in-memory cache
- **AND** the operation SHALL NOT require filesystem access

#### Scenario: Read character not found
- **WHEN** Read is called with a non-existent character ID
- **THEN** the provider SHALL return an error indicating the character was not found

#### Scenario: Update character modifies JSON file
- **WHEN** Update is called with a modified Character whose name has not changed
- **THEN** the provider SHALL overwrite the existing JSON file
- **AND** the provider SHALL update the in-memory cache

#### Scenario: Update character renames file if name changed
- **WHEN** Update is called with a Character whose name has changed
- **THEN** the provider SHALL write a new JSON file with the updated lowercased name prefix
- **AND** the provider SHALL delete the old JSON file
- **AND** the provider SHALL update the in-memory cache

#### Scenario: Update character with invalid name rejected
- **WHEN** Update is called with a Character whose new name contains non-alphanumeric characters (except spaces)
- **THEN** the provider SHALL return an error indicating invalid character name
- **AND** no changes SHALL be made to the file or cache

#### Scenario: Delete character removes JSON file
- **WHEN** Delete is called with a valid character ID
- **THEN** the provider SHALL remove the JSON file from the filesystem
- **AND** the provider SHALL remove the character from the in-memory cache

#### Scenario: List characters returns all cached characters
- **WHEN** List is called with a query
- **THEN** the provider SHALL return characters from the in-memory cache matching the query criteria

---

### Requirement: Thread-Safe Operations

The dfdb filesystem provider SHALL support concurrent access from multiple goroutines safely.

#### Scenario: Concurrent reads do not block
- **WHEN** multiple goroutines call Read simultaneously
- **THEN** all reads SHALL complete without blocking each other

#### Scenario: Writes are serialized
- **WHEN** multiple goroutines call Create, Update, or Delete simultaneously
- **THEN** operations SHALL be serialized to prevent data corruption

---

### Requirement: Character Query

The dfdb package SHALL provide a CharacterQuery type for filtering characters in List operations.

#### Scenario: Query by spirit type
- **WHEN** List is called with a query specifying spirit (vampire, ghoul, human)
- **THEN** only characters matching the spirit type SHALL be returned

#### Scenario: Query by group
- **WHEN** List is called with a query specifying group (pc, npc)
- **THEN** only characters matching the group SHALL be returned

#### Scenario: Query by player
- **WHEN** List is called with a query specifying a player username
- **THEN** only characters belonging to that player SHALL be returned

#### Scenario: Query with multiple filters
- **WHEN** List is called with multiple filter criteria
- **THEN** only characters matching ALL criteria SHALL be returned (AND logic)

#### Scenario: Query all characters
- **WHEN** List is called with an empty query
- **THEN** all characters SHALL be returned

---

### Requirement: Provider Factory

The dfdb package SHALL provide a factory function to instantiate providers by type.

#### Scenario: Create filesystem provider via factory
- **WHEN** New is called with provider type "filesystem" and configuration
- **THEN** a configured filesystem provider SHALL be returned

#### Scenario: Default to filesystem provider
- **WHEN** New is called with an unknown provider type
- **THEN** a filesystem provider SHALL be returned as the default

---

### Requirement: Character Model (dfm package)

The dflib/dfm package SHALL define the complete Character model with JSON and YAML struct tags for serialization.

#### Scenario: Character model supports all spirit types
- **WHEN** a Character is instantiated
- **THEN** it SHALL support vampire, ghoul, and human spirit types with their respective attributes

#### Scenario: Character serializes to JSON per specification
- **WHEN** a Character is marshaled to JSON
- **THEN** the output SHALL conform to the format defined in docs/characters_json_format.md

#### Scenario: Character deserializes from JSON
- **WHEN** a JSON file conforming to the character format is unmarshaled
- **THEN** a valid Character struct SHALL be populated with all attributes

#### Scenario: Vampire characters include disciplines and hunger stress
- **WHEN** a Character has spirit "vampire"
- **THEN** it SHALL include disciplines array and hungerStress attributes

#### Scenario: Model includes all character attributes
- **WHEN** a Character is created
- **THEN** it SHALL include: id, player, category, spirit, group, name, gender, aliases, tags, collectives, embrace_year, setting_year, notes, refresh, aspects, skills, stunts, consequences, physicalStressLimit, physicalStressCurrent, mentalStressLimit, mentalStressCurrent
