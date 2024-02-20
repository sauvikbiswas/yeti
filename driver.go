package yeti

import (
	"context"

	"github.com/sauvikbiswas/yeti/config"
)

type Driver interface {
	Configure(config.DriverConfig) error
	NewSession(context.Context, config.SessionConfig) (Session, error)
	Close(context.Context)
}
