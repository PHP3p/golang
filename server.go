package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func NewServer(ip string,port int) *Server{
	server:=&Server{Ip:ip,Port:port}
	return server

}

func (s *Server) Handler(conn net.Conn) {
	//	 do something
	fmt.Println("connect success")

}
func (s *Server) Start() {
	res, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer res.Close()
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
