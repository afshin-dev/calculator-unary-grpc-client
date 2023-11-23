compile_proto:
	 protoc -IPATH="." --go_out="."  --go-grpc_out="." .\protod\calculator.proto


.PHONY: compile_proto