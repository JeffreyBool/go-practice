package main

import (
	"fmt"

	_ "github.com/JeffreyBool/go-practice/golang从入门到精通/listen5/calc"
)

var a int = 1000
var b int = 2000

func init() {
	fmt.Println(a, b)
	fmt.Println("init func in main")
}

func init() {
	fmt.Println("init2 func in main")
}

func main() {
	sum = calc.Add(2, 3)
	fmt.Printf("sum=%d\n", sum)
}
