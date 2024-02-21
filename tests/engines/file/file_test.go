package file

import (
	"context"
	"testing"

	"github.com/sauvikbiswas/yeti"
	"github.com/sauvikbiswas/yeti/config"
	"github.com/sauvikbiswas/yeti/tests/protostruct"
	"github.com/sauvikbiswas/yeti/yetidb/engines/file"
	"github.com/stretchr/testify/require"
)

func TestFileDriver(t *testing.T) {
	fd := file.NewFileDriver()

	err := fd.Configure(config.DriverConfig{Path: "./output"})
	require.NoError(t, err)

	session, err := fd.NewSession(context.Background(), config.SessionConfig{})
	require.NoError(t, err)

	rec := &protostruct.TestProto{
		Name: "Sauvik",
		Age:  99,
	}

	_, err = session.Execute(context.Background(),
		func(tx yeti.Transaction) (any, error) {
			err := tx.Write(rec)
			return nil, err
		},
	)

	require.NoError(t, err)

}
