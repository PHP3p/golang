package main

import (
	"errors"
	"fmt"
	"runtime"
)
//阶乘
//4*3！
//5*4！
//6*5！
func fn(n int) int {
	if n<1 {
		return 1
	}
	return n*fn(n-1)
}

func taijie(t int) int{
	if t==1{
		return t
	}else if t==2{
		return t
	}
	return fn(t-1) + fn(t-2)
}

func cpuNum(){
	fmt.Println("cpu数量",runtime.NumCPU())
	fmt.Println("cpu数量的一半",runtime.GOMAXPROCS(runtime.NumCPU()/2))
}
func main() {
	cpuNum()
	tj:=taijie(3)
	fmt.Println("上台阶问题",tj)
	jiecheng:=fn(6)
	fmt.Println("阶乘",jiecheng)
	err := panicRecover()
	if err != nil {
		fmt.Println("自定义错误",err)
		panic(err)//是否需要抛出错误并且终止执行
	}
	fmt.Println("逻辑1")
	fmt.Println("业务代码1")
}

// 异常捕获处理
func panicRecover() (err error) {
/*	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:", err)
		}
	}()*/
	a := 10
	b := 2
	//b = 0
	if b == 0 {
		return errors.New(">>>>错误友好提示，除数不可以为0")
	} else {
		c := a / b
		fmt.Println(c)
		return nil
	}
}
