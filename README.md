# Book service


#### Intall the dependencies
```sh
go get -d
```

#### Generate the gRPC service and message
```sh
protoc -I=proto --go_out=proto --go-grpc_out=proto ./proto/proto.proto
```

#### Run test
```
go test -v ./...
```

#### Run the program
```sh
go run main.go
```

#### Docker build and send to docker hub
```sh
docker build -t ng-auth-service .
docker tag ng-auth-service:latest alaurentino/ng-auth-service:latest
docker push alaurentino/ng-auth-service:latest
```

#### Deploy to kubernetes
```sh
kubectl apply -f .\kube-dev.yaml
```