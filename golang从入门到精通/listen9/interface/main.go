package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	s := "hello world"
	describe(s)

	i := 55
	describe(i)

	strt := struct {
		name string
	}{
		name: "JeffreyBool",
	}
	describe(strt)
}
