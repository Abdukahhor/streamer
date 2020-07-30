package storages

import (
	"context"
	"time"
)

//Cache -
type Cache interface {
	Get(ctx context.Context, key string) (val string, err error)
	Set(ctx context.Context, key string, val string, timeout time.Duration) error
	Del(ctx context.Context, key string) (err error)
	Close() error
	IsNotFound(err error) bool
}
