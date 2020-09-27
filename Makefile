# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
#latest
all: test build
update:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOCMD) build -o server_simple cmd/main.go
	mkdir cmd/linux
	mv ./server_simple cmd/linux
	scp cmd/linux/server_simple root@140.143.188.219:/www/bin/server_simple
	rm -rf cmd/linux
test:
	$(GOTEST) -v ./...

clean:
	rm -rf target/

download:
	export GOPROXY=https://goproxy.io
	$(GOCMD) mod download

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

build:
    $(GOBUILD) cmd/main.go

stop:
	pkill -f target/logic
	pkill -f target/job
	pkill -f target/comet
