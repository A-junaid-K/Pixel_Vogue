# export PATH=$PATH:$HOME/go/bin
path:
	@echo "export PATH=\$$PATH:\$$HOME/go/bin"

gen:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. pkg/grpc/token.proto

run:
	go run cmd/main.go
