## Go Private

`export GOPRIVATE=github.com/your-username/your-repo-name`

- Test private repo:
    - `export GOPRIVATE=github.com/giovannymassuia/learning-go-test-private-module`

### Authentication ways

- Using `~/.netrc` file:
    - `machine github.com login <your-username> password <your-token>`
    - `chmod 600 ~/.netrc`

- Using `~/.gitconfig` file:
  - `git config --global --add url."
  - `git config --global --add url."https://github.com/your-username/".insteadOf "
  - ```
    [url "ssh://git@github.com/"]
        insteadOf = https://github.com/
    ```

### More resources

- https://www.digitalocean.com/community/tutorials/how-to-use-a-private-go-module-in-your-own-project


## Go Proxy vs Vendor

- Proxy
  - [docs](https://proxy.golang.org)
  - The proxy is a module mirror that maintains a consistent view of modules to serve to users, even as modules are published and unpublished.
  - The proxy caches information about modules, versions, and packages that it has seen, and responds to queries about those modules from go get, go list, go mod, and so on.
  - The proxy does not store the contents of packages.

- Vendor
  - [docs](https://golang.org/ref/mod#go-mod-vendor)
  - The go mod vendor command constructs a directory named vendor in the main module's root directory that contains copies of all packages needed to support builds and tests of packages in the main module.
  - The go command uses the vendor directory to satisfy most dependencies: it uses local copies, even if the corresponding packages are available in a module downloaded from a remote repository.
  - The go command ignores the vendor directory when using modules to compute target versions, build requirements, and so on.

## Vendor

- `go mod vendor`
  - The go mod vendor command constructs a directory named vendor in the main module's root directory that contains copies of all packages needed to support builds and tests of packages in the main module.