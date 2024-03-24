gen:
	protoc --proto_path=proto proto/*.proto --go_out=server --go-grpc_out=server
	protoc --proto_path=proto proto/*.proto --go_out=client --go-grpc_out=client


clean:
	rm -rf server/pb/
	rm -rf client/pb/

start-server:
	go run server/main.go

start-client:
	go run client/main.go

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	brew install protobuf
	brew install clang-format
	brew install grpcurl
	export PATH=$PATH:$(go env GOPATH)/bin

docker-build:
	sudo docker compose build

docker-run:
	sudo docker compose run --rm client

docker-down:
	sudo docker compose down

test:
	go test -cover -race ./pkg/merkletree/...

