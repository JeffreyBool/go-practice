package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	_, err := copyFile("target.txt", "main.go")
	if err != nil {
		fmt.Printf("copy file failed, err:%v\n", err)
		return
	}
	fmt.Println("copy done !")
}

func copyFile(dstName,srcName string) (written int64,err error)  {
	srcFd,err := os.Open(srcName)
	if err != nil {
		fmt.Printf("open source file %s failed, err: %v\n",srcName,err)
		return
	}
	defer srcFd.Close()
	dstFd,err := os.OpenFile(dstName,os.O_CREATE| os.O_WRONLY,0644)
	if err != nil {
		fmt.Printf("open dest file %s failed, err:%v\n", dstName, err)
		return
	}
	defer dstFd.Close()
	return io.Copy(dstFd,srcFd)
}
