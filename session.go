package yeti

import "context"

type Session interface {
	Execute(context.Context, func(tx Transaction) (any, error)) (any, error)
	Close(context.Context)
}
