.PHONY: get-protoc-bin
get-protoc-bin:
	go install \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: gen
gen:
	mkdir -p gen/go
	protoc -I . -I proto/third_party/ \
		--go_out=./gen/go --go-grpc_out=./gen/go --grpc-gateway_out=./gen/go \
		proto/example.proto

.PHONY: server
server:
	go run server/main.go

.PHONY: gateway
gateway:
	go run gateway/main.go

.PHONY: echo
echo:
	@curl -X POST localhost:8081/echo -d '{"aMap": {"someString": "test", "emptyString": "", "nullableString": null}}'
	@echo
