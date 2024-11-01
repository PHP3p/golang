package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	//online user list
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	//dispatch msg channel
	Msg chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{Ip: ip, Port: port, OnlineMap: make(map[string]*User), Msg: make(chan string)}
	return server

}

//监听广播消息 通知其他用户
func (s *Server) ListenActiveOnline() {
	for {
		msg := <-s.Msg
		s.mapLock.Lock()
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}
		s.mapLock.Unlock()

	}
}

//广播消息
func (s *Server) BroadCast(u *User, msg string) {
	sendMsg := fmt.Sprintf("[%s]%s:%s", u.Addr, u.Name, msg)
	s.Msg <- sendMsg
}
func (s *Server) Handler(conn net.Conn) {
	//	 do something
	fmt.Println("connect success")
	//	user Online and dispatch msg
	user := NewUser(conn, s)
	user.Online()
	isLive := make(chan bool)
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.OffOnline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err", err)
				return
			}
			msg := string(buf[:n-1])
			//s.BroadCast(user,msg)迭代第4版容易产生逻辑错误 user.dealMsg()
			user.dealMsg(msg)

			isLive <- true
		}
	}()
	//	 当前handler阻塞
	/*select {
	}*/
	for {
		select {
		case <- isLive:
		case <- time.After(time.Second * 10):
			user.SendUserMsg("time out 被踢了")
			close(user.C)
			conn.Close()
			return //或者runtime.Goexit()
		}
	}

}
func (s *Server) Start() {
	res, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer res.Close()
	go s.ListenActiveOnline()
	for {
		//accept
		conn, err := res.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//start goroutine  do handler
		go s.Handler(conn)
	}
}
