package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	name string

	age int
}

//go的协程执行顺序不好说

//name和age不一定相等，因为有sleep

//加锁可以解决

var p Person
var lock sync.Mutex // 新增互斥锁
func update(name string, age int) {
	lock.Lock() // 加锁
	p.name = name

	time.Sleep(time.Millisecond * 200)
	p.age = age
	lock.Unlock() // 解锁
}

func main() {

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {

		name, age := fmt.Sprintf("nobody:%v", i), i

		go func() {

			defer wg.Done()

			update(name, age)

		}()

	}

	wg.Wait()
	fmt.Printf("p.name= %s\np.age=%v", p.name, p.age)
}

