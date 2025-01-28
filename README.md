# GoQtgRPC

An exploratory work to embed a Go gRPC server inside a C++ app using QML to interact with the gRPC API.


## Build

```sh
buf generate

cd cposlib && CGO_ENABLED=1 go build -buildmode=c-archive -o cposlib.a ./cposlib.go
```

## Pitfalls
- id conficting with qml id
- the GrpcHttp2Channel hostUri can't be changed
- timestamp.proto imports
