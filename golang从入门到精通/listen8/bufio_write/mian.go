package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fd, err := os.OpenFile("./test.dat",os.O_CREATE|os.O_WRONLY,0666)
	if err != nil {
		fmt.Println("open file failed, err:",err)
		return
	}

	defer fd.Close()
	write := bufio.NewWriter(fd)
	for i := 0; i < 10; i ++ {
		write.WriteString("hello world\n")
	}
	write.Flush()
}
