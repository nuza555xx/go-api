# GO API

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

This repository contains a collection of Go programs and libraries that
demonstrate the language, standard libraries, and tools.

## Clone the project

```
$ git clone https://github.com/nu-ch/go-api.git
$ cd go-api
```
https://github.com/nu-ch/go-api.git is the canonical Git repository.
It is mirrored at https://github.com/nu-ch/go-api.git.
## [example](example/)

```
$ cd example
$ go build
```
A trivial "Hello, world" program that uses a go-api package.

Command [hello](hello/) covers:

A web server that answers the question: "Is Go 1.x out yet?"

Topics covered:

* Web servers ([net/http](//golang.org/pkg/net/http/))
* Logging ([log](//golang.org/pkg/log/))
* Synchronizing data access between goroutines ([sync](//golang.org/pkg/sync/))
* Exporting server state for monitoring ([expvar](//golang.org/pkg/expvar/))
* Unit and integration tests ([testing](//golang.org/pkg/testing/))
* Dependency injection
* Time ([time](//golang.org/pkg/time/))
