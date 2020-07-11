package main

import (
	"fmt"
)

func main() {
	var s interface{} = 56
	assert(s)
}

func assert(i interface{})  {
	s := i.(int)
	fmt.Println(s)
}
