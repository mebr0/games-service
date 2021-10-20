# WebSocket Service

Service that receives messages from websocket endpoint and log them in console

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
GIN_MODE=release    # For prod

HTTP_PORT=8080
HTTP_READ_TIMEOUT=10s
HTTP_WRITE_TIMEOUT=10s
HTTP_MAX_HEADER_MEGABYTES=1
```
