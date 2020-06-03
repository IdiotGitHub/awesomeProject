package main

import "fmt"

var userId string
var userPwd string

func main() {
	var loop bool = true
	var key int
	for loop {
		fmt.Println("----------------欢迎登陆多人聊天系统-----------------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出聊天室")
		fmt.Println("\t\t\t 请选择(1-3)")
		_, err := fmt.Scanln(&key)
		if err != nil {
			fmt.Println("输入错误", err)
		}
		switch key {
		case 1:
			fmt.Println("请输入用户ID：")
			fmt.Scanln(&userId)
			fmt.Println("请输入密码：")
			fmt.Scanln(&userPwd)
			err := login(userId, userPwd)
			if err != nil {
				fmt.Println("登陆失败")
			} else {
				fmt.Println("登陆成功")
			}
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出聊天室")
			loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
