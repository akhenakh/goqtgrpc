# GoQtgRPC

An exploratory work to embed a Go gRPC server inside a C++ app using QML to interact with the gRPC API.


## Build

```sh
buf generate

cd poslib && CGO_ENABLED=1 go build -buildmode=c-archive -o poslib.a ./poslib.go
```
