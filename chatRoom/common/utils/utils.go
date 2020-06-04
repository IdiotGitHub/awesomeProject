package utils

import (
	"awesomeProject/chatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func ConnectServer(netProtocol, IPAddress string) (conn net.Conn, err error) {
	//1.连接服务器
	conn, err = net.Dial(netProtocol, IPAddress)
	if err != nil {
		fmt.Println("connect server error", err)
		return
	}
	return
}

//将消息发送封装成函数
func WritePkg(conn net.Conn, bytes []byte) (err error) {
	//4.1发送数据长度
	//因为conn.Write()只能接受一个切片为参数，所以需要将包长度转为一个切片
	//这里需要用到一个叫encode的包
	//声明包长度
	var pkgLen uint32
	//获取包长度
	pkgLen = uint32(len(bytes))
	//声明一个数组
	var arrLen [4]byte
	//使用binary.BigEndian.PutUint32()将pkgLen转为字节切片，并存入bytes
	binary.BigEndian.PutUint32(arrLen[0:4], pkgLen)
	//发送长度数据
	n, err := conn.Write(arrLen[0:4])
	if n != 4 || err != nil {
		fmt.Println("send data length error ", err)
		return err
	} else {
		fmt.Println("send data length success")
	}
	//发送消息数据
	_, err = conn.Write(bytes)
	if err != nil {
		fmt.Println("send message error", err)
	} else {
		fmt.Println("send message success, the message is ", string(bytes))
	}
	return
}

//将消息接收封装成函数
func ReadPkg(conn net.Conn) (mes message.Message, err error) {
	//声明一个字符缓冲切片存放消息
	var buffer = make([]byte, 8096)
	//接收客户端发送的数据
	_, err = conn.Read(buffer)

	//返回EOF error是客户端关闭了连接
	//问题，如果将两个err错误判断放在一起还是会出错，报错为<nil>，。。。。。
	//这里的错误有可能是客户端关闭连接产生的io.EOF错误，先不做处理，
	if err != nil {
		fmt.Println("receive message length error ", err)
		return
	}
	fmt.Println("received data length ", buffer[:4])
	//将消息长度转回uint32
	pkgLen := binary.BigEndian.Uint32(buffer[:4])
	//读消息
	n, err := conn.Read(buffer[:pkgLen])
	//这里的错误有可能是客户端关闭连接产生的io.EOF错误，先不做处理，
	if n != int(pkgLen) || err != nil {
		fmt.Println("read message error ", err)
		return
	}
	//反序列化消息
	err = json.Unmarshal(buffer[:pkgLen], &mes)
	if err != nil {
		fmt.Println("message unmarshal error ", err)
		return
	}
	return
}
