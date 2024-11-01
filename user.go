package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
	server *Server
}
//create user api
func NewUser(conn net.Conn,server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{userAddr, userAddr, make(chan string), conn,server}
	go user.ListenMsg() //
	return user
}

func (u *User) Online()  {
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.mapLock.Unlock()
	//	and dispatch msg
	u.server.BroadCast(u, "已上线")
}
func (u *User) OffOnline()  {
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap,u.Name)
	u.server.mapLock.Unlock()
	//	and dispatch msg
	u.server.BroadCast(u, "下线")
}
func (u *User) dealMsg(msg string)  {
	u.server.BroadCast(u, msg)
}
//listen current user
func (u *User) ListenMsg()  {
	for  {
		msg:=<-u.C
		u.conn.Write([]byte(msg+"\n"))
	}
}
