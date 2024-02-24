FILES := $(shell find . -name '*.proto')
OPTION_FILES := $(shell find ./proto/options -name '*.proto')

yeti-proto: $(FILES)
	protoc --proto_path=.  --go_out=. --go_opt paths=source_relative --go-yeti_out=. --go-yeti_opt paths=source_relative $(FILES)

yeti-option: $(OPTION_FILES)
	protoc --proto_path=. --go_out=. --go_opt paths=source_relative $(OPTION_FILES)

install-yeti-plugin: yeti-option
	cd cmd/protoc-gen-go-yeti; go install .

yeti-tests:
	go test -v ./...

tidy:
	cd cmd/protoc-gen-go-yeti; go mod tidy
	go mod tidy

.PHONY: install-yeti-plugin yeti-proto yeti-option yeti-tests tidy