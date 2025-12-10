# Users

Users are the logged-in users of the Dark Fate TUI application. If the username is in the db/users.json file, they potentially have access to chronicles, campaigns, and characters information.


## Users JSON-file

Below is an example of the users.json file. It contains a JSON array of objects. Each object is a representation of one user. The username attribute is the same as the login username. 
The chronicles array contains strings of chronicle identifiers; if the user has an identifier that matches a chronicle in the chronicles.json file, they can see information about the chronicle in the chronicles tab. 

```json
[
    {
        "username": "",
        "chronicles": [],
        "campaigns": [],
        "characters": []
    }
[
```
