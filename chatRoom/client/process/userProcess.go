package process

import (
	"awesomeProject/chatRoom/client/model"
	"awesomeProject/chatRoom/client/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (u *UserProcess) Login(userId, userPwd string) error {
	//return nil
	//创建Transmitter实例
	transmitter := utils.Transmitter{
		Conn: u.Conn,
	}
	//2.组装消息
	var loginMessage = model.LoginMes{
		UserId:  userId,
		UserPwd: userPwd,
	}
	//2.1序列化登录消息
	loginMessageBytes, err := json.Marshal(loginMessage)
	if err != nil {
		fmt.Println("loginMessage Marshal error ", err)
		return err
	}
	var mes = model.Message{
		Type: model.LoginMesType,
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
	err = transmitter.WritePkg(messageBytes)
	if err != nil {
		fmt.Println("send message error", err)
	}
	fmt.Println("client send message success,the message length is", len(messageBytes), "the content is ", string(messageBytes))
	//fmt.Println("waiting server message...")
	//创建一个缓冲器
	serverMes, err := transmitter.ReadPkg()
	if err != nil {
		return err
	}
	//反序列化服务器消息，
	var loginResultMes model.LoginResultMes
	err = json.Unmarshal([]byte(serverMes.Data), &loginResultMes)
	if err != nil {
		fmt.Println("unmarshal server message error ", err)
		return err
	}
	if loginResultMes.Code == 200 {
		//处理登录成功
		fmt.Println("login success")
		fmt.Println("online users:")
		for _, v := range loginResultMes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println("user :\t", v)
		}
		return nil
	} else {
		//fmt.Println(loginResultMes.Error)
		return errors.New(loginResultMes.Error)
	}
}

//register function
func (u *UserProcess) Register(userId, userPwd, userName string) (err error) {
	//combine user information to RegisterMes object
	var registerMes = &model.RegisterMes{
		UserId:   userId,
		UserPwd:  userPwd,
		UserName: userName,
	}
	//marshal registerMes
	registerMesBytes, err := json.Marshal(registerMes)
	if err != nil {
		return
	}
	//combine []byte registerMesBytes to Message object
	mes := &model.Message{
		Type: model.RegisterMesType,
		Data: string(registerMesBytes),
	}
	//marshal mes
	mesBytes, err := json.Marshal(mes)
	if err != nil {
		return
	}
	//send message to server
	//use write function
	//create a transmitter object
	transmitter := &utils.Transmitter{Conn: u.Conn}
	err = transmitter.WritePkg(mesBytes)
	if err != nil {
		return
	}
	//wait server deal message
	//read result
	//use readPkg function
	resultMes, err := transmitter.ReadPkg()
	if err != nil {
		return
	}
	//create a registerResultMes
	var registerResultMes model.RegisterResultMes
	//unmarshal message
	err = json.Unmarshal([]byte(resultMes.Data), &registerResultMes)
	if err != nil {
		return
	}
	if registerResultMes.Code == 200 {
		return nil
	} else {
		err = errors.New(registerResultMes.Error)
		return
	}
}
