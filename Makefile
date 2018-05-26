protoc:
	@cd poll/ && protoc *.proto --go_out=plugins=grpc:. --python_out=plugins=grpc:.