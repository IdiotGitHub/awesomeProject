package main

import (
	"awesomeProject/chatRoom/server/model"
	"awesomeProject/chatRoom/server/process"
	"awesomeProject/chatRoom/server/utils"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//处理客户端连接
func (p *Processor) Processor() {
	transmitter := &utils.Transmitter{Conn: p.Conn}
	for {
		mes, err := transmitter.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("client closed connection ")
				return
			}
			fmt.Println("read message error ", err)
			return
		}
		fmt.Println("the client message is ", mes)
		p.ServerProcessMes(&mes)
	}
}

//消息分类处理
func (p *Processor) ServerProcessMes(mes *model.Message) {
	userProcessor := &process.UserProcessor{Conn: p.Conn}
	switch mes.Type {
	case model.LoginMesType:
		err := userProcessor.ServerLoginProcess(mes)
		if err != nil {
			fmt.Println("login process error ", err)
		}
	case model.RegisterMesType:
		err := userProcessor.Register(mes)
		if err != nil {
			fmt.Println("register error ", err)
		} else {
			fmt.Println("register success")
		}
	default:
		fmt.Println("messageType error ")

	}
}
