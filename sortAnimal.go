package main

import (
	"fmt"
)
//猫狗鱼顺序打印10次 --利用ch阻塞 不使用sync或select{}阻塞方式
func main()  {
	ch1:=make(chan bool)
	ch2:=make(chan bool)
	ch3:=make(chan bool)
	chStop:=make(chan bool)
	go func() {
		for i:=1;i<=10;i++{
			select {
			case <-ch1:
				fmt.Println("cat",i)
				ch2<-true
			}

		}
	}()
	go func() {
		for i:=1;i<=10;i++{
			select {
			case <-ch2:
				fmt.Println("dog",i)
				ch3<-true
			}
		}
	}()
	go func() {
		for i:=1;i<=10;i++{
			select {
			case <-ch3:
				fmt.Println("fish",i)
				if i >9{
					chStop<-true
					return
				}
				ch1<-true
			}
		}
	}()
	ch1<-true
	<-chStop
}