## gRPC in Go

### Client-Server communication patterns

- Unary
  - Client sends a single request to the server and gets a single response back.
- Server Streaming
  - Client sends a single request to the server and gets a stream of responses back.
- Client Streaming
  - Client sends a stream of requests to the server and gets a single response back.
- Bidirectional Streaming
  - Client sends a stream of requests to the server and gets a stream of responses back.

### Protocol Buffers and proto files

- Protocol Buffers are a language-neutral, platform-neutral, extensible way of serializing structured data for use in communications protocols, data storage, and more.
- Proto files are used to define the structure of the data that we want to serialize using Protocol Buffers.

### Official docs

- [gRPC](https://grpc.io/)
    - [gRPC Go](https://grpc.io/docs/languages/go/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers/)


### Plugins to work with Protocol Buffers and gRPC in Go

- protoc
  - `$ brew install protobuf`
  - https://grpc.io/docs/protoc-installation/

- protoc-gen-go
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`
- protoc-gen-go-grpc
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`

### Code generation

- From the root of the project:
  - `protoc --go_out=. --go-grpc_out=. ./proto/course_category.proto`