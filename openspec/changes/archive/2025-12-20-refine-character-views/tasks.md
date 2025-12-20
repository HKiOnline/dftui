# Implementation Tasks

## 1. Character Data Model Updates
- [x] 1.1 Add `fatePoint` field to Character struct in dflib/dfm/character.go
- [x] 1.2 Update JSON and YAML tags for the new field
- [x] 1.3 Add appropriate default value (0) for fatePoint

## 2. Documentation Updates
- [x] 2.1 Update docs/characters_json_format.md to include fatePoint field
- [x] 2.2 Update sample character JSON files (vampire_character.json, ghoul_character.json, human_character.json) with fatePoint field
- [x] 2.3 Ensure documentation mentions fatePoint is placed after refresh field

## 3. Character Detail View Updates
- [x] 3.1 Remove ID field from character detail view display
- [x] 3.2 Move player field just below character name
- [x] 3.3 Add Fate section after basic section with "Refresh" and "Points Available" display
- [x] 3.4 Capitalize skill names and add space before skill rank in skill section
- [x] 3.5 Rename Disciplines section to "Beast and Blood" and include Blood Potency
- [x] 3.6 Capitalize discipline names
- [x] 3.7 Remove redundant Vampire-information section
- [x] 3.8 Add notes section at bottom displaying character notes

## 4. Character List View Updates
- [x] 4.1 Remove type indicators (PC/NPC) from character names in list
- [x] 4.2 Add dash and description after character name in list
- [x] 4.3 Truncate description to 50 characters in list view
- [x] 4.4 Ensure characters are grouped by PC/NPC sections

## 5. Testing and Validation
- [x] 5.1 Test character data loading with new fatePoint field
- [x] 5.2 Verify character detail view displays all information correctly
- [x] 5.3 Test character list view formatting and truncation
- [x] 5.4 Validate JSON schema compatibility