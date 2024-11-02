package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	SeverIP    string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{SeverIP: serverIp, ServerPort: serverPort, flag: 999}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error", err)
	}
	client.conn = conn
	return client
}
func (client *Client) SelectUser() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("select user error", err)
		return
	}
}
func (client *Client) PrivateChat() {
	var clientUser string
	var chatMsg string
	client.SelectUser()
	fmt.Println("请输入聊天对象【用户名】,exit退出")
	fmt.Scanln(&clientUser)
	for clientUser != "exit" {
		fmt.Println("请输入聊天内容，exit退出")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "to|" + clientUser + "|" + chatMsg + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("sendMsg error", err)
					break
				}

			}
			chatMsg = ""
			fmt.Println("<<<请输入聊天内容，exit退出>>>")
			fmt.Scanln(&chatMsg)
		}
		client.SelectUser()
		fmt.Println("请输入聊天对象【用户名】,exit退出")
		fmt.Scanln(&clientUser)
	}
}
func (client *Client) PublicChat() {
	//	提示用户输入信息
	var chatMsg string
	fmt.Println("请输入聊天内容，exit退出")
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit" {
		//	发送到服务端
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("sendMsg error", err)
				break
			}

		}
		chatMsg = ""
		fmt.Println("<<<请输入聊天内容，exit退出>>>")
		fmt.Scanln(&chatMsg)
	}
}
func (client *Client) UpdateUserName() bool {

	fmt.Println("请输入用户名，exit退出")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|"+client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("client rename error", err)
		return false
	}
	return true
}
var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址，默认地址ip为127.0.0.1")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口8888")
}
func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		switch client.flag {
		case 1:
			//fmt.Println("1.公聊模式\n")
			client.PublicChat()
			break
		case 2:
			//fmt.Println("2.私聊模式\n")
			client.PrivateChat()
			break
		case 3:
			//fmt.Println("3.更新用户名\n")
			client.UpdateUserName()
			break
		}

	}
}
func (client *Client) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	fmt.Scanln(&flag)
	if flag >= 0 && flag < 4 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>正确的范围<<<<")
		return false
	}
}
func (client *Client) DealResponse(){
	//一但 client.conn有值 阻塞监听 处理回执消息
	io.Copy(os.Stdout,client.conn)
	/*等同for{
		buf:=make()
		client.conn.Read(buf)
		fmt.Println("hhhh")
	}*/
}
func main() {
	//命令解析
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>链接服务器失败")
		return
	}
	go client.DealResponse()
	fmt.Println("链接服务器成功>>>")
	client.Run()
	//select {}
}
