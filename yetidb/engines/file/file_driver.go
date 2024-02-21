package file

import (
	"context"
	"fmt"
	"os"

	"github.com/sauvikbiswas/yeti"
	"github.com/sauvikbiswas/yeti/config"
)

type FileDriver struct {
	folder string
	active bool
}

func NewFileDriver() *FileDriver {
	return &FileDriver{}
}

func (f *FileDriver) Configure(dcfg config.DriverConfig) error {
	if dcfg.Path == "" {
		return fmt.Errorf("a FileDriver cannot be configured with an empty path")
	}

	if _, err := os.Stat(dcfg.Path); os.IsNotExist(err) {
		return fmt.Errorf("path %s specified in configuration does not exist or is not accessible", dcfg.Path)
	}

	f.folder = dcfg.Path
	f.active = true

	return nil
}

func (f *FileDriver) NewSession(ctx context.Context, scfg config.SessionConfig) (yeti.Session, error) {
	if !f.active {
		return nil, fmt.Errorf("this FileDriver is inactive")
	}
	return NewFileSession(f), nil
}

func (f *FileDriver) GetPath() string {
	return f.folder
}

func (f *FileDriver) Close(_ context.Context) {
}
