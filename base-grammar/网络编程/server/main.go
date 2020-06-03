package main

import (
	"fmt"
	"net"
	"strings"
)

/*
网络编程
	没啥好说的，都是那些套路，唯一需要注意的地方是，在读取客户端发送过来的数据的时候，必须要配合读取的长度n来进行字符串的转换！
*/
func main() {
	//1.tcp--使用的网络协议
	//2.监听端口8888
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error ", err)
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待客户端连接")
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept error ", err)
		} else {
			fmt.Println("Accept success", accept, "address is", accept.RemoteAddr().String())
		}
		go readMessage(accept)
	}
	//fmt.Printf("listen :%v\t%T\n", listen, listen)

}
func readMessage(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("客户端关闭错误", err)
		}
	}()
	for {
		var bytes []byte = make([]byte, 1024)
		n, err := conn.Read(bytes)
		if err != nil {
			fmt.Println("读取消息错误", err)
			return
		}
		//这个位置不能直接转，必须配合n来使用，否则，
		message := strings.Replace(string(bytes[:n]), " ", "", -1)
		if message == "exit" {
			return
		}
		fmt.Printf("读取了%d个消息，%v\n", n, message)
	}
}
