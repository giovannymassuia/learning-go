package main

// run: go run main.go

// build: go build main.go
// run compiled: ./main

// build for windows: GOOS=windows GOARCH=386 go build main.go
// run compiled: main.exe

// build for linux: GOOS=linux GOARCH=386 go build main.go
// run compiled: ./main

// build for mac: GOOS=darwin go build main.go
// run compiled: ./main

// GOOS: linux, darwin, windows, freebsd, openbsd, plan9, solaris
//   - linux: linux
//   - darwin: mac
//   - windows: windows

// GOARCH: 386, amd64, arm, arm64
//   - 386: 32 bits
//   - amd64: 64 bits
//   - arm: arm 32 bits
//   - arm64: arm 64 bits

// default of the machine: go env GOOS GOARCH

// `go build`: works when there is a module
//     - this way the compiled will be named as the module name

// `go build -o`: specify the name of the compiled

func main() {
	println("Hello World")
}
