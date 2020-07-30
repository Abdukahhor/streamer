package app

import (
	"context"
	"time"

	"github.com/abdukahhor/streamer/internal"
	"github.com/abdukahhor/streamer/models"
	"github.com/abdukahhor/streamer/storages"
)

//Core - main layer
type Core struct {
	cache storages.Cache
	cfg   models.Config
}

//New - инициализировать новый экземпляр Core
func New(s storages.Cache, c models.Config) *Core {
	return &Core{cache: s, cfg: c}
}

//GetURL -
func (c Core) GetURL(ctx context.Context) (url string, val string, err error) {

	ix := internal.GetRand(0, c.cfg.Ln)
	url = c.cfg.URLs[ix]
reCheck:
	val, err = c.cache.Get(ctx, url)
	if err == nil {
		return
	}

	if !c.cache.IsNotFound(err) {
		return
	}

	_, err = c.cache.Get(ctx, "txn"+url)
	if err == nil {
		time.Sleep(1 * time.Second)
		goto reCheck
	}

	timeout := internal.GetRand(c.cfg.MinTimeout, c.cfg.MaxTimeout)
	err = c.cache.Set(ctx, "txn"+url, "processing", time.Duration(timeout+60)*time.Second)
	if err != nil {
		return
	}
	defer c.cache.Del(ctx, "txn"+url)

	val, err = internal.GetURL(url)
	if err != nil {
		return
	}
	//cache the response
	err = c.cache.Set(ctx, url, val, time.Duration(timeout)*time.Second)
	return
}
