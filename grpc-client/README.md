# gRPC Client

gRPC client that generates random messages and send them to grpc server

## Protobuf

Generate code from protobuf:

```shell
make proto
```

## Build & Run

Build project:

```shell
make build
```

Run project (include build):

```shell
make run
```

Example of `.env` file

```dotenv
GRPC_PORT=9000
GRPC_PERIOD=4s
```
