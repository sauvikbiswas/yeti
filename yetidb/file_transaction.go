package yetidb

import (
	"context"

	"github.com/sauvikbiswas/yeti"
)

type FileTransaction struct {
}

func NewFileTransaction() *FileTransaction {
	return &FileTransaction{}
}

func (ft *FileTransaction) Read(ctx context.Context) ([]yeti.Result, error) {
	return nil, nil
}

func (ft *FileTransaction) Write(r yeti.Record) {
}

func (ft *FileTransaction) Commit(ctx context.Context) {
}

func (ft *FileTransaction) Rollback() {
}

func (ft *FileTransaction) Close() {
}
