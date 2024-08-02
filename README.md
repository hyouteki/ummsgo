> Minimal User Management Microservice in gRPC Go

## Getting Started
- Install necessary dependencies.
```bash
chmod +x setup.sh
sudo ./setup.sh
```
- Compile the Protocol Buffer file to Go module.
```bash
protoc --go_out=./pb --go-grpc_out=./pb service.proto
```
- Run the server-client setup.
```bash
chmod +x build.sh
sudo ./build.sh
```

## Courtesy
- [go-mode - emacs](https://github.com/dominikh/go-mode.el/)
