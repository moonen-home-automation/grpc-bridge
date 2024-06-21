proto-gen:
	git submodule update --init --recursive
	protoc --go_out=. --go-grpc_out=. proto/hass.proto