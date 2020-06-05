package main

import (
	"awesomeProject/chatRoom/server/model"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	userDao := &model.UserDao{Pool: &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: 100, //最大空闲时间，单位秒
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}}
	user, err := userDao.GetUserById("100")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User ", user)
}
