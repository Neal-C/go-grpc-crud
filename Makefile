proto-codegen:
	protoc --proto_path=proto-stuff proto-stuff/*.proto --go_out=. --go-grpc_out=.

# optional fields are behind an experimental flag 
# --experimental_allow_proto3_optional
# --go-grpc_out=. will generate gRPC code in the indicated folder specified in quote.proto