package main

import (
	"testing"
)

func TestPanicRecover(t *testing.T){

	err := panicRecover()
	if err != nil {
		t.Fatal("捕获到错误，执行成功")
	}
	t.Logf("没有异常发生")
}
