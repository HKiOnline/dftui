# Tasks: Add dfdb Character Database Package

## 1. Create dflib Package Structure
- [x] 1.1 Create `dflib/dfm/` directory for models
- [x] 1.2 Create `dflib/dfdb/` directory for database provider

## 2. Implement Character Models (dfm)
- [x] 2.1 Create `dflib/dfm/character.go` with full Character struct (JSON/YAML tags)
- [x] 2.2 Create `dflib/dfm/aspect.go` with Aspect struct
- [x] 2.3 Create `dflib/dfm/skill.go` with Skill struct
- [x] 2.4 Create `dflib/dfm/stunt.go` with Stunt struct
- [x] 2.5 Create `dflib/dfm/discipline.go` with Discipline struct
- [x] 2.6 Create `dflib/dfm/consequence.go` with Consequence struct
- [x] 2.7 Create `dflib/dfm/query.go` with CharacterQuery struct (spirit, player, group filters)

## 3. Implement dfdb Provider Interface
- [x] 3.1 Create `dflib/dfdb/provider.go` with Provider interface and ProviderConfiguration
- [x] 3.2 Create `dflib/dfdb/new.go` with factory function

## 4. Implement Filesystem Provider
- [x] 4.1 Create `dflib/dfdb/fs_provider.go` with FsProvider struct
- [x] 4.2 Implement directory creation if missing on initialization
- [x] 4.3 Implement cache loading with UUID v4 pattern matching for file identification
- [x] 4.4 Implement Create method (write JSON with `{name}_{id}.json` naming, update cache)
- [x] 4.5 Implement Read method (return from cache)
- [x] 4.6 Implement Update method (rename file if name changed, update contents, update cache)
- [x] 4.7 Implement Delete method (remove file, remove from cache)
- [x] 4.8 Implement List method with query filtering (spirit, player, group)
- [x] 4.9 Add mutex protection for thread safety

## 5. Testing
- [x] 5.1 Create `dflib/dfm/character_test.go` - test JSON marshal/unmarshal
- [x] 5.2 Create `dflib/dfdb/fs_provider_test.go` - test CRUD operations
- [x] 5.3 Test file naming convention (lowercased `{name}_{uuid}.json`)
- [x] 5.4 Test file rename on character name change
- [x] 5.5 Test character name validation (reject non-alphanumeric except spaces)
- [x] 5.6 Test directory auto-creation
- [x] 5.7 Test concurrent access scenarios
- [x] 5.8 Test query filtering (by spirit, group, player, and combinations)

## 6. Integration
- [x] 6.1 Remove `models/` directory
- [x] 6.2 Update `services/backend.go` to use dfdb provider and dfm models
- [x] 6.3 Update `ui/characters.go` to import from `dflib/dfm/`
- [x] 6.4 Update `models/character_test.go` tests or migrate to dfm package (tests migrated to dflib/dfm/)
- [x] 6.5 Create `db/characters/` as default storage directory
- [x] 6.6 Verify existing TUI functionality works unchanged

## Dependencies
- Tasks 2.x must complete before 4.x (models needed for provider)
- Tasks 3.x and 4.x can proceed in parallel after 2.x
- Task 5.x requires 4.x completion
- Task 6.x requires 4.x and 5.x completion
