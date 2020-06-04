package main

import (
	"awesomeProject/chatRoom/common/message"
	"awesomeProject/chatRoom/common/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func login(conn net.Conn, userId, userPwd, userName string) error {
	//return nil

	//2.组装消息
	var loginMessage = message.LoginMes{
		UserId:   userId,
		UserPwd:  userPwd,
		UserName: userName,
	}
	//2.1序列化登录消息
	loginMessageBytes, err := json.Marshal(loginMessage)
	if err != nil {
		fmt.Println("loginMessage Marshal error ", err)
		return err
	}
	var mes = message.Message{
		Type: message.LoginMesType,
		Data: string(loginMessageBytes),
	}
	//3.序列化消息
	messageBytes, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("message marshal error ", err)
		return err
	}
	//4.发送数据
	//4.1发送数据长度
	//因为conn.Write()只能接受一个切片为参数，所以需要将包长度转为一个切片
	//这里需要用到一个叫encode的包
	//声明包长度
	/*var pkgLen uint32
	//获取包长度
	pkgLen = uint32(len(messageBytes))
	//声明一个数组
	var bytes [4]byte
	//使用binary.BigEndian.PutUint32()将pkgLen转为字节切片，并存入bytes
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	//发送长度数据
	n, err := conn.Write(bytes[0:4])
	if n != 4 || err != nil {
		fmt.Println("send data length error ", err)
		return err
	}
	//发送消息数据
	_, err = conn.Write(messageBytes)
	if err != nil {
		fmt.Println("send message error", err)
	} else {
		fmt.Println("send message success")
	}*/
	err = utils.WritePkg(conn, messageBytes)
	if err != nil {
		fmt.Println("send message error", err)
	} else {
		fmt.Println("send message success")
	}
	fmt.Println("client send message success,the message length is", len(messageBytes), "the content is ", string(messageBytes))
	fmt.Println("waiting server message...")
	//创建一个缓冲器
	serverMes, err := utils.ReadPkg(conn)
	if err != nil {
		return err
	}
	//反序列化服务器消息，
	var loginResultMes message.LoginResultMes
	err = json.Unmarshal([]byte(serverMes.Data), &loginResultMes)
	if err != nil {
		fmt.Println("unmarshal server message error ", err)
		return err
	}
	if loginResultMes.Code == 200 {
		return nil
	} else {
		fmt.Println(loginResultMes.Error)
		return errors.New(loginResultMes.Error)
	}
}
