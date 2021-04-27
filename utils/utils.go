package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// 打包zip文件
func PackageZip(srcDir string, zipFileName string) {

	// 删除旧文件
	err := os.RemoveAll(zipFileName)
	if err != nil {
		return
	}

	// 创建zip文件
	create, _ := os.Create(zipFileName)

	defer create.Close()

	// 打开zip文件
	archive := zip.NewWriter(create)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == srcDir {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, srcDir+`/`)

		// 判断文件是否文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置 zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建 压缩包头部信息
		writer, _ := archive.CreateHeader(header)

		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}

		return nil
	})

}

// 解压zip文件
func UnZip(srcDir string) {

	// file read
	// 打开并读取压缩文件中的内容
	fr, err := zip.OpenReader(srcDir)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer fr.Close()

	// r.read.file 是一个集合，里面包括了压缩包里面的所有文件
	for _, file := range fr.Reader.File {
		// 判断文件该目录是否为文件夹
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(file.Name, 0644)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		// 打开文件
		r, err := file.Open()
		// 文件为空时候打印错误
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 文件名及路径
		fmt.Println("unzip:", file.Name)

		// 在对应的目录中创建相同的文件
		NewFile, err := os.Create(file.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 将内容复制
		io.Copy(NewFile, r)

		// 关闭文件
		NewFile.Close()
		r.Close()
	}
}

// 通过通道实现abc按顺序输出100次
func PrintAbc() {
	var count, forRand = 0, 100
	var ch1, ch2, ch3 = make(chan int), make(chan int), make(chan int)

	var wg sync.WaitGroup

	defer close(ch1)
	defer close(ch2)
	defer close(ch3)

	go func() {
		ch3 <- 1
	}()
	wg.Add(forRand)

	for i := 0; i < forRand; i++ {
		go func() {
			<-ch3
			fmt.Print("a,")
			ch1 <- 1
		}()

		go func() {
			<-ch1
			fmt.Print("b,")
			ch2 <- 1
		}()

		go func() {
			defer wg.Done()
			<-ch2
			fmt.Print("c,")
			count++
			fmt.Print(count, "\n")
			if count != forRand {
				ch3 <- 1
			}
		}()
	}

	wg.Wait()

}

// nil的值比较
func NilCompare() {

	// 指针类型的nil比较
	fmt.Println((*int64)(nil) == (*int64)(nil))

	// channel 类型的nil比较
	fmt.Println((chan int)(nil) == (chan int)(nil))

	// func类型的比较
	// fmt.Println((func())(nil) == (func())(nil))

	// interface类型的比较
	fmt.Println((interface{})(nil) == (interface{})(nil))

	// map类型的mil比较
	// fmt.Println((map[string]int)(nil) == (map[string]int)(nil))

	// slice类型的比较
	// fmt.Println(([]int)(nil) == ([]int)(nil))
}
