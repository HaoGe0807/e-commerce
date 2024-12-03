# e-commerce
this is a study project

protoc -I grpc/ grpc/eCommerce.proto --go_out=./grpc/ --go-grpc_out=require_unimplemented_servers=false:./grpc/
