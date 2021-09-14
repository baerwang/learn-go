// Package main
/**
* @Author: wangxiaoxiong
* @Date: 2021/9/14 6:19
* @Desc：逃逸分析
**/
package main

var z *int

// go tool compile -m escape.go
func escape() {
	a := 1
	z = &a
}

func main() {
	l := new(int)
	*l = 42

	m := &l
	n := &m
	z = **n
}
