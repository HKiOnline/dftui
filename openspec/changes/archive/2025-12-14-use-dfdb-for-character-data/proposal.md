# Use DFDB for Character Data

## Summary
Replace the current mock character data in the Characters tab with actual character data loaded from the dflib/dfdb package. The system will read character JSON files from the `db/characters` directory and display them based on the logged-in user's username.

## Motivation
Currently, the Characters tab uses stub/mock data. This proposal implements real character data loading using the existing dfdb filesystem provider, allowing users to see their actual characters (PCs) and NPCs visible to them in the game.

## Goals
- Load character data from `db/characters` directory using dfdb
- Display PCs for the logged-in user (matching player name)
- Display all NPCs (for gamemaster visibility)
- Show character details in an organized view
- Maintain backward compatibility with existing UI structure

## Non-Goals
- User authentication integration (assume username is available from SSH session)
- Character creation or editing functionality
- Multi-user concurrent edits to character data
