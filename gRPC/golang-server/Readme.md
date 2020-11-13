# gRPC in Golang
## Requirements
We need to have installed:
- Protobuf

## Dependencies
We need to have installed to Golang:
```
$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/proto
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

Commands to execute:
```
$ protoc -I.. --go_out=plugins=grpc:cases ../cases.proto
```
or
```
$ protoc -I=.. --go_out=plugins=grpc:cases ../cases.proto
```

*this command is used to compile the **.proto** file and obtain the **.pb.go** file to create the gRPC client*

To run the projecto will be execute
```
$ go run main.go
```