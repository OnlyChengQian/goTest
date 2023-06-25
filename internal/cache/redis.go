package cache

import (
	"advt/internal/file"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"sync"
	"time"
)

type RedisPoolProvider interface {
	GetRedisConnection() *redis.Pool
}

type RedisProvider struct {
	ConfigReader file.ConfigReader
	redisPool    *redis.Pool
	mutex        sync.Mutex
}

func NewRedisPoolProvider(configReader file.ConfigReader) *RedisProvider {
	return &RedisProvider{
		ConfigReader: configReader,
	}
}

func (r *RedisProvider) LazyInitPool() {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.redisPool == nil {
		config := r.ConfigReader.GetRedisConfig()
		maxIdle, _ := strconv.Atoi(config.MaxIdle)
		maxActive, _ := strconv.Atoi(config.MaxActive)
		idleTimeout, _ := strconv.Atoi(config.IdleTimeout)
		pass := redis.DialPassword(config.Auth)
		db := redis.DialDatabase(0)
		r.redisPool = &redis.Pool{
			MaxIdle:     maxIdle,
			MaxActive:   maxActive,
			IdleTimeout: time.Duration(idleTimeout),
			Dial: func() (conn redis.Conn, err error) {
				return redis.Dial("tcp", config.Host+":"+config.Port, db, pass)
			},
		}
	}
}

func (r *RedisProvider) GetRedisConnection() *redis.Pool {
	r.LazyInitPool()
	return r.redisPool
}
