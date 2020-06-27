package main

import (
	"fmt"
)

func main() {
	var a map[string]int
	fmt.Printf("a:%v\n",a)

	a = make(map[string]int,16)
	a["stu01"] = 1000
	a["stu02"] = 2000
	a["stu03"] = 3000
	fmt.Printf("a=%#v\n",a)
}
