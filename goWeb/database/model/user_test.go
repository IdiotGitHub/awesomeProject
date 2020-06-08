package model

import (
	"fmt"
	"testing"
)

func TestUser_AddUser(t *testing.T) {
	user := &User{}
	user.AddUser()
	user.AddUser2()
}
func TestUser_GetUserById(t *testing.T) {
	user := &User{Id: 2}
	var err error
	user, err = user.GetUserById()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user)
}
func TestUser_GetUsers(t *testing.T) {
	user := User{}
	users, err := user.GetUsers()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(users); i++ {
		fmt.Println(*users[i])
	}
}
