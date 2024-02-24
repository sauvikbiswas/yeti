module github.com/sauvikbiswas/yeti/cmd/protoc-gen-go-yeti

go 1.22.0

replace github.com/sauvikbiswas/yeti => ../..

require (
	github.com/sauvikbiswas/yeti v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.32.0
)
