package service

import (
	"errors"
	"go-web/config"
	"go-web/datasource"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	SessionRedis "github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

type RedisService struct {
}

func (self *RedisService) GetSessionDb() *SessionRedis.Database {
	db := SessionRedis.New(SessionRedis.Config{
		Network:   "tcp",
		Addr:      config.Setting.RedisHost + ":" + config.Setting.RedisPort,
		Timeout:   time.Duration(30) * time.Second,
		MaxActive: 10,
		Password:  "",
		Database:  strconv.Itoa(config.Setting.RedisDB),
		Prefix:    "",
		Delim:     "-",
		Driver:    SessionRedis.Redigo(), // redis.Radix() can be used instead.
	})
	return db
}

//设置缓存
func (self *RedisService) Set(key, val string, ttl time.Duration) error {
	conn := datasource.GetRedis()
	defer conn.Close()

	r, err := redis.String(conn.Do("SET", key, val, "EX", ttl.Seconds()))

	if err != nil {
		return err
	}

	if r != "OK" {
		return errors.New("NOT OK")
	}

	return nil
}

//获取缓存
func (self *RedisService) Get(key string) (string, error) {
	conn := datasource.GetRedis()
	defer conn.Close()

	r, err := redis.String(conn.Do("GET", key))

	if err != nil {
		return "", err
	}

	return r, nil
}
