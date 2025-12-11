# Design: Basic UI Skeleton

## Context
This is the initial implementation of dftui, establishing the foundational architecture for an SSH-based terminal UI. We need to make key decisions about project structure, state management, and how Charm Wish and Bubble Tea integrate.

**Constraints:**
- Must use Charm Wish for SSH server functionality
- Bubble Tea (part of Charm ecosystem) for TUI state management
- Keep it simple - this is just the skeleton, not full functionality
- User identification comes from SSH login (not separate auth)

**Stakeholders:**
- Future developers extending the UI
- Users connecting via SSH

## Goals / Non-Goals

**Goals:**
- Working SSH server that accepts connections on configurable port
- User identification from SSH session username
- Tabbed interface with keyboard navigation (tab key or arrow keys)
- Characters tab displaying a list with PC/NPC indicators
- Clean separation between UI, models, and backend services
- Extensible structure for adding remaining tabs later

**Non-Goals:**
- Real backend integration (stubs only)
- Full functionality of any tab
- Database or persistent storage
- Advanced SSH features (SFTP, port forwarding, etc.)
- User authentication beyond SSH key-based login

## Decisions

### Project Structure
**Decision:** Use standard Go project layout with separation by concern
```
dftui/
├── main.go              # Entry point, SSH server setup
├── go.mod
├── ui/
│   ├── model.go         # Main Bubble Tea model
│   ├── tabs.go          # Tab navigation logic
│   └── characters.go    # Characters tab view
├── models/
│   └── character.go     # Data models
└── services/
    └── backend.go       # Backend service stubs
```

**Rationale:** 
- Keeps UI logic separate from data models
- Services layer provides clear integration points
- Flat structure (no deep nesting) appropriate for early stage
- Easy to understand and extend

**Alternatives considered:**
- Feature-based structure (characters/, sessions/, etc.) - Too early, only one tab implemented
- Monolithic single file - Would become unmaintainable quickly

### Charm Framework Integration
**Decision:** Use Wish for SSH + Bubble Tea for TUI state management

**Rationale:**
- Wish provides SSH server with session management built-in
- Bubble Tea is the standard TUI framework from Charm (Elm architecture)
- Well-integrated ecosystem with Lip Gloss for styling later
- Good middleware support for extending functionality

**Pattern:**
```go
// main.go - SSH server with Wish
s, err := wish.NewServer(
    wish.WithAddress(":2222"),
    wish.WithMiddleware(
        bubbletea.Middleware(teaHandler),
        // ... other middleware
    ),
)

// teaHandler creates the Bubble Tea program for each SSH session
func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
    username := s.User()
    return NewMainModel(username), []tea.ProgramOption{tea.WithAltScreen()}
}
```

### Tab Navigation
**Decision:** Use a simple index-based tab model with keyboard shortcuts

**Tabs (5 total, implement 1 now):**
1. Characters (implemented)
2. Sessions (placeholder)
3. Chronicles (placeholder)
4. Campaigns (placeholder)
5. Fate Tracker (placeholder)

**Navigation:**
- Tab key or Right Arrow: Next tab
- Shift+Tab or Left Arrow: Previous tab
- Number keys 1-5: Direct tab access
- q or Ctrl+C: Quit

**Rationale:**
- Simple and familiar to terminal users
- Easy to implement with Bubble Tea's key message handling
- Extensible as we add more tabs

### User Identification
**Decision:** Extract username directly from SSH session, no additional authentication

**Implementation:**
```go
username := s.User() // From ssh.Session
```

**Rationale:**
- User has already authenticated via SSH keys
- Simplifies initial implementation
- Username can be used to query backend for user's data
- Aligns with SSH-based access model

**Future consideration:** May need user ID mapping if SSH username ≠ backend user identifier

### Mock Data Strategy
**Decision:** Hardcode mock data in backend stubs with TODO comments

**Example:**
```go
// GetUserCharacters returns the characters visible to a user
// TODO: Implement actual backend integration
func GetUserCharacters(username string) ([]models.Character, error) {
    // Mock data for now
    return []models.Character{
        {Name: "Gandalf", Type: "PC", ID: "1"},
        {Name: "Shadowfax", Type: "NPC", ID: "2"},
    }, nil
}
```

**Rationale:**
- Allows UI development without backend dependency
- Clear TODO markers for future implementation
- Easy to test UI rendering

## Risks / Trade-offs

### Risk: SSH Host Key Management
**Issue:** SSH server needs a host key. Where should it be stored?

**Mitigation:** 
- Generate on first run if not present
- Store in `~/.dftui/id_rsa` (configurable via flag)
- Document in README

### Risk: Bubble Tea Learning Curve
**Issue:** Team may not be familiar with Elm architecture

**Mitigation:**
- Keep initial model simple
- Add inline comments explaining Update/View pattern
- Link to Bubble Tea docs in code comments

### Trade-off: Mock Data vs. Backend Interface
**Decision:** Define interface now, implement later

**Pro:** 
- Clear contract for future backend work
- UI can evolve independently

**Con:**
- Interface may need changes when real backend is integrated

**Acceptance:** Expected and acceptable for skeleton phase

## Migration Plan

N/A - This is the initial implementation with no existing code to migrate.

## Open Questions

1. **SSH Port:** What port should the SSH server listen on by default?
   - Suggestion: 2222 (configurable via flag)
   - Standard SSH (22) requires root, avoid for dev

2. **Configuration:** Should we use config file or just flags?
   - Suggestion: Start with flags, add config file later if needed
   - Keep it simple for now

3. **Logging:** What logging library/approach?
   - Suggestion: Use standard `log` package initially
   - Can add structured logging (e.g., zerolog) later if needed

4. **Character List Pagination:** How many characters to show per page?
   - Suggestion: Defer until we have real data and know typical sizes
   - Start with simple scrollable list
