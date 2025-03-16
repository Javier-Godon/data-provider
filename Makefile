.PHONY: run generate

run:
	go run app/main.go

generate:
	protoc --proto_path=proto \
	    --go_out=app/usecases/cpu/get_cpu_user_usage/grpc \
	    --go-grpc_out=app/usecases/cpu/get_cpu_user_usage/grpc \
	    proto/get_cpu_user_usage.proto

	protoc --proto_path=proto \
	    --go_out=app/usecases/cpu/get_cpu_system_usage/grpc \
	    --go-grpc_out=app/usecases/cpu/get_cpu_system_usage/grpc \
	    proto/get_cpu_system_usage.proto	
