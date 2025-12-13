# Characters JSON-format character sheets

Dark Fate as a role-playing game is character-focused. Generally, characters are divided into two groups: player characters and non-player characters (NPCs). Dark Fate characters use Fate Condensed rules or expanded rules. Character data is transported in JSON format.

## Human Characters

Human characters in Dark Fate use the basic Fate Condensed rules to create the characters.

## Vampire the Eternal Fate Characters

Vampire the Eternal Fate generally has vampire, ghoul, and human characters. Human characters are plain Fate Condensed characters. Ghoul characters are Fate Condensed characters with the permission to take limited supernatural stunts. Vampire characters are expanded Fate Condensed characters with additional properties such as blood potency, disciplines, and an additional hunger stress track.

## Ghoul Characters

Ghoul characters in Dark Fate use the basic Fate Condensed rules to create the characters like humans. However they have some differences in character aspects.

## Character JSON-attributes

Below are the character attributes used in JSON format, their explanation, type, and default value. If the attribute is used with only certain "spirit" or type of characters, it is mentioned; otherwise, the attribute is present in all character types.

  - id: unique identifier of the character used in data storages (string, default uuid v4)
  - player: username of the player, (string, default ""), npc-characters' player uses the username of the gamemaster
  - category: "character" (string, default "character")
  - spirit: "vampire" OR "ghoul" OR "human", (string, default "human") 
  - group: "pc" for player characters OR "npc" for non-player characters, (string, default "npc")
  - name: full name of the character, (string, default "")
  - gender: gender of the character, "male" OR "female" (string, default "male") 
  - aliases: list of name aliases, (array of strings, default [])
  - tags: list of tags for the character, (array of strings, default [])
  - collectives: list of collectives the character is affiliated with, (array of strings, default [])
  - embrace_year: year, positive if after Christ, negative if before (number, default 1982)
  - setting_year: year, positive if after Christ, negative if before (number, default 1982)  "description": "",
  - notes: free form notes about the character (string, default "")
  - refresh: fate point refresh, (number, default 3)
  - bloodPotency: potency of vampire's blood, (number, default 1), only in vampire characters
  - aspects: list of character aspects, (array of objects, see character aspect defaults)
  - skills: list of character skills, (array of objects, see skill defaults)
  - stunts: list of character stunts, (array of objects, see stunt defaults)
  - disciplines: list of vampire disciplines, (array of objects, see discipline defaults), only in vampire characters
  - consequences: list of consequence aspects, (array of objects, see consequence defaults)
  - physicalStressLimit: physical stress slots available for the character, (number, default 3)
  - physicalStressCurrent: physical stress slots used, (number, default 0)
  - mentalStressLimit: mental stress slots available for the character, (number, default 3)
  - mentalStressCurrent: mental stress slots used, (number, default 0)
  - hungerStressLimit: hunger stress slots available for the character, (number, default 3), only in vampire characters
  - hungerStressCurrent: hunger stress slots used, (number, default 0), only in vampire characters

### Character aspect defaults

Fate Condensed tells that "aspects are short phrases that describe who your character is or what is important to them. They can relate to your character’s physical or mental qualities, history, beliefs, training, relationships, or even particularly important equipment".

In Fate Condensed there are four types of aspects: high concept, trouble, relationship aspects and free aspects. These same aspects are used with all characters with the human spirit. Characters with ghoul and vampire spirit have slightly adjusted aspects.

Characters with vampire spirit have high concept, trouble, clan and covenant aspects. Clan aspect is unique to vampires. Optionally vampire character may have also relationship and free aspects.

Characters with ghoul spirit have high concept, trouble and covenant aspects and at least one relationship aspect. The mandatory relationship aspect depicts the ghouls relationship with its master. Optionally ghoul character may have also more relationships and free aspects.

Character aspects are part of the aspects attribute in the JSON-formatted character sheet. It is an array of objects. Each item (object) contains one aspect. Each aspect object has type, title, and description attributes. Possible values for the type attribute are high concept, trouble, relationship, free, clan, and covenant strings. Title is a string used to display the name of the aspect. Description is a free-form string used to explain the effect of the aspect. The description is optional and may not always be present.

Below are the defaults for character aspects and their default values.

**Human character default aspects:**

```json
"aspects":[
    {
      "type": "high concept",
      "title": "",
      "description": ""
    },
    {
      "type": "trouble",
      "title": "",
      "description": ""
    },
    {
      "type": "relationship",
      "title": "",
      "description": ""
    },
    {
      "type": "free",
      "title": "",
      "description": ""
    }
]
```
**Vampire character default aspects:**

```json
"aspects":[
    {
      "type": "high concept",
      "title": "",
      "description": ""
    },
    {
      "type": "trouble",
      "title": "",
      "description": ""
    },
    {
      "type": "clan",
      "title": "",
      "description": ""
    },
    {
      "type": "covenant",
      "title": "",
      "description": ""
    }
]
```

**Ghoul character default aspects:**

```json
"aspects":[
    {
      "type": "high concept",
      "title": "",
      "description": ""
    },
    {
      "type": "trouble",
      "title": "",
      "description": ""
    },
    {
      "type": "covenant",
      "title": "",
      "description": ""
    },
    {
      "type": "relationship",
      "title": "",
      "description": ""
    }
]
```

### Skill defaults

Skills tell what characters can do. Below are the skills and their default ratings. 

