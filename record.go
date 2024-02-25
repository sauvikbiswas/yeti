package yeti

type Record interface {
	New() Record
	YetiSerialize() ([]byte, error)
	YetiDeserialize([]byte) error
	YetiKey() (string, error)
	YetiType() string
}
