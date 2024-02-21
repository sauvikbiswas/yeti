package yeti

type Record interface {
	Serialize() ([]byte, error)
	Key() (string, error)
}
