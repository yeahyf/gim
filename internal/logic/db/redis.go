package db

import (
	"gim/config"
	"gim/pkg/logger"

	"github.com/go-redis/redis"
)

var RedisCli *redis.Client

func InitDB() {
	addr := config.LogicConf.RedisIP //获取redis的ip地址

	//获取redis的连接池
	RedisCli = redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	_, err := RedisCli.Ping().Result()
	if err != nil {
		logger.Sugar.Error("redis err ")
		panic(err)
	}
}
