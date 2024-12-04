gen:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

clean:
	rm -f proto/*.pb.go

script:
	go run script/main.go

run:
	go run cmd/main.go

test:
	go test ./...

.PHONY: gen clean script run test