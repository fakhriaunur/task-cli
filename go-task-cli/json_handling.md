## Pseudocode for JSON handling

### function readJSON signature: param: json path, return type: struct, error

- get the fullpath of the json file
- join for the full path
- open the json file
- create a var of the target struct
- create a new json decoder pointer
- decode the json into the target struct with the decoder
- return the struct with nil error


### function writeJSON signature: param: pointer to Task struct, return type: error

- create or open the json file if not exist
- open the pointer to the Task struct
- create a new json encoder pointer
- encode the struct into the json file
- write back the json file with minimum 644 octal permission
- return nil


