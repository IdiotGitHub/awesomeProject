package main

import (
	"awesomeProject/chatRoom/client/process"
	"fmt"
	"net"
)

var userId string
var userPwd string
var userName string

func main() {
	//连接服务器
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("connect server is error", err)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("connection close error", err)
		} else {
			fmt.Println("connection closed")
		}
	}()
	//创建UserProcess实例
	userProcess := &process.UserProcess{Conn: conn}
	var loop bool = true
	var key int
	for loop {
		fmt.Println("----------------welcome to chat room-----------------")
		fmt.Println("\t 1 login")
		fmt.Println("\t 2 register")
		fmt.Println("\t 3 exit")
		fmt.Println("\t make your choice(1-3)")
		_, err := fmt.Scanln(&key)
		if err != nil {
			fmt.Println("wrong entity", err)
		}
		switch key {
		case 1:
			fmt.Println("enter user id：")
			_, _ = fmt.Scanln(&userId)
			fmt.Println("enter password：")
			_, _ = fmt.Scanln(&userPwd)
			err := userProcess.Login(userId, userPwd)
			if err != nil {
				fmt.Println(err)
			} else {
				process.ReadServerMes(conn)
				process.ShowMenu()
			}
			//loop = false
		case 2:
			fmt.Println("register")
			fmt.Println("enter user id：")
			_, _ = fmt.Scanln(&userId)
			fmt.Println("enter password：")
			_, _ = fmt.Scanln(&userPwd)
			fmt.Println("enter a nickName：")
			_, _ = fmt.Scanln(&userName)
			err := userProcess.Register(userId, userPwd, userName)
			if err != nil {
				fmt.Println("register failed, the problem is ", err)
			} else {
				fmt.Println("register success, please login then")
			}

		case 3:
			fmt.Println("exit")
			loop = false
		default:
			fmt.Println("wrong entity，please enter again")
		}
	}
}
