proto:
	cd src/pkg/grpc/proto_type/proto && for file in *.proto; do \
		protoc --go_out=../ --go_opt=paths=source_relative \
		--go-grpc_out=../ --go-grpc_opt=paths=source_relative \
		$$file; \
	done

proto-install:
	PB_REL="https://github.com/protocolbuffers/protobuf/releases" && \
	curl -LO $$PB_REL/download/v30.2/protoc-30.2-linux-x86_64.zip && \
	unzip protoc-30.2-linux-x86_64.zip -d $$HOME/.local

# update to ~/.bashrc
# export PATH="$PATH:$HOME/.local/bin"
