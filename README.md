# Skii Resort Daemon
the Skii Resort Daemon is a simple gRPC server that can traverse a graph and provide all possible routes from a given point given a specific diffculty

## Requirements
- `go 1.19`
## Development requirements
if you want to develop for this project you'll need
- [Protoc](https://grpc.io/docs/protoc-installation/)
 you will also need to install protoc-gen plugins
 ```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
you can generate the protobuffer files with `make generate-protoc`


## Installation
With go it's fairly simple to install by using the native install command

```bash
$ go install github.com/edjroz/skii/cmd/skiid
```

## Usage
Once installed you can call the CLI.

```bash
$ skiid help

Skii is a compute engine that can retrieve all available paths for a skiier from a given point based on their difficulty as measured descending (black|red|blue)

Usage:
  skii [command]

Available Commands:
  client      client to interact with gRPC server
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  start       starts skii daemon

Flags:
  -h, --help   help for skii

Use "skii [command] --help" for more information about a command.
```
 
### Start
To start the server.
```bash
$ skiid start --path ./skii-resort.gsv
```
you can pass in any different gsv graph with different values to play around with it

### Interact with Server
The CLi comes in equipped with a quick CLI to interact with the app.

```bash
$ skiid client query --color black --start a
2022/11/09 00:00:00 Routes: route:{node:"a" node:"b"} route:{node:"a" node:"c"}
``

