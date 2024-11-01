package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}
//create user api
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{userAddr, userAddr, make(chan string), conn}
	go user.ListenMsg() //
	return user
}
//listen current user
func (u *User) ListenMsg()  {
	for  {
		msg:=<-u.C
		u.conn.Write([]byte(msg+"\n"))
	}
}
