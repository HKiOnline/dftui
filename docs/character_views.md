# Character Views

This document describes the user interfaces (views) for displaying Dark Fate charter information.

## Characters Tab

Character tab displays a list of selectable characters the logged in user has access (see [user management](users,md)). If the character is selected it displays the character detail view.

## Character Detail View

Character detail view shows the full character sheet displayed pleasingly. The view has a clear and easy way back to the characters tab. Character -data is loaded from a JSON-file in the db/characters -directory.

## Character Data Model

Character data model is based on the character JSON-format character sheet. The data itself is stored as JSON-files in the db/characters -directory. The character JSON-files are named using the following format: character name where white space is replace by underscores followed by an underscore and the charter unique id. See [character JSON-format](character_json_format.md)). Characters are stored in plain JSON-files loaded when needed. Users and characters are associated via [users.json -file](users.md) in the db-directory.
