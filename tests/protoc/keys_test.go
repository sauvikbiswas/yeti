package protoc

import (
	"log"
	"testing"

	"github.com/sauvikbiswas/yeti/proto/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var logger = log.Default()

func TestProtoWithSinglePrimaryKey(t *testing.T) {
	rec := &test.TestProto{
		Name: "Sauvik",
		Age:  99,
	}

	key, err := rec.YetiKey()

	if err != nil {
		logger.Printf("error fetching key for %s, %s", rec.YetiType(), err.Error())
	}
	require.NoError(t, err)

	assert.Equal(t, key, rec.GetName())
}

func TestProtoWithCompositePrimaryKey(t *testing.T) {
	rec := &test.TestProtoWithCompositeKey{
		Name:        "Sauvik",
		AgeAsString: "99",
	}

	key, err := rec.YetiKey()

	if err != nil {
		logger.Printf("error fetching key for %s, %s", rec.YetiType(), err.Error())
	}
	require.NoError(t, err)

	assert.Equal(t, key, rec.GetName()+rec.GetAgeAsString())
}

func TestProtoWithNonStringPrimaryKey(t *testing.T) {
	rec := &test.TestProtoWithNonStringPrimaryKey{
		Name: "Sauvik",
		Age:  99,
	}

	key, err := rec.YetiKey()

	if err != nil {
		logger.Printf("error fetching key for %s, %s", rec.YetiType(), err.Error())
	}
	require.NoError(t, err)

	assert.Equal(t, key, rec.GetName())
}

func TestProtoWithPrimaryKeyFieldUnset(t *testing.T) {
	rec := &test.TestProtoWithCompositeKey{
		AgeAsString: "99",
	}

	_, err := rec.YetiKey()

	if err != nil {
		logger.Printf("error fetching key for %s, %s", rec.YetiType(), err.Error())
	}
	require.Error(t, err)
}
