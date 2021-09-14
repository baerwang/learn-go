// Package main
/**
* @Author: wangxiaoxiong
* @Date: 2021/9/13 22:43
* @Desc：变量捕获
**/
package main

import "fmt"

// go tool compile -m=2 main.go | grep capturing
func main() {
	a := 1
	b := 2
	go func() {
		fmt.Println(a, b)
	}()
	a = 99
}
