package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	chStop := make(chan bool)
	go func() {
		i := 1
		for {
			select {
			case <-ch1:
				if i%2 != 0 {
					fmt.Println(i)
				}
				i++
				ch2 <- true
			}
		}
	}()
	go func() {
		i := 1
		for {
			select {
			case <-ch2:
				if i%2 == 0 {
					fmt.Println(i)
				}
				i++
				if i > 10 {
					chStop <- true
					return
				}
				ch1 <- true
			}
		}
	}()

	// 启动时，需要给奇数协程发送一个初始信号
	ch1 <- true
	// 等待所有协程完成
	<-chStop

}
