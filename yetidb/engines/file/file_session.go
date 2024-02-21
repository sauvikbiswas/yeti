package file

import (
	"context"

	"github.com/sauvikbiswas/yeti"
)

type FileSession struct {
	driver yeti.Driver
	active bool
}

func NewFileSession(driver yeti.Driver) *FileSession {
	return &FileSession{
		active: true,
		driver: driver,
	}
}

func (fs *FileSession) Execute(ctx context.Context, f func(tx yeti.Transaction) (any, error)) (any, error) {
	tx := NewFileTransaction(fs.driver, fs)
	return f(tx)
}

func (fs *FileSession) Close(ctx context.Context) {
	fs.active = false
}
