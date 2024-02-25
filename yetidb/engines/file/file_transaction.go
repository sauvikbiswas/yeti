package file

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"

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

func (ft *FileTransaction) Read(ctx context.Context, r yeti.Record) (map[string]yeti.Record, error) {
	if !ft.driver.IsActive() {
		return nil, fmt.Errorf("this FileDriver is inactive")
	}

	if ft.session.GetSesionId() == "" {
		return nil, fmt.Errorf("cannot execute transaction in an inactive session")
	}

	res := make(map[string]yeti.Record)

	resSession, _ := ft.getRecordFromDatabaseFolder(r)
	for key, record := range resSession {
		res[key] = record
	}

	resTransaction, _ := ft.getRecordFromTransactionFolder(r)
	for key, record := range resTransaction {
		res[key] = record
	}

	return res, nil
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

	yType := r.YetiType()

	if err := os.WriteFile(path.Join(ft.driver.GetPath(), ft.GetTransactionId(), yType+"_"+key+".json"), byteArray, 0666); err != nil {
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

	if err := checkFolder(folder); err != nil {
		return err
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

func (ft *FileTransaction) getRecordFromTransactionFolder(rec yeti.Record) (map[string]yeti.Record, error) {
	folder := path.Join(ft.driver.GetPath(), ft.GetTransactionId())
	return readRecordFromFolder(folder, rec)
}

func (ft *FileTransaction) getRecordFromDatabaseFolder(rec yeti.Record) (map[string]yeti.Record, error) {
	folder := ft.driver.GetPath()
	return readRecordFromFolder(folder, rec)
}

func checkFolder(folder string) error {
	if info, err := os.Stat(folder); errors.Is(err, fs.ErrExist) {
		if !info.IsDir() {
			return fmt.Errorf("%s exists and is not a directory", folder)
		}
	}
	return nil
}

func readRecordFromFolder(folder string, rec yeti.Record) (map[string]yeti.Record, error) {

	if err := checkFolder(folder); err != nil {
		return nil, err
	}

	prefix := rec.YetiType()

	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	res := make(map[string]yeti.Record, 0)

	for _, file := range files {
		if strings.HasPrefix(file.Name(), prefix) {
			if b, err := readFile(path.Join(folder, file.Name())); err == nil {
				rx := rec.New()
				if err := rx.YetiDeserialize(b); err == nil {
					if key, err := rx.YetiKey(); err == nil {
						res[key] = rx
					}
				}
			}
		}
	}

	return res, nil
}

func readFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	b := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(b)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return b, nil
}
