package yeti

import "context"

type Transaction interface {
	Read(context.Context) ([]Result, error)
	Write(Record) error
	Commit(context.Context)
	Rollback()
	Close()
}
