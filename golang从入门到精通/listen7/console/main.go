package main

import (
	"fmt"
	"os"
)

func main() {
	var buf [16]byte
	// 获取标准输入
	if _, err := os.Stdin.Read(buf[:]); err != nil {
		panic(err)
	}
	fmt.Println("read:", string(buf[:]))

	// 打印到标准输出
	os.Stdout.WriteString(string(buf[:]))
}
