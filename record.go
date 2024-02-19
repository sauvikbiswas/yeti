package yeti

type Record interface {
	Serialize() ([]byte, error)
	Deserialize([]byte) error
}
