### Go Mod Init
`go mod init github.com/your-username/your-repo-name`

### Go Mod Tidy
`go mod tidy`

- This will add any missing dependencies to your go.mod file and remove any unused dependencies from your `go.mod` file.
- It will also download the dependencies to your `go.sum` file.

`go.sum`

- This file contains the expected cryptographic checksums of the content of specific module versions.
- It a lock file for your dependencies.
- It is intended to be checked into source control.

`go get <package-name>`

- This will add a dependency to your `go.mod` file.
- It will also download the dependency to your `go.sum` file.
- It will not remove any unused dependencies from your `go.mod` file.

`go get -u`

- This will update all of your dependencies to the latest minor or patch versions.
- It will not update to the latest major version.
- It will also update your `go.sum` file.

`go get -u=patch`

- This will update all of your dependencies to the latest patch versions.
- It will not update to the latest minor or major versions.

`go get -u=patch github.com/your-username/your-repo-name`

- This will update a specific dependency to the latest patch version.
- It will not update to the latest minor or major version.

