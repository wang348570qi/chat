package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"chat/common/message"
	"net"
)

func login(userId int, userPwd string) (err error) {

	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMesType
	//创建一个loginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//序列化loginMes
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.marshal err", err)
		return
	}
	//把data赋值给mes.data
	mes.Data = string(data)
	//将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	_, err = conn.Write(buf[:4])
	if err != nil {
		fmt.Println("conn.write fail", err)
		return
	}
	fmt.Println("客户端,发送消息的长度=%d 内容=%s", len(data), string(data))
	return
}
