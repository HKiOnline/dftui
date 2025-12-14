# Tasks for Blood Potency Implementation

1. Update Character struct in dflib/dfm/character.go to include BloodPotency field
2. Update vampire_character.json example file to include bloodPotency attribute
3. Update character tests in dflib/dfm/character_test.go to include blood potency validation
4. Verify JSON serialization/deserialization works correctly
5. Run existing tests to ensure no regressions
6. Update documentation if needed to reflect the implementation

## Dependencies
- No external dependencies
- All changes are self-contained within the dfm package and docs directory

## Validation Criteria
- Existing vampire character JSON data can still be parsed
- New bloodPotency field is properly serialized and deserialized
- All existing tests pass
- Vampire character example file includes the new field