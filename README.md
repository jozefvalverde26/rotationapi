# Rotate Matrix API
- [Getting Started](#getting-started)
  * [Requirements](#requirements)
  * [Installation](#installation)
    + [Check Tools](#check-tools)
    + [Project Dependencies](#project-dependencies)
- [Compile and Run](#compile-and-run)
- [Run Test](#run-test)
- [Example](#example)

## Getting Started

### Requirements

- golang 1.12
- [shellcheck](https://github.com/koalaman/shellcheck#installing)

### Installation

#### Check Tools

In order to make all the checks in this codebase, there are some external go tools that must be installed:

```sh
$ make install-tools
```

**PS:** remember to update your PATH correctly: `export PATH="$GOPATH/bin:$PATH"`

#### Project Dependencies

```sh
$ make install-deps
```

## Compile and Run

### Local Environment with docker
```sh
$ docker-compose up
```

### Local Environment development
* up and running
```sh
$ go run cmd/rotationapi/main.go
```
## Run Test
```sh
$ go test ./...
```

## Example
```sh
$ curl --data '{"matrix": [[1,2], [3,4]]}' -H "content-type: application/json"  -X POST "http://localhost:5000/matrix/rotate"
```
* Response
```json
{"data":[[2,4],[1,3]]}
```
