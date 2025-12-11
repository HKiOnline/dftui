# Design: Character Detail View

## Context
The current Characters tab displays a simple list of character names with PC/NPC indicators. Users need to be able to select characters and view more details. This is the first interactive list component in the application and will establish patterns for other list-based tabs (Sessions, Chronicles, Campaigns).

The Charm Bubble Tea framework follows the Elm architecture (Model-Update-View), so state changes happen through message passing in Update(), and the View() function renders the current state.

## Goals / Non-Goals

**Goals:**
- Enable keyboard-driven selection in character list
- Display character details in a dedicated view
- Provide clear navigation between list and detail views
- Establish reusable patterns for list selection in other tabs

**Non-Goals:**
- Character editing functionality (read-only for now)
- Character creation/deletion
- Filtering or searching characters
- Loading additional character data beyond what's currently in the Character model

## Decisions

### Decision: View Mode State Machine
Use a simple enum-based state machine with two modes: "list" and "detail".

**Rationale:**
- Clear separation of concerns between list and detail rendering
- Easy to extend later with additional modes (e.g., edit mode)
- Matches Bubble Tea's state-based rendering model

**Alternatives considered:**
- Stack-based navigation (push/pop views) - overly complex for two views
- Component-based approach - premature abstraction before patterns emerge

### Decision: Cursor Position Tracking
Track selected index in the character list with bounds checking.

**Rationale:**
- Simple integer index is efficient and sufficient
- Works well with slice-based character storage
- Easy to persist across view mode changes

**Implementation notes:**
- Index should persist when switching from list to detail and back
- Bounds checking needed when character list changes (add/remove)
- Initialize to 0 or -1 for empty list

### Decision: Key Bindings
- **Up/Down arrows**: Navigate list (standard TUI convention)
- **Enter**: Select character and view details
- **ESC**: Return to list from detail view
- **q/Ctrl+C**: Quit application (unchanged)

**Rationale:**
- Follows common TUI conventions (vim, less, etc.)
- ESC for "back" is intuitive and standard
- Doesn't conflict with existing tab navigation (Tab/Shift+Tab, 1-5)

### Decision: Detail View Content
Display all fields currently available in the Character model (ID, Name, Type) with clear labels.

**Rationale:**
- Start minimal with available data
- Placeholder for future character attributes (stats, description, etc.)
- Demonstrates the detail view pattern without waiting for backend expansion

**Future extension:**
When backend provides additional character data (stats, description, inventory, etc.), the detail view can be enhanced without changing the navigation pattern.

## Risks / Trade-offs

**Risk: Empty List Handling**
- **Mitigation:** Check list length before rendering, show "No characters" message, disable selection when empty

**Risk: Terminal Size Constraints**
- **Trade-off:** Detail view may need scrolling for small terminals or large character data
- **Mitigation:** Start with simple vertical layout, add scrolling later if needed

**Risk: State Consistency**
- **Trade-off:** Selected index could become invalid if character list changes
- **Mitigation:** Reset or clamp index when characters are reloaded

## Migration Plan

No migration needed - this is a new feature that enhances existing functionality without breaking changes. Users will see:
1. Existing character list display remains the same (backward compatible)
2. New interaction: arrow keys now highlight characters
3. New interaction: enter key opens detail view
4. New interaction: ESC returns to list

Rollback: Simply revert the code changes; no data migration required.

## Open Questions

1. Should the selected character index persist across tab switches? (i.e., if user switches to Sessions tab and back, should the same character still be selected?)
   - **Proposed answer:** Yes, persist the index for better UX

2. What should happen if the backend updates character list while viewing details?
   - **Proposed answer:** For now, accept eventual consistency; handle in future real-time update feature

3. Should we display a loading state when fetching character details (if future backend call required)?
   - **Proposed answer:** Not needed yet since all data is loaded upfront; revisit when lazy-loading details
