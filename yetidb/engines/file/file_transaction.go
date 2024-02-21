package file

import (
	"context"
	"os"
	"path"

	"github.com/sauvikbiswas/yeti"
)

type FileTransaction struct {
	driver           yeti.Driver
	session          yeti.Session
	transactionCache map[string]string
}

func NewFileTransaction(driver yeti.Driver, session yeti.Session) *FileTransaction {
	return &FileTransaction{
		driver:           driver,
		session:          session,
		transactionCache: make(map[string]string),
	}
}

func (ft *FileTransaction) Read(ctx context.Context) ([]yeti.Result, error) {
	return nil, nil
}

func (ft *FileTransaction) Write(r yeti.Record) error {
	byteArray, err := r.Serialize()
	if err != nil {
		return err
	}

	key, err := r.Key()
	if err != nil {
		return err
	}

	if err := os.WriteFile(path.Join(ft.driver.GetPath(), key), byteArray, 0666); err != nil {
		return err
	}

	return nil
}

func (ft *FileTransaction) Commit(ctx context.Context) {
}

func (ft *FileTransaction) Rollback() {
}

func (ft *FileTransaction) Close() {
}
