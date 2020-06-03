package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/*
Redis
	适合做缓存，也可以持久化，也成为数据结构数据库
Redis 五大数据类型：
	1.String(字符串)、Hash、List、Set、zset
1.String
	string是redis最基本的类型，一个key对应一个value。
	string类型是二进制安全的，除普通的字符串外，也可以存放图片等数据。
	redis中的字符串value最大是512M
	CRUD
		set key value(如果key存在就是修改，如果不存在就是新增)
		get key (查询)
		del key	(删除)
		setex (set with expire)设置有效时间 -->setex keyName second value
		mset (同时设置一个或多个key-value) mset key value [key value].....
		mget (同时获取一个或多个value) mget key1 [key2...]
2.Hash--类似go map
	CURD
		1.hset keyName attrName value
		2.hget keyName attrName
		3.hmset (一次设置多个键值对)
		4.hgetall(获取所有的属性以及值)
		5.hmget (一次获取多个键值对)
		6.hlen 统计一个hash有多少个键值对
		7.hexists 查询一个hash中是否包含某字段
3.List
	CURD
		1.lpush 左推
		2.rpush	右推
		3.lrange
			lrange key start stop
			返回列表key中指定区间内的元素，区间以偏移量start和stop指定。
			下标（index）参数start和stop都是以0为底，也就是说，以0表示列表的第一个元素。
			也可以使用负数-1表示最后一个元素
		4.lpop 左取
		5.rpop 右取
		6.del	删除一个list
	如果将某一个list中的全部value都pop出去了，那么这个list就会被删除
4.Set
	set是string类型的无序集合
	底层是HashTable数据结构，Set也是存放很多字符串元素，字符串元素是无序的，而且元素的值不能重复
	CURD
		1.sadd setName	value
		2.smembers setName	取出所有的值
		3.sismember setName	value 判断是否是成员值
		4.srem	setName	value	删除指定值
常用指令
	1.添加key-value	set keyName value
	2.查看当前redis的所有key 	keys *
	3.获取key对应的值	get keyName
	4.切换redis数据库	select index(0-15)
	5.查看当前数据库的key-value数量		dbsize
	6.清空当前数据库的key-value和清空所有数据库的key-value	flushdb\flushall
*/
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: 100, //最大空闲时间，单位秒
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
func main() {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("set", "name", "大许")
	if err != nil {
		fmt.Println("写入数据错误", err)
		return
	}
	//取出
	s, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get data error", err)
		return
	}
	fmt.Println("redis get data name is ", s)
	//要想从pool取出连接，要保证pool未关闭
	pool.Close()
	conn2 := pool.Get()
	_, err = conn2.Do("set", "name2", "大许2")
	if err != nil {
		fmt.Println("写入数据错误", err)
		return
	}
	//取出
	s2, err := redis.String(conn.Do("get", "name2"))
	if err != nil {
		fmt.Println("get data error", err)
		return
	}
	fmt.Println("redis get data name is ", s2)
}
