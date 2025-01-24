LOCAL_BIN := $(CURDIR)/bin

install-deps-windows:
	go env -w GOBIN=$(LOCAL_BIN)
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	
install-deps-macOS:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u google.golang.org/grpc/status
	go get -u google.golang.org/protobuf/reflect/protoreflect
	go get -u google.golang.org/protobuf/runtime/protoimpl
	go get -u google.golang.org/protobuf/types/known/emptypb
	go get -u google.golang.org/protobuf/types/known/timestamppb
	go get -u google.golang.org/protobuf/types/known/wrapperspb
	go get -u google.golang.org/grpc

generate-user-api-v1-windows:
	protoc --proto_path api/user_api_v1 \
	--go_out=pkg/user_api_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go.exe \
	--go-grpc_out=pkg/user_api_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc.exe \
	api/user_api_v1/user_api_v1.proto

build:
	go env -w GOOS=linux
	go env -w GOARCH=amd64
	go build -o auth cmd/server/main.go

send:
	pscp -i C:/Users/KanzA/Downloads/putty-0.73-ru-17/PuTTY/123.ppk  $(CURDIR)/auth root@90.156.156.247:

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/kaiwoch/test-server:v0.0.2 .
	docker login -u token --password-stdin < pass.txt cr.selcloud.ru/kaiwoch
	docker push cr.selcloud.ru/kaiwoch/test-server:v0.0.2
