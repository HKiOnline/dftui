# Implementation Tasks

## 1. Character Data Model Updates
- [ ] 1.1 Add `fatePoint` field to Character struct in dflib/dfm/character.go
- [ ] 1.2 Update JSON and YAML tags for the new field
- [ ] 1.3 Add appropriate default value (0) for fatePoint

## 2. Documentation Updates
- [ ] 2.1 Update docs/characters_json_format.md to include fatePoint field
- [ ] 2.2 Update sample character JSON files (vampire_character.json, ghoul_character.json, human_character.json) with fatePoint field
- [ ] 2.3 Ensure documentation mentions fatePoint is placed after refresh field

## 3. Character Detail View Updates
- [ ] 3.1 Remove ID field from character detail view display
- [ ] 3.2 Move player field just below character name
- [ ] 3.3 Add Fate section after basic section with "Refresh" and "Points Available" display
- [ ] 3.4 Capitalize skill names and add space before skill rank in skill section
- [ ] 3.5 Rename Disciplines section to "Beast and Blood" and include Blood Potency
- [ ] 3.6 Capitalize discipline names
- [ ] 3.7 Remove redundant Vampire-information section
- [ ] 3.8 Add notes section at bottom displaying character notes

## 4. Character List View Updates
- [ ] 4.1 Remove type indicators (PC/NPC) from character names in list
- [ ] 4.2 Add dash and description after character name in list
- [ ] 4.3 Truncate description to 50 characters in list view
- [ ] 4.4 Ensure characters are grouped by PC/NPC sections

## 5. Testing and Validation
- [ ] 5.1 Test character data loading with new fatePoint field
- [ ] 5.2 Verify character detail view displays all information correctly
- [ ] 5.3 Test character list view formatting and truncation
- [ ] 5.4 Validate JSON schema compatibility