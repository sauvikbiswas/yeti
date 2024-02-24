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

func (fd *FileDriver) Configure(dcfg config.DriverConfig) error {
	if dcfg.Path == "" {
		return fmt.Errorf("a FileDriver cannot be configured with an empty path")
	}

	if _, err := os.Stat(dcfg.Path); os.IsNotExist(err) {
		return fmt.Errorf("path %s specified in configuration does not exist or is not accessible", dcfg.Path)
	}

	fd.folder = dcfg.Path
	fd.active = true

	return nil
}

func (fd *FileDriver) NewSession(ctx context.Context, scfg config.SessionConfig) (yeti.Session, error) {
	if !fd.IsActive() {
		return nil, fmt.Errorf("this FileDriver is inactive")
	}
	return NewFileSession(fd), nil
}

func (fd *FileDriver) GetPath() string {
	return fd.folder
}

func (fd *FileDriver) IsActive() bool {
	return fd.active
}

func (fd *FileDriver) Close(_ context.Context) {
	fd.active = false
}
