package record

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/sauvikbiswas/yeti/tests/protostruct"
)

func TestNewJSONRecord(t *testing.T) {
	testProto := &protostruct.TestProto{
		Name: "Test",
		Age:  30,
	}

	jsonRecord := NewJSONRecord(testProto)
	assert.Equal(t, testProto, jsonRecord.ProtoBuf)
}

func TestJSONRecord_Serialize(t *testing.T) {
	testProto := &protostruct.TestProto{
		Name: "Test",
		Age:  30,
	}

	jsonRecord := NewJSONRecord(testProto)

	serialized, err := jsonRecord.Serialize()
	assert.NoError(t, err)

	// Deserialize into a map to check contents
	decoded := &structpb.Struct{}
	err = protojson.Unmarshal(serialized, decoded)
	assert.NoError(t, err)

	assert.Equal(t, "Test", decoded.GetFields()["name"].GetStringValue())
	assert.Equal(t, float64(30), decoded.GetFields()["age"].GetNumberValue())
}