Skills in the JSON-format character sheet are an array of objects. Each object contains attributes and values of a single skill. Each skill has title, group, and rating. Title is a string. In the example below are all the values a skill title can have. Skill group is a string; it can have a value of mental, physical, or social. Rating is an integer number. The default is zero, 0.

All character spirit types use the same skills.

Note: Dark Fate uses "larceny" instead of "burglary" as the skill title for the burglary skill.

```json
"skills": [
    {
      "title": "academics",
      "group": "mental",
      "rating": 0
    },
    {
      "title": "athletics",
      "group": "physical",
      "rating": 0
    },
    {
      "title": "contacts",
      "group": "social",
      "rating": 0
    },
    {
      "title": "crafts",
      "group": "mental",
      "rating": 0
    },
    {
      "title": "deceive",
      "group": "social",
      "rating": 0
    },
    {
      "title": "drive",
      "group": "physical",
      "rating": 0
    },
    {
      "title": "empathy",
      "group": "social",
      "rating": 0
    },
    {
      "title": "fight",
      "group": "physical",
      "rating": 0
    },
    {
      "title": "investigate",
      "group": "mental",
      "rating": 0
    },
    {
      "title": "larceny",
      "group": "physical",
      "rating": 0
    },
    {
      "title": "lore",
      "group": "mental",
      "rating": 0
    },
    {
      "title": "notice",
      "group": "mental",
      "rating": 0
    },
    {
      "title": "physique",
      "group": "physical",
      "rating": 0
    },
    {
      "title": "provoke",
      "group": "social",
      "rating": 0
    },
    {
      "title": "rapport",
      "group": "social",
      "rating": 0
    },
    {
      "title": "resources",
      "group": "social",
      "rating": 0
    },
    {
      "title": "shoot",
      "group": "physical",
      "rating": 0
    },
    {
      "title": "stealth",
      "group": "physical",
      "rating": 0
    },
    {
      "title": "technology",
      "group": "mental",
      "rating": 0
    },
    {
      "title": "will",
      "group": "social",
      "rating": 0
    }
]
```

### Stunt defaults

Fate Condensed tells that "stunts are the cool techniques, tricks, or bits of equipment that make characters unique and interesting."

Stunts in the JSON-format character sheet are stored in an array of objects called stunts. Each object represents one stunt. Each stunt has title and description. Title is a string, used as a display name. Description is a free-form string, a longer explanation of what the stunt accomplishes. 

Below is an example of stunts array with one stunt with its title and description empty, which is the default. Each character has typically three stunts. All character spirit types use the same format.

```json
"stunts": [
    {
        "title":"",
        "description":""
    }
]
```

### Discipline defaults

Disciplines are unique to characters with vampire spirit. No other characters have them.

Disciplines in the JSON-format character sheet are stored in an array of objects called disciplines. Each object represents one discipline. Each discipline has title and rating. Title is a string and used as the display name. Rating is a number with zero as the default. Often disciplines are displayed as a title and rating combination. For example, Animalism 1 with the first letter capitalized in UIs.

Below are all the disciplines in their default ratings and titles. Most characters have rating above 0 only in one or three disciplines.

```json
"disciplines": [
    {
      "title": "animalism",
      "rating": 0
    },
    {
      "title": "auspex",
      "rating": 0
    },
    {
      "title": "celerity",
      "rating": 0
    },
    {
      "title": "dominate",
      "rating": 0
    },
    {
      "title": "majesty",
      "rating": 0
    },
    {
      "title": "nightmare",
      "rating": 0
    },
    {
      "title": "obfuscate",
      "rating": 0
    },
    {
      "title": "protean",
      "rating": 0
    },
    {
      "title": "resilience",
      "rating": 0
    },
    {
      "title": "vigor",
      "rating": 0
    },
    {
      "title": "coils of the ascendant",
      "rating": 0
    },
    {
      "title": "coils of the sanguine",
      "rating": 0
    },
    {
      "title": "coils of the wyrm",
      "rating": 0
    },
    {
      "title": "coils of the voivode",
      "rating": 0
    },
    {
      "title": "crúac",
      "rating": 0
    },
    {
      "title": "theban sorcery",
      "rating": 0
    }
]
```

### Consequences defaults

Fate Condensed tells that consequences "are new aspects you write on your character sheet when your character takes a hit, representing the real harm and injury your character suffers".

In JSON-format character sheets, consequences are an array of objects. Each object represents one consequence. Each consequence object has level, isActive, and title attributes. Level is a number; possible values are 2, 4, and 6. The isActive attribute tells if the consequence is in play. It is a boolean with a default of false. Title is a string, used as the display name in UIs. Titles are filled only if the consequence isActive value is true.

The consequences array has by default three consequence objects with levels of 2, 4, and 6. Some arrays may have an additional second level-2 consequence object. All character spirit types use the same format. Below is an example of the typical consequences array.

```json
"consequences": [
    {
      "level": 2,
      "isActive": false,
      "title": ""
    },
    {
      "level": 4,
      "isActive": false,
      "title": ""
    },
    {
      "level": 6,
      "isActive": false,
      "title": ""
    }
]
```

## Example JSON-files

- vampire_character.json - vampire character sheet in JSON-format
- ghoul_character.json - ghoul character sheet in JSON-format
- human_character.json - human character sheet in JSON-format

Note: this document overrides any inconsistencies in the example JSON files listed above.

