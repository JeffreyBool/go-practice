package main

import (
	"fmt"
)

func main() {
	findType("hello")
	findType(77)
	findType(89.88)
	findType(struct {
		name string
	}{
		name: "JeffreyBool",
	})
}

func findType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("值为字符串类型，值:%s\n", v)
	case int:
		fmt.Printf("值为整数类型，值:%d\n", v)
	default:
		fmt.Printf("未知类型: %T, %v\n", v, v)
	}
}
