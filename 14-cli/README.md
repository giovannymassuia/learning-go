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

- Build the Cobra project:
  - `go build -o cobra-cli main.go`

- Build the Cobra project with custom name:
  - `go build -o <custom-name> main.go`