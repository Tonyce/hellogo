gen-pbs:
	@cd proto && protoc --go_out=paths=source_relative:../proto_gen helloworld/*.proto \
		--go-grpc_out=paths=source_relative:../proto_gen helloworld/*.proto 
