package main

func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()

}
//go run ./*.go
//go run main.go server.go
//go build -o server main.go server.go

//检测端口命令
//nc 127.0.0.1 8888

