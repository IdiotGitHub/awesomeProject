package process

import (
	"awesomeProject/chatRoom/server/model"
	"awesomeProject/chatRoom/server/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

//使用结构体存储公共字段
type UserProcessor struct {
	Conn net.Conn
}

//编写单独的函数来处理不同的消息类型--处理登录请求
func (u *UserProcessor) ServerLoginProcess(mes *model.Message) (err error) {
	//反序列化mes.Data,获取客户端发送的数据，并验证登陆信息
	var loginMes model.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		//fmt.Println("unmarshal message error", err)
		return
	}
	//声明服务器返回信息
	var resultMes model.LoginResultMes
	//验证登录信息，目前没有链接数据库，先写死
	//使用redis数据库进行验证
	user, err := model.MyUserDao.GetUserById(loginMes.UserId)
	if err != nil {
		if err == model.UserNotExists {
			resultMes.Code = 500
			resultMes.Error = err.Error()
		} else {
			resultMes.Code = 500
			resultMes.Error = err.Error()
		}

	} else if loginMes.UserPwd == user.UserPwd {
		resultMes.Code = 200
	}
	//将验证结果返回客户端
	//序列化验证结果,不能直接将resultMes返回给客户端，要规范起来，将resultMes封装到Message中返回
	resultMesBytes, err := json.Marshal(resultMes)
	if err != nil {
		//fmt.Println("marshal message error", err)
		return
	}
	//声明新的message返回给客户端
	var mesToClient = model.Message{
		Type: model.LoginResultMesType,
		Data: string(resultMesBytes),
	}
	mesToClientBytes, err := json.Marshal(mesToClient)
	if err != nil {
		//fmt.Println("marshal message error", err)
		return
	}
	//创建一个Transmitter实例
	transmitter := &utils.Transmitter{Conn: u.Conn}
	err = transmitter.WritePkg(mesToClientBytes)
	if err != nil {
		//fmt.Println("send message error", err)
		return
	} else {
		//fmt.Println("send message success , the message is ", string(mesToClientBytes))
	}
	return
}

//the deeper operation with redis in UserDao is better
//deal with register
func (u *UserProcessor) Register(mes *model.Message) (err error) {
	var registerResultMes model.RegisterResultMes
	//unmarshal client message mes
	var registerMes model.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		registerResultMes.Code = 500
		registerResultMes.Error = err.Error()
	}
	//select user information from redis
	_, err = model.MyUserDao.GetUserById(registerMes.UserId)
	if err != nil {
		if err == model.UserNotExists {
			registerResultMes.Code = 200
			_, err = model.MyUserDao.Pool.Get().Do("HSet", "users", registerMes.UserId, mes.Data)
			if err != nil {
				registerResultMes.Code = 500
				registerResultMes.Error = err.Error()
			}
		}
	} else {
		registerResultMes.Code = 500
		registerResultMes.Error = model.UserIsExists.Error()
	}
	//return result to client
	//marshal
	registerResultMesByte, err := json.Marshal(registerResultMes)
	if err != nil {
		registerResultMes.Code = 500
		registerResultMes.Error = err.Error()
	}
	var messageToClient = &model.Message{
		Type: model.RegisterResultMesType,
		Data: string(registerResultMesByte),
	}
	//marshal messageToClient
	messageToClientBytes, err := json.Marshal(messageToClient)
	if err != nil {
		fmt.Println("marshal messageToClient error ", err)
	}
	//call writePkg()
	transmitter := &utils.Transmitter{Conn: u.Conn}
	err = transmitter.WritePkg(messageToClientBytes)
	if err != nil {
		return
	}
	//recording to registerResultMes.Code judge that if user register successfully
	if registerResultMes.Code == 200 {
		return nil
	} else {
		return errors.New(registerResultMes.Error)
	}
}
