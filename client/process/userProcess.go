package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"chat/client/utils"
	"chat/common/message"
	"net"
	"os"
)

type UserProcess struct {
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.dial err=", err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("Conn.write fail", err)
		return
	}
	fmt.Printf("客户端发送消息长度%d,内容%s", len(data), string(data))
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.write fail", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readpkg err=", err)
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")
		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.dial err=", err)
		return
	}
	defer conn.Close()
	var mes message.Message
	mes.Type = message.RegisterMesType
	var RegisterMes message.RegisterMes
	RegisterMes.User.UserId = userId
	RegisterMes.User.UserPwd = userPwd
	RegisterMes.User.UserName = userName
	data, err := json.Marshal(RegisterMes)
	if err != nil {
		fmt.Println("json,marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.marshal err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)

	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readpkg err=", err)
		return
	}
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功,重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}
