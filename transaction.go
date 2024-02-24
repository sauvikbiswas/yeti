package yeti

import "context"

type Transaction interface {
	Read(context.Context) ([]Result, error)
	Write(context.Context, Record) error
	Commit(context.Context) error
	Rollback(context.Context) error
	Close(context.Context)
	GetTransactionId() string
}
