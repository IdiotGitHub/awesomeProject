package main

import (
	"fmt"
	"net"
)

/*
在编写网络编程项目时，如果牵扯到发送数据读取数据时，一定要注意读取长度！！！！！在给定缓冲器时，如果不确定长度可以使用两个利用for循环循环读取，再确定长度，如果已知长度，务必确定读取长度！！！
在编写网络编程项目时，如果牵扯到发送数据读取数据时，一定要注意读取长度！！！！！在给定缓冲器时，如果不确定长度可以使用两个利用for循环循环读取，再确定长度，如果已知长度，务必确定读取长度！！！
在编写网络编程项目时，如果牵扯到发送数据读取数据时，一定要注意读取长度！！！！！在给定缓冲器时，如果不确定长度可以使用两个利用for循环循环读取，再确定长度，如果已知长度，务必确定读取长度！！！
在编写网络编程项目时，如果牵扯到发送数据读取数据时，一定要注意读取长度！！！！！在给定缓冲器时，如果不确定长度可以使用两个利用for循环循环读取，再确定长度，如果已知长度，务必确定读取长度！！！
在编写网络编程项目时，如果牵扯到发送数据读取数据时，一定要注意读取长度！！！！！在给定缓冲器时，如果不确定长度可以使用两个利用for循环循环读取，再确定长度，如果已知长度，务必确定读取长度！！！
在编写网络编程项目时，如果牵扯到发送数据读取数据时，一定要注意读取长度！！！！！在给定缓冲器时，如果不确定长度可以使用两个利用for循环循环读取，再确定长度，如果已知长度，务必确定读取长度！！！
在编写网络编程项目时，如果牵扯到发送数据读取数据时，一定要注意读取长度！！！！！在给定缓冲器时，如果不确定长度可以使用两个利用for循环循环读取，再确定长度，如果已知长度，务必确定读取长度！！！
在编写网络编程项目时，如果牵扯到发送数据读取数据时，一定要注意读取长度！！！！！在给定缓冲器时，如果不确定长度可以使用两个利用for循环循环读取，再确定长度，如果已知长度，务必确定读取长度！！！
在编写网络编程项目时，如果牵扯到发送数据读取数据时，一定要注意读取长度！！！！！在给定缓冲器时，如果不确定长度可以使用两个利用for循环循环读取，再确定长度，如果已知长度，务必确定读取长度！！！
*/
/*
使用redis时，可以使用hash结构，<对象名称	主键	json信息>
*/

func init() {
	initRedisPool(16, 0, 100, "localhost:6379")
}
func main() {
	//hint
	fmt.Println("Listen port 8888")
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Listen port 8888 error :", err)
		return
	}
	//waiting for connection
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept client error ")
			continue
		}
		go processor(conn)
	}
}

func processor(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("process conn close error", err)
			return
		} else {
			fmt.Println("server close connection with client")
		}
	}()
	pr := &Processor{Conn: conn}
	pr.Processor()
}
