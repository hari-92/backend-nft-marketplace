package distributed_lock

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

type IDistributedLock interface {
	Lock(ctx context.Context, key string, duration time.Duration) (bool, error)
	Release(ctx context.Context, key string) error
}

type DistributedLock struct {
	mux   sync.Mutex
	redis *redis.Client
}

func NewDistributedLock(redisClient *redis.Client) IDistributedLock {
	if redisClient == nil {
		panic("[DistributedLock] Redis client is nil")
	}
	return &DistributedLock{
		redis: redisClient,
	}
}

func (dl *DistributedLock) Lock(ctx context.Context, key string, duration time.Duration) (bool, error) {
	dl.mux.Lock()
	defer dl.mux.Unlock()

	boolCmd := dl.redis.SetNX(ctx, key, "locking", duration)

	if err := boolCmd.Err(); err != nil {
		return false, err
	}
	if !boolCmd.Val() {
		return false, nil
	}
	return true, nil
}

func (dl *DistributedLock) Release(ctx context.Context, key string) error {
	return dl.redis.Del(ctx, key).Err()
}
