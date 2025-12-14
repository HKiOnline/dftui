# Design Document: Blood Potency Implementation

## Overview
This document outlines the design for implementing blood potency support in vampire characters within the dftui system.

## Architecture
The implementation requires changes across multiple components:
1. Data model in `dflib/dfm/character.go`
2. Documentation examples in `docs/vampire_character.json`
3. JSON schema compliance

## Implementation Details
The blood potency field should:
- Be an integer type
- Have a default value of 1 (as specified in documentation)
- Only be present for vampire characters
- Be properly serialized/deserialized in JSON

## Considerations
- The field should be optional in the struct to maintain compatibility with existing data
- Should follow existing code patterns and naming conventions
- Must preserve backward compatibility for existing vampire character data

## Validation Approach
- Update existing tests to include blood potency validation
- Ensure JSON marshaling/unmarshaling works correctly
- Verify that the field is properly handled in all character operations