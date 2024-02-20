package yetidb

import (
	"context"
)

type FileSession struct {
	active bool
}

func NewFileSession() *FileSession {
	return &FileSession{
		active: true,
	}
}

func (fs *FileSession) Execute(ctx context.Context, f func(tx FileTransaction) (any, error)) (any, error) {
	tx := NewFileTransaction()
	return f(*tx)
}

func (fs *FileSession) Close(ctx context.Context) {
	fs.active = false
}
