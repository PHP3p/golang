package main

import (
	"fmt"
	"net"
	"strings"
)

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

func (u *User) SendUserMsg(msg string)  {
	u.conn.Write([]byte(msg))
}

func (u *User) dealMsg(msg string)  {

	if msg == "who" {
		u.server.mapLock.Lock()
		for _,user :=range u.server.OnlineMap{
			onlineWho :=fmt.Sprintf("[%s]%s%s\n",user.Addr,user.Name,"在线...")
			u.SendUserMsg(onlineWho)
		}
		u.server.mapLock.Unlock()
	}else if len(msg)>7&&msg[:7] == "rename|"{
		newName:=strings.Split(msg,"|")[1]
		_,ok:=u.server.OnlineMap[newName]
		if ok{
			u.SendUserMsg("当前用户被使用")
		}else{
			u.server.mapLock.Lock()
			delete(u.server.OnlineMap,u.Name)
			u.server.OnlineMap[newName]=u
			u.server.mapLock.Unlock()
			u.Name=newName
			u.SendUserMsg("您已经更新用户名:"+u.Name+"\n")
		}

	}else if len(msg)>4&&msg[:3] == "to|"{
		//消息格式 to|张三|你好
		findName:=strings.Split(msg,"|")[1]
		if  findName ==""{
			u.SendUserMsg("消息格式不正确,请使用正确格式")
			return
		}
		findUser,ok:=u.server.OnlineMap[findName]
		if !ok{
			u.SendUserMsg("当前用户不存在")
			return
		}
		content:=strings.Split(msg,"|")[2]
		if content== "" {
			u.SendUserMsg("消息内容为空，不能发送")
			return
		}
		findUser.SendUserMsg(u.Name+"对您说 "+content)
		/*else{
			u.server.mapLock.Lock()
			delete(u.server.OnlineMap,u.Name)
			u.server.OnlineMap[findName]=u
			u.server.mapLock.Unlock()
			u.Name=findName
			u.SendUserMsg("您已经更新用户名:"+u.Name+"\n")
		}
*/
	}else{
		u.server.BroadCast(u, msg)
	}
}
//listen current user
func (u *User) ListenMsg()  {
	for  {
		msg:=<-u.C
		u.conn.Write([]byte(msg+"\n"))
	}
}
