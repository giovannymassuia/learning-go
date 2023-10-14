## golang-standards / project-layout
[project-layout](github.com/golang-standards/project-layout) is a set of common templates for Go projects.

## Important folders

- `/internal`: Private application and library code
- `/pkg`: Library code that's ok to use by external applications
- `/cmd`: Main applications for this project
- `/configs`: Configuration file examples or default configs
- `/test`: Additional external test apps and test data
- `/api`: OpenAPI/Swagger specs, JSON schema files, protocol definition files.

## Swagger
[Swagger Go repo](https://github.com/swaggo/swag)
`go install github.com/swaggo/swag/cmd/swag@latest`

- `go install`: Installs the packages named by the import paths, along with their dependencies.

- Adding `go/bin` to your `PATH` variable will allow you to run the `swagger` command from anywhere.
  - `export PATH=$PATH:$(go env GOPATH)/bin`

- `swag init`: Creates a `docs` folder in your project root with `swagger.json` and `swagger.yaml` files.
  - `swag init -g cmd/api/main.go`: Generates `swagger.json` and `swagger.yaml` files by analyzing the `main.go` file.