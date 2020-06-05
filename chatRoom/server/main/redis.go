package main

import (
	"awesomeProject/chatRoom/server/model"
	"github.com/garyburd/redigo/redis"
	"time"
)

func initRedisPool(maxIdle, maxActive int, idleTimeout time.Duration, address string) {
	model.MyUserDao = &model.UserDao{
		Pool: &redis.Pool{
			MaxIdle:     maxIdle,     //最大空闲连接数
			MaxActive:   maxActive,   //表示和数据库的最大连接数，0表示没有限制
			IdleTimeout: idleTimeout, //最大空闲时间，单位秒
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", address)
			},
		},
	}
}
