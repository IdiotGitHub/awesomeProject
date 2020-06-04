package process

import (
	message "awesomeProject/chatRoom/server/model"
	"awesomeProject/chatRoom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

//使用结构体存储公共字段
type UserProcessor struct {
	Conn net.Conn
}

//编写单独的函数来处理不同的消息类型--处理登录请求
func (u *UserProcessor) ServerLoginProcess(mes *message.Message) (err error) {
	//反序列化mes.Data,获取客户端发送的数据，并验证登陆信息
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("unmarshal message error", err)
		return
	}
	//声明服务器返回信息
	var resultMes message.LoginResultMes
	//验证登录信息，目前没有链接数据库，先写死
	if loginMes.UserId == "12345" && loginMes.UserPwd == "123" {
		resultMes.Code = 200
	} else {
		resultMes.Code = 500
		resultMes.Error = "userId or userPwd error"
	}
	//将验证结果返回客户端
	//序列化验证结果,不能直接将resultMes返回给客户端，要规范起来，将resultMes封装到Message中返回
	resultMesBytes, err := json.Marshal(resultMes)
	if err != nil {
		fmt.Println("marshal message error", err)
		return
	}
	//声明新的message返回给客户端
	var mesToClient = message.Message{
		Type: message.LoginResultMesType,
		Data: string(resultMesBytes),
	}
	mesToClientBytes, err := json.Marshal(mesToClient)
	if err != nil {
		fmt.Println("marshal message error", err)
		return
	}
	//创建一个Transmitter实例
	transmitter := &utils.Transmitter{Conn: u.Conn}
	err = transmitter.WritePkg(mesToClientBytes)
	if err != nil {
		fmt.Println("send message error", err)
		return
	} else {
		fmt.Println("send message success , the message is ", string(mesToClientBytes))
	}
	return
}
