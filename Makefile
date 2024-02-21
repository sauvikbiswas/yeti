FILES := $(shell find . -name '*.proto')

proto: $(FILES)
	protoc --proto_path=.  --go_out=./tests --go_opt paths=source_relative --go-yeti_out=./tests --go-yeti_opt paths=source_relative $(FILES)

install:
	cd cmd/protoc-gen-go-yeti
	go install .

.PHONY: proto install