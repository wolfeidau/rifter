/*
Redis backend based on counters, this uses key expiry to
cleanup old entries.
*/
package backend

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/mailgun/vulcan/timeutils"
)

type RedisConfig struct {
	Prefix   string
	Server   string // host:port of server
	Password string
}

type RedisBackend struct {
	Config       *RedisConfig
	TimeProvider timeutils.TimeProvider
	Pool         *redis.Pool
	Expires      time.Duration
}

func NewRedisBackend(config *RedisConfig, timeProvider timeutils.TimeProvider) (*RedisBackend, error) {
	backend := &RedisBackend{
		Config:       config,
		TimeProvider: timeProvider,
		Expires:      time.Hour * 48, // default redis key expiry
	}
	backend.initialize()

	return backend, nil
}

func (b *RedisBackend) GetCount(key string, period time.Duration) (int64, error) {

	conn := b.Pool.Get()
	defer conn.Close()

	redisKey := buildRedisKey(b, key, period)

	if exists, err := redis.Bool(conn.Do("EXISTS", redisKey)); err != nil {
		return 0, err
	} else {
		if exists {
			if counter, err := redis.Int64(conn.Do("GET", redisKey)); err != nil {
				return 0, err
			} else {
				return counter, nil
			}
		}
		return 0, nil
	}

}

func (b *RedisBackend) UpdateCount(key string, period time.Duration, increment int64) error {

	conn := b.Pool.Get()
	defer conn.Close()

	redisKey := buildRedisKey(b, key, period)

	conn.Send("MULTI")
	conn.Send("INCRBY", redisKey, increment)
	// expire key based on the configured duration
	conn.Send("EXPIREAT", redisKey, time.Now().Add(b.Expires).Unix())

	_, err := conn.Do("EXEC")
	if err != nil {
		return err
	}
	return nil
}

func (b *RedisBackend) initialize() {
	b.Pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", b.Config.Server)
			if err != nil {
				return nil, err
			}
			if b.Config.Password != "" {
				if _, err := c.Do("AUTH", b.Config.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (b *RedisBackend) DeleteCount(key string, period time.Duration) (int64, error) {
	conn := b.Pool.Get()
	defer conn.Close()
	redisKey := buildRedisKey(b, key, period)

	return redis.Int64(conn.Do("DEL", redisKey))
}

func (b *RedisBackend) UtcNow() time.Time {
	return b.TimeProvider.UtcNow()
}

func buildRedisKey(b *RedisBackend, key string, period time.Duration) string {
	return b.Config.Prefix + ":" + timeutils.GetHit(b.UtcNow(), key, period)
}
