# Add Blood Potency to Vampire Characters

## Summary
This change adds the `bloodPotency` attribute to vampire character JSON format and implementation, aligning with the existing documentation specification.

## Problem
The `bloodPotency` attribute is documented in `docs/characters_json_format.md` as a valid field for vampire characters with default value 1, but it is missing from:
1. The Go struct implementation in `dflib/dfm/character.go`
2. The vampire character example JSON file in `docs/vampire_character.json`

## Solution
1. Add `BloodPotency` field to the Character struct in `dflib/dfm/character.go`
2. Update `docs/vampire_character.json` to include the `bloodPotency` field
3. Ensure proper JSON serialization/deserialization

## Impact
- Adds support for vampire blood potency in the system
- Maintains backward compatibility with existing character data
- Aligns documentation with implementation