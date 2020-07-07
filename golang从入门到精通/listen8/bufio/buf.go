package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fd, err :=  os.Open("./buf.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:",err)
			return
		}
		fmt.Println(line)
	}
}
