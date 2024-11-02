package main

import (
	"fmt"
	"net"
)

type Client struct {
	SeverIP string
	ServerPort int
	Name string
	conn net.Conn
}

func NewClient(serverIp string,serverPort int)*Client  {
	client:=&Client{SeverIP:serverIp,ServerPort:serverPort}
	conn,err:=net.Dial("tcp",fmt.Sprintf("%s:%d",serverIp,serverPort))
	if err != nil {
		fmt.Println("net.Dial error",err)
	}
	client.conn=conn
	return client
}
func main()  {
	client:=NewClient("127.0.0.1",8888)
	if client == nil {
		fmt.Println(">>>链接服务器失败")
		return
	}
	fmt.Println("链接服务器成功>>>")
	select {

	}
}
