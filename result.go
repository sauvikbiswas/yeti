package yeti

import "context"

type Result interface {
	Err() error
	ThisRecord() *Record
	NextRecord(context.Context) (*Record, bool)
	Collect(context.Context) ([]*Record, error)
}
