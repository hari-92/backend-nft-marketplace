proto:
	cd src/pkg/grpc/proto_type/proto && for file in *.proto; do \
		protoc --go_out=../ --go_opt=paths=source_relative \
		--go-grpc_out=../ --go-grpc_opt=paths=source_relative \
		$$file; \
	done
