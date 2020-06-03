package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/*
go 使用redis
*/

func main() {
	//通过go向redis写如数据和读取数据
	//1.连接到redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis connection failed", err)
		return
	}
	defer conn.Close()
	fmt.Println("redis connection success")
	//返回的是redis操作结果(其实是一个接口)或一个错误信息
	reply, err := conn.Do("set", "name", "大许")
	if err != nil {
		fmt.Println("redis set key-value failed", err)
	}
	fmt.Println("redis set operation success", reply)
	//name, err := conn.Do("get", "name")
	//不能这样使用，它返回的是一个数组，可以直接使用包中的具体类型方法进行获取
	//name = name.(string)
	name, err := redis.String(conn.Do("get", "name"))
	fmt.Println("redis get name is", name)
}
