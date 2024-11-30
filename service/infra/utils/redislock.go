package utils

import (
	"context"
	"e-commerce/service/infra/log"
	"gopkg.in/redis.v5"
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"
)

const (
	randomLen       = 16
	tolerance       = 500 // milliseconds
	millisPerSecond = 1000
	lockCommand     = `if redis.call("GET", KEYS[1]) == ARGV[1] then
    redis.call("SET", KEYS[1], ARGV[1], "PX", ARGV[2])
    return "OK"
else
    return redis.call("SET", KEYS[1], ARGV[1], "NX", "PX", ARGV[2])
end`
	delCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end`
)

// A RedisLock is a redis lock.
type RedisLock struct {
	store   *redis.Client
	seconds uint32
	key     []string
	id      string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewRedisLock returns a RedisLock.
func NewRedisLock(store *redis.Client, key []string) *RedisLock {
	return &RedisLock{
		store: store,
		key:   key,
		id:    GetRandomString(randomLen),
	}
}

// Acquire acquires the lock.
func (rl *RedisLock) Acquire() (bool, error) {
	return rl.AcquireCtx(context.Background())
}

// AcquireCtx acquires the lock with the given ctx.
func (rl *RedisLock) AcquireCtx(ctx context.Context) (bool, error) {
	seconds := atomic.LoadUint32(&rl.seconds)
	val, err := rl.store.Eval(lockCommand, rl.key, rl.id, strconv.Itoa(int(seconds)*millisPerSecond+tolerance)).Result()

	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		log.Errorf("Error on acquiring lock for %s, %s", rl.key, err.Error())
		return false, err
	} else if val == nil {
		return false, nil
	}

	reply, ok := val.(string)
	if ok && reply == "OK" {
		return true, nil
	}

	log.Errorf("Unknown reply when acquiring lock for %s: %v", rl.key, val)
	return false, nil
}

// Release releases the lock.
func (rl *RedisLock) Release() (bool, error) {
	return rl.ReleaseCtx(context.Background())
}

// ReleaseCtx releases the lock with the given ctx.
func (rl *RedisLock) ReleaseCtx(ctx context.Context) (bool, error) {
	resp, err := rl.store.Eval(delCommand, rl.key, rl.id).Result()
	if err != nil {
		return false, err
	}

	reply, ok := resp.(int64)
	if !ok {
		return false, nil
	}

	return reply == 1, nil
}

// SetExpire sets the expiration.
func (rl *RedisLock) SetExpire(seconds int) {
	atomic.StoreUint32(&rl.seconds, uint32(seconds))
}
