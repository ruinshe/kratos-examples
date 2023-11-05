package apis

//go:generate protoc -I. --go_out=paths=source_relative:. --go-http_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. apis.proto
