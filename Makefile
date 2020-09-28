# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
#latest

download:
	export GOPROXY=https://goproxy.io
	$(GOCMD) mod download

build:
	$(GOBUILD) cmd/main.go

run:
	cd internal/models/proto && protoc --go_out=../protoCompiles ./blogChatRecord.proto
	cd internal/models/proto && protoc --go_out=../protoCompiles ./blogArticle.proto
	cd internal/models/proto && protoc --go_out=../protoCompiles ./blogComment.proto
	cd internal/models/proto && protoc --go_out=../protoCompiles ./blogType.proto
	cd internal/models/proto && protoc --go_out=../protoCompiles ./blogView.proto
	cd internal/models/proto && protoc --go_out=../protoCompiles ./errors.proto
	cd internal/models/proto && protoc --go_out=../protoCompiles ./resp.proto
	cd internal/models/proto && protoc --go_out=../protoCompiles ./ws.proto

	cd internal/models/proto && protoc --js_out=import_style=commonjs,binary:../protoCompiles/js ./blogChatRecord.proto
	cd internal/models/proto && protoc --js_out=import_style=commonjs,binary:../protoCompiles/js ./blogArticle.proto
	cd internal/models/proto && protoc --js_out=import_style=commonjs,binary:../protoCompiles/js ./blogComment.proto
	cd internal/models/proto && protoc --js_out=import_style=commonjs,binary:../protoCompiles/js ./blogType.proto
	cd internal/models/proto && protoc --js_out=import_style=commonjs,binary:../protoCompiles/js ./blogView.proto
	cd internal/models/proto && protoc --js_out=import_style=commonjs,binary:../protoCompiles/js ./errors.proto
	cd internal/models/proto && protoc --js_out=import_style=commonjs,binary:../protoCompiles/js ./resp.proto
	cd internal/models/proto && protoc --js_out=import_style=commonjs,binary:../protoCompiles/js ./ws.proto


	export GOPROXY=http://goproxy.io
	$(GOCMD) run cmd/main.go