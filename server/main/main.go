package main

import (
	"chat/server/model"
	"fmt"
	"net"
	"time"
)

func process2(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process()
	if err != nil {
		fmt.Println("客户端和服务端通讯协程错误", err)
		return
	}
}
func init() {
	initPool("localhost:26379", 16, 0, 300*time.Second)
	initUserDao()
}
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", ":8889")
	if err != nil {
		fmt.Println("net.listen err=", err)
		return
	}
	defer listen.Close()
	//开始监听，等待客户端链接
	for {
		fmt.Println("等待客户端链接服务器....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.accept err=", err)
		}
		go process2(conn)
	}

}
