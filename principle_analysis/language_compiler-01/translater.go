// Package principle_analysis
/**
* @Author: wangxiaoxiong
* @Date: 2021/9/13 22:19
* @Desc：词法解析
**/
package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	src := []byte("cos(+) +2i*sin(x) // Euler")

	// 初始化
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	s.Init(file, src, nil, scanner.ScanComments)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s|t%q\n", fset.Position(pos), tok, lit)
	}
}
