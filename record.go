package yeti

type Record interface {
	YetiSerialize() ([]byte, error)
	YetiKey() (string, error)
	YetiName() string
}
