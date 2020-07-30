package rdb

import (
	"context"
	"time"

	"github.com/abdukahhor/streamer/storages"
	"github.com/go-redis/redis/v8"
)

type db struct {
	conn *redis.Client
}

//Close closes redis
func (r *db) Close() error {
	return r.conn.Close()
}

//Connect to Redis cache database,
func Connect(addr string, poolSize int) (storages.Cache, error) {

	c := redis.NewClient(&redis.Options{
		Addr: addr, PoolSize: poolSize,
	})
	//to check the connection with Redis server
	err := c.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return &db{conn: c}, nil
}

func (r db) Get(ctx context.Context, key string) (val string, err error) {
	val, err = r.conn.Get(ctx, key).Result()
	return
}

func (r db) Set(ctx context.Context, key string, val string, timeout time.Duration) error {
	return r.conn.Set(ctx, key, val, timeout).Err()
}

//IsNotFound -
func (r db) IsNotFound(err error) bool {
	return err == redis.Nil
}

func (r db) Del(ctx context.Context, key string) (err error) {
	err = r.conn.Del(ctx, key).Err()
	return
}
