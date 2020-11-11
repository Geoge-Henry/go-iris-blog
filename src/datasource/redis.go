package datasource

import (
	"errors"
	"go-web/config"
	"time"

	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

// 初始化redis
func Init() (err error) {
	redisPool = &redis.Pool{
		IdleTimeout: time.Duration(30) * time.Minute,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.Setting.RedisHost, redis.DialDatabase(config.Setting.RedisDB))
		},
	}

	conn := GetRedis()
	defer conn.Close()

	if r, _ := redis.String(conn.Do("PING")); r != "PONG" {
		err = errors.New("redis connect failed.")
	}

	return
}

// 获取redis连接
func GetRedis() redis.Conn {
	return redisPool.Get()
}

// 关闭redis
func CloseRedis() {
	if redisPool != nil {
		redisPool.Close()
	}
}
