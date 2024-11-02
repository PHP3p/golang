package main

import (
	"flag"
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

var serverIp string
var serverPort int

func init()  {
	flag.StringVar(&serverIp,"ip","127.0.0.1","设置服务器IP地址，默认地址ip为127.0.0.1")
	flag.IntVar(&serverPort,"port",8888,"设置服务器端口8888")
}
func main()  {
	//命令解析
	flag.Parse()
	client:=NewClient(serverIp,serverPort)
	if client == nil {
		fmt.Println(">>>链接服务器失败")
		return
	}
	fmt.Println("链接服务器成功>>>")
	select {

	}
}
