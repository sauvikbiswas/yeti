package file

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sauvikbiswas/yeti"
)

type FileSession struct {
	driver    yeti.Driver
	sessionId string
}

func NewFileSession(driver yeti.Driver) *FileSession {
	return &FileSession{
		sessionId: uuid.New().String(),
		driver:    driver,
	}
}

func (fs *FileSession) Execute(ctx context.Context, f func(tx yeti.Transaction) (any, error)) (any, error) {
	if !fs.driver.IsActive() {
		return nil, fmt.Errorf("this FileDriver is inactive")
	}

	if fs.GetSesionId() == "" {
		return nil, fmt.Errorf("cannot execute transaction in an inactive session")
	}

	tx := NewFileTransaction(fs.driver, fs)

	output, err := f(tx)

	if err != nil {
		tx.Rollback(ctx)
	} else {
		tx.Commit(ctx)
	}

	tx.Close(ctx)

	return output, err
}

func (fs *FileSession) GetSesionId() string {
	return fs.sessionId
}

func (fs *FileSession) Close(ctx context.Context) {
	fs.sessionId = ""
}
