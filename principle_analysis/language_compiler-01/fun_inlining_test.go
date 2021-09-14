// Package main
/**
* @Author: wangxiaoxiong
* @Date: 2021/9/13 22:49
* @Desc：函数内联
**/
package main

import (
	"testing"
)

/*
禁止函数优化
go test fun_inlining_test.go -bench=.
*/
//go:noinline
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var Result int

func BenchmarkMax(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = max(-1, i)
	}
	Result = r
}
