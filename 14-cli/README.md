## CLI with Go and Cobra

### Cobra

- Cobra is a library for creating powerful modern CLI interfaces.
- It is used by popular projects such as Kubernetes, Hugo, and Github CLI.
- It is a library for creating CLI applications in Go.
- It is a CLI framework for Go.

Docs:
  - https://cobra.dev/

### Cobra Installation

- `go install github.com/spf13/cobra-cli@latest`
  - This will install the `cobra-cli` command.

### Cobra Usage

- Initialize a new Cobra project:
  - `cobra-cli init`

- Run the Cobra project:
  - `go run main.go`

- Create a new Cobra command:
  - `cobra-cli add <command-name>`
- Create a new Cobra as a child of another command:
  - `cobra-cli add <command-name> -p <parent-command-name>`

- Build the Cobra project:
  - `go build -o cobra-cli main.go`

- Build the Cobra project with custom name:
  - `go build -o <custom-name> maiin.go`

### Cobra Flags

- Cobra flags are used to pass data to a command.

- Cobra flags can be defined as:
  - Persistent flags
    - Persistent flags are global flags that are available to all commands from the specified command down.
  - Local flags
    - Local flags are flags that are only available to a specific command.