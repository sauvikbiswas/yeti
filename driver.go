package yeti

import (
	"context"

	"github.com/sauvikbiswas/yeti/config"
)

type Driver interface {
	Configure(config.DriverConfig)
	NewSession(context.Context, config.SessionConfig) error
	Close(context.Context)
}
