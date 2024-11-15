package main

import "fmt"

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	go func() {
		i := 1
		for {
			select {
			case <-ch1:
				fmt.Println(i)
				i++
				fmt.Println(i)
				i++
				ch2<-true
			}
		}
	}()
	go func() {
		i := 'A'
		for {
			select {
			case <-ch2:
				fmt.Println(string(i))
				i++
				fmt.Println(string(i))
				i++
				if  i>'Z'{
					ch3<-true
					return

				}
				ch1<-true
			}
		}
	}()
	ch1<-true
	<-ch3
}
