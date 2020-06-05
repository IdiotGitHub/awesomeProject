package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//set global var
var (
	MyUserDao *UserDao
)

type UserDao struct {
	Pool *redis.Pool
}
type User struct {
	UserId   string `json:"user_id"`
	UserPwd  string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

/*func (u *UserDao) NewPool(pool *redis.Pool) {
	u.Pool = pool
}*/
func (u *UserDao) GetUserById(userId string) (user User, err error) {
	var conn = u.Pool.Get()
	defer func() {
		err2 := conn.Close()
		if err2 != nil {
			fmt.Println("connection close error from pool of redis ")
		}
	}()
	userJsonString, err := redis.String(conn.Do("HGet", "users", userId))
	if err != nil {
		if err == redis.ErrNil {
			//fmt.Println("user not exists")
			err = UserNotExists
		}
		//if user info not exists, return void user struct and UserNotExists error
		return
	}
	//or create new user object
	//unmarshal userJsonString into user
	err = json.Unmarshal([]byte(userJsonString), &user)
	if err != nil {
		err = JsonMarshalError
		return
	}
	return
}

//add register logical
