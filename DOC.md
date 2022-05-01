# REPL KV store

## Code Structure

- `internal`: Holds the main code base for the application
- [cmd](./internal/cmd): Holds the type of commands and implementation of execution of those commands
- [factory](./internal/factory): Is the wrapper implementing factory pattern to return the desired type datastore
- [store](./internal/store): Holds the packages of different type of datastores.

## Run the application

Run `make run`, it will start the application, then try out the commands.

## Development

- The commands are added in package [cmd](./internal/cmd) which implements the interface [Executor](./internal/cmd/types.go/Executor).

- The constructor or a new commands are added in the [GetCommandExecutor](./internal/cmd/types.go/GetCommandExecutor).

- To support a new data store instead of inmemory, we need to implement [Store](./internal/store/interface.go/Store) interface, and the package in [store](./internal/store) package.
Also, add the wrapper for that datastore in the [factory](./internal/factory/store.go)

### Some Important Commands

#### Run Unit Tests

``make test``

#### Build the application

``make build``

#### Run linter

``make lint``


