protoc: protoc-sub protoc-reverse

protoc-sub:
	@cd poll/ && protoc -I/usr/local/include -I. \
				-I$(GOPATH)/src \
				-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
				--go_out=plugins=grpc:. poll.proto

protoc-reverse:
	@cd poll/ && protoc -I/usr/local/include -I. \
				-I$(GOPATH)/src \
				-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
				--grpc-gateway_out=logtostderr=true:. \
				poll.proto