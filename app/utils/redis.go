package utils

import (
	facade2 "advt/internal/facade"
	"github.com/gomodule/redigo/redis"
	"sync"
)

type RedisInstancePool struct{}

func RedisClient() redis.Conn {
	var facade = new(facade2.Facade)
	return facade.GetRedisInstance()
}

func (r RedisInstancePool) Set(key string, value string, expire any) bool {
	if expire == false {
		expire = 86400
	}
	conn := RedisClient()
	defer conn.Close()
	_, err := conn.Do("set", key, value, "EX", expire)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (r RedisInstancePool) Get(key string) (string, error) {
	conn := RedisClient()
	defer conn.Close()
	info, err := redis.String(conn.Do("get", key))
	return info, err
}

var mu sync.Mutex

func (r RedisInstancePool) LazyGet(isFunc func(item ...interface{}) string, key string, exp any) (string, error) {
	mu.Lock()
	defer mu.Unlock()
	cacheData, _ := r.Get(key)
	if cacheData == "" {
		info := isFunc()
		r.Set(key, info, exp)
		return info, nil
	}
	return cacheData, nil
}
