# GKE log parser
Because JSON logs suck.

## Running
1. Pull code locally 
2. Either build the binary or just use `go run`
3. The full command will be `go run main.go FILENAME SEARCH_STRING`
    e.g. `go run main.go file.json error`