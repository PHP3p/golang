package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	counterOnce sync.Once // 重命名为 counterOnce，用于初始化 counter
)

func initializeCounter() {
	counter = 42 // 假设初始值为 42
}

type MySingleton struct{}

var (
	instance *MySingleton
	singletonOnce sync.Once // 重命名为 singletonOnce，用于初始化单例
)

func GetInstance() *MySingleton {
	singletonOnce.Do(func() {
		instance = &MySingleton{}
	})
	return instance
}

func main() {
	// 获取单例实例
	singleton1 := GetInstance()
	singleton2 := GetInstance()

	// 输出实例的地址，确保它们是同一个实例
	fmt.Printf("Singleton1: %p\n", singleton1)
	fmt.Printf("Singleton2: %p\n", singleton2)

	var wg sync.WaitGroup

	// 启动多个 goroutine 尝试初始化 counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counterOnce.Do(initializeCounter) // 使用 counterOnce
			fmt.Println("Counter:", counter)
		}()
	}

	wg.Wait()
}