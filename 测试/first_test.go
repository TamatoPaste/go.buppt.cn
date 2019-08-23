package main

import (
	"fmt"
	"testing"
)

type cb func(int) int

func TestFirstTry(t *testing.T) {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})
	fmt.Print("主程序跑完啦！")
}

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}
