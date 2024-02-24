package file

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/google/uuid"
	"github.com/sauvikbiswas/yeti"
)

type FileTransaction struct {
	driver           yeti.Driver
	session          yeti.Session
	transactionId    string
	transactionCache map[string]string
}

func NewFileTransaction(driver yeti.Driver, session yeti.Session) *FileTransaction {
	return &FileTransaction{
		driver:           driver,
		session:          session,
		transactionId:    uuid.New().String(),
		transactionCache: make(map[string]string),
	}
}

func (ft *FileTransaction) Read(ctx context.Context) ([]yeti.Result, error) {
	return nil, nil
}

func (ft *FileTransaction) Write(ctx context.Context, r yeti.Record) error {
	if !ft.driver.IsActive() {
		return fmt.Errorf("this FileDriver is inactive")
	}

	if ft.session.GetSesionId() == "" {
		return fmt.Errorf("cannot execute transaction in an inactive session")
	}

	if err := ft.makeTransactionFolderIfNotPresent(); err != nil {
		return err
	}

	byteArray, err := r.YetiSerialize()
	if err != nil {
		return err
	}

	key, err := r.YetiKey()
	if err != nil {
		return err
	}

	name := r.YetiName()

	if err := os.WriteFile(path.Join(ft.driver.GetPath(), ft.GetTransactionId(), name+"_"+key+".json"), byteArray, 0666); err != nil {
		return err
	}

	return nil
}

func (ft *FileTransaction) Commit(ctx context.Context) error {
	return ft.moveDataFromTransactionFolder()
}

func (ft *FileTransaction) Rollback(ctx context.Context) error {
	folder := path.Join(ft.driver.GetPath(), ft.GetTransactionId())
	return os.RemoveAll(folder)
}

func (ft *FileTransaction) Close(ctx context.Context) {
}

func (ft *FileTransaction) GetTransactionId() string {
	return ft.transactionId
}

func (ft *FileTransaction) makeTransactionFolderIfNotPresent() error {
	folder := path.Join(ft.driver.GetPath(), ft.GetTransactionId())

	if info, err := os.Stat(folder); errors.Is(err, fs.ErrExist) {
		if !info.IsDir() {
			return fmt.Errorf("%s exists and is not a directory", folder)
		}
	}

	return os.Mkdir(path.Join(ft.driver.GetPath(), ft.transactionId), os.ModePerm)
}

func (ft *FileTransaction) moveDataFromTransactionFolder() error {
	folder := path.Join(ft.driver.GetPath(), ft.GetTransactionId())

	entries, err := os.ReadDir(folder)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		err := os.Rename(path.Join(folder, entry.Name()), path.Join(ft.driver.GetPath(), entry.Name()))
		if err != nil {
			return err
		}
	}

	return os.RemoveAll(folder)
}
