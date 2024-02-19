package yeti

import "context"

type Transaction interface {
	Read(context.Context) ([]Result, error)
	Write(Record)
	Commit(context.Context)
	Rollback()
	Close()
}
