package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// test1()
	test2()
}

func test1()  {
	var str string
	fmt.Scanf("%s",&str)
	fmt.Println("read from fmt:",str)
}

func test2()  {
	// 获取终端标准输出
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("read from bufio:",str)
}
