package main

import (
	"fmt"
	"net"
)

//
////编写单独的函数来处理不同的消息类型--处理登录请求
//func serverLoginProcess(conn net.Conn, mes *message.Message) (err error) {
//	//反序列化mes.Data,获取客户端发送的数据，并验证登陆信息
//	var loginMes message.LoginMes
//	err = json.Unmarshal([]byte(mes.Data), &loginMes)
//	if err != nil {
//		return
//	}
//	//声明服务器返回信息
//	var resultMes message.LoginResultMes
//	//验证登录信息，目前没有链接数据库，先写死
//	if loginMes.UserId == "12345" && loginMes.UserPwd == "123" {
//		resultMes.Code = 200
//	} else {
//		resultMes.Code = 500
//		resultMes.Error = "userId or userPwd error"
//	}
//	//将验证结果返回客户端
//	//序列化验证结果,不能直接将resultMes返回给客户端，要规范起来，将resultMes封装到Message中返回
//	resultMesBytes, err := json.Marshal(resultMes)
//	if err != nil {
//		return
//	}
//	//声明新的message返回给客户端
//	var mesToClient = message.Message{
//		Type: message.LoginResultMesType,
//		Data: string(resultMesBytes),
//	}
//	mesToClientBytes, err := json.Marshal(mesToClient)
//	if err != nil {
//		return
//	}
//	err = utils.WritePkg(conn, mesToClientBytes)
//	if err != nil {
//		return
//	}
//	return
//}
////消息分类处理
//func serverProcessMes(conn net.Conn, mes *message.Message) {
//	switch mes.Type {
//	case message.LoginMesType:
//		err := serverLoginProcess(conn, mes)
//		if err != nil {
//			fmt.Println("login process error ", err)
//		}
//	case message.RegisterMesType:
//	default:
//		fmt.Println("messageType error ")
//
//	}
//}
////处理客户端连接
//func process(conn net.Conn) {
//	defer func() {
//		err := conn.Close()
//		if err != nil {
//			fmt.Println("process conn close error", err)
//			return
//		}
//	}()
//
//	for {
//		mes, err := utils.ReadPkg(conn)
//		if err != nil {
//			if err == io.EOF {
//				fmt.Println("client closed connection ")
//				return
//			}
//			//fmt.Println("read message error ", err)
//			return
//		}
//		fmt.Println("the client message is ", mes)
//		serverProcessMes(conn, &mes)
//	}
//}
func main() {
	//提示信息
	fmt.Println("监听端口8888")
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("监听端口错误:", err)
		return
	}
	//循环等待，客户端连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接受客户端连接错误")
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
		}
	}()
	pr := &Processor{Conn: conn}
	pr.Processor()
}
