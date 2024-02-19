package record

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// JSONRecord is a record that is stored as a JSON string.
type JSONRecord struct {
	ProtoBuf proto.Message
}

// NewJSONRecord creates a new JSONRecord.
func NewJSONRecord(p proto.Message) *JSONRecord {
	return &JSONRecord{
		ProtoBuf: p,
	}
}

// Serialize converts the JSONRecord to a JSON string bytearray.
func (r *JSONRecord) Serialize() ([]byte, error) {
	return protojson.Marshal(r.ProtoBuf)
}
