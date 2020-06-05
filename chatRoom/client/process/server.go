package process

import (
	"awesomeProject/chatRoom/client/utils"
	"fmt"
	"net"
	"os"
)

/*
和服务器沟通
*/
func ShowMenu() {
	for {
		fmt.Println("------------login successfully-----------")
		fmt.Println("1.show users who online")
		fmt.Println("2.show message list")
		fmt.Println("3.exit")
		fmt.Println("make your choice(1-4)")
		var key int
		_, err := fmt.Scanf("%d\n", &key)
		if err != nil {
			fmt.Println("read input error ", err)
			continue
		}
		switch key {
		case 1:
			fmt.Println("show users who are online...")
		case 2:
			fmt.Println("show messages...")
		case 3:
			fmt.Println("System exit...")
			os.Exit(0)
		}

	}
}
func ReadServerMes(conn net.Conn) {
	//创建Transmitter实例
	transmitter := &utils.Transmitter{Conn: conn}
	for {
		mes, err := transmitter.ReadPkg()
		if err != nil {
			fmt.Println("read server message error ", err)
			continue
		}
		fmt.Println("receive message from server ", mes.Data)

	}
}
