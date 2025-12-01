# SSH Server Capability

## ADDED Requirements

### Requirement: SSH Server Initialization
The system SHALL provide an SSH server that accepts incoming connections on a configurable port.

#### Scenario: Server starts successfully
- **WHEN** the application is launched with default settings
- **THEN** an SSH server SHALL listen on port 2222
- **AND** the server SHALL generate a host key if one does not exist

#### Scenario: Custom port configuration
- **WHEN** the application is launched with `--port 3000` flag
- **THEN** the SSH server SHALL listen on port 3000

#### Scenario: Host key persistence
- **WHEN** the application starts for the first time
- **THEN** a host key SHALL be generated and stored in `~/.dftui/id_rsa`
- **AND** subsequent starts SHALL reuse the existing host key

### Requirement: User Identification
The system SHALL identify users from their SSH session credentials.

#### Scenario: User connects via SSH
- **WHEN** a user connects with username "alice"
- **THEN** the system SHALL extract "alice" as the current user
- **AND** this username SHALL be available throughout the session

#### Scenario: Multiple concurrent users
- **WHEN** multiple users connect simultaneously
- **THEN** each SSH session SHALL maintain its own user context
- **AND** sessions SHALL NOT interfere with each other

### Requirement: Graceful Shutdown
The system SHALL handle shutdown signals cleanly.

#### Scenario: SIGINT received
- **WHEN** the server receives SIGINT (Ctrl+C)
- **THEN** it SHALL stop accepting new connections
- **AND** allow existing sessions to complete
- **AND** exit cleanly within 5 seconds

#### Scenario: SIGTERM received
- **WHEN** the server receives SIGTERM
- **THEN** it SHALL perform graceful shutdown
- **AND** exit cleanly within 5 seconds
