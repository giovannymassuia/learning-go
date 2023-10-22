## Dependency Injection in Go

### What is Dependency Injection?

Dependency Injection is a technique whereby one object (or static method) supplies the dependencies of another object. A dependency is an object that can be used (a service). An injection is the passing of a dependency to a dependent object (a client) that would use it. The service is made part of the client's state. Passing the service to the client, rather than allowing a client to build or find the service, is the fundamental requirement of the pattern.

### Why Dependency Injection?

Dependency Injection is a software design pattern that allows us to develop loosely coupled code. It allows us to develop components that are independent of each other. It allows us to develop highly cohesive code with minimal coupling.

### How to Implement Dependency Injection in Go

There are two types of dependency injection: constructor injection and method injection.

#### Constructor Injection

Constructor injection is when a dependency is injected into a class via its constructor. The dependency is then stored in a private field. The class can then use the dependency.

#### Method Injection

Method injection is when a dependency is injected into a class via a method. The dependency is then stored in a private field. The class can then use the dependency.

---

### Dependency Injection in Go with libraries

- [Wire](https://github.com/google/wire)
- [Dig](https://github.com/uber-go/dig)
- [Fx](https://github.com/uber-go/fx)

#### Working with Googles Wire

- Install
    - `go install github.com/google/wire/cmd/wire@latest`
- Create a `wire.go` file in the root of your project
- Generate Wire code
    - `wire`
