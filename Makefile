protoc: protoc-sub protoc-reverse protoc-swagger

protoc-sub:
	@cd poll/ && protoc -I/usr/local/include -I. \
				-I$(GOPATH)/src \
				-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
				--go_out=plugins=grpc:. \
				poll.proto

protoc-reverse:
	@cd poll/ && protoc -I/usr/local/include -I. \
				-I$(GOPATH)/src \
				-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
				--grpc-gateway_out=logtostderr=true:. \
				poll.proto

protoc-swagger:
	@cd poll/ && protoc -I/usr/local/include -I. \
				-I$(GOPATH)/src \
				-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
				--swagger_out=logtostderr=true:. \
				poll.proto