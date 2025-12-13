# Character JSON Format Specification

## ADDED Requirements

### Blood Potency Attribute
#### Requirement: Vampire characters include blood potency attribute
##### Scenario: Vampire characters include blood potency attribute
- Given a vampire character JSON document
- When the document is parsed
- Then it should include a bloodPotency field with integer value

##### Scenario: Blood potency defaults to 1 for vampire characters
- Given a new vampire character JSON document without bloodPotency specified
- When the document is parsed
- Then the bloodPotency field should default to 1

##### Scenario: Blood potency is serialized correctly
- Given a vampire character with bloodPotency set to 3
- When the document is serialized to JSON
- Then the resulting JSON should include "bloodPotency": 3

## MODIFIED Requirements

### Vampire Character Specification
#### Requirement: Vampire character JSON includes blood potency field
##### Scenario: Vampire character JSON includes blood potency field
- Given the current vampire character specification in docs/characters_json_format.md
- When the specification is updated to include blood potency
- Then the example vampire_character.json should reflect this change

## REMOVED Requirements

No requirements are removed by this change.