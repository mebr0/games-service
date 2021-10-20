# gRPC Server

gRPC server that receives messages about games and send them to websocket endpoint

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
WS_HOST=localhost:8080
```
