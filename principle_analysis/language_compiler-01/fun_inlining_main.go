// Package main
/**
* @Author: wangxiaoxiong
* @Date: 2021/9/13 22:49
* @Desc：函数内联
**/
package main

// go tool compile -m=2 fun_inlining_main.go
func small() string {
	s := "hello, " + "world"
	return s
}

func fib(index int) int {
	if index < 2 {
		return index
	}
	return fib(index-1) + fib(index-2)
}

func main() {
	small()
	fib(65)
}
