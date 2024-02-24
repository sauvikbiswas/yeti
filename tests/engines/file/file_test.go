package file

import (
	"context"
	"log"
	"testing"

	"github.com/sauvikbiswas/yeti"
	"github.com/sauvikbiswas/yeti/config"
	"github.com/sauvikbiswas/yeti/proto/test"
	"github.com/sauvikbiswas/yeti/yetidb/engines/file"
	"github.com/stretchr/testify/require"
)

var logger = log.Default()

func TestFileDriver(t *testing.T) {
	fd := file.NewFileDriver()

	err := fd.Configure(config.DriverConfig{Path: "./output"})
	require.NoError(t, err)

	ctx := context.Background()

	session, err := fd.NewSession(ctx, config.SessionConfig{})
	require.NoError(t, err)

	rec1 := &test.TestProto{
		Name: "Sauvik",
		Age:  99,
	}

	_, err = session.Execute(ctx, getTransactionFunction(ctx, rec1))
	if err != nil {
		logger.Printf("error executing session, %s", err.Error())
	}

	require.NoError(t, err)

	rec2 := &test.TestProtoWithCompositeKey{
		Name:        "Sauvik",
		AgeAsString: "99",
	}

	_, err = session.Execute(ctx, getTransactionFunction(ctx, rec2))
	if err != nil {
		logger.Printf("error executing session, %s", err.Error())
	}

	require.NoError(t, err)
}

func TestUnsetKeyError(t *testing.T) {
	fd := file.NewFileDriver()

	err := fd.Configure(config.DriverConfig{Path: "./output"})
	require.NoError(t, err)

	ctx := context.Background()

	session, err := fd.NewSession(ctx, config.SessionConfig{})
	require.NoError(t, err)

	rec1 := &test.TestProto{
		Age: 99,
	}

	_, err = session.Execute(ctx, getTransactionFunction(ctx, rec1))
	if err != nil {
		logger.Printf("error executing session, %s", err.Error())
	}

	require.Error(t, err)

	rec2 := &test.TestProtoWithCompositeKey{
		Name: "Sauvik",
	}

	_, err = session.Execute(ctx, getTransactionFunction(ctx, rec2))
	if err != nil {
		logger.Printf("error executing session, %s", err.Error())
	}

	require.Error(t, err)
}

func TestClosedSession(t *testing.T) {
	fd := file.NewFileDriver()

	err := fd.Configure(config.DriverConfig{Path: "./output"})
	require.NoError(t, err)

	ctx := context.Background()

	session, err := fd.NewSession(ctx, config.SessionConfig{})
	require.NoError(t, err)

	session.Close(ctx)

	rec := &test.TestProto{
		Age: 98,
	}

	_, err = session.Execute(ctx, getTransactionFunction(ctx, rec))
	if err != nil {
		logger.Printf("error executing session, %s", err.Error())
	}

	require.Error(t, err)
}

func TestFolderNotPresent(t *testing.T) {
	fd := file.NewFileDriver()

	err := fd.Configure(config.DriverConfig{Path: "./folder-not-present"})

	if err != nil {
		logger.Printf("error configuring driver, %s", err.Error())
	}

	require.Error(t, err)

}

func getTransactionFunction(ctx context.Context, rec yeti.Record) func(tx yeti.Transaction) (any, error) {
	return func(tx yeti.Transaction) (any, error) {
		logger.Printf("executing transcation %s", tx.GetTransactionId())
		err := tx.Write(ctx, rec)
		return nil, err
	}
}
