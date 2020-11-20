# Book service


#### Intall the dependencies
```sh
go get -d
```


#### Generate the gRPC service and message
```sh
protoc -I=proto --go_out=proto --go-grpc_out=proto ./proto/proto.proto
```


#### Run the program
```sh
go run main.go
```