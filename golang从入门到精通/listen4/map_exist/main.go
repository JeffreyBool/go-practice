package main

import (
	"fmt"
)

func main() {
	var a map[string]int

	a = make(map[string]int, 16)
	a["stu01"] = 1000
	a["stu02"] = 1000
	a["stu03"] = 1000
	fmt.Printf("a=%#v\n", a)

	var (
		result int
		ok     bool
	)
	result, ok = a["stu03"]
	if !ok {
		fmt.Printf("key %s is not exist\n", "stu03")
	} else {
		fmt.Printf("key %s is %d\n", "stu03", result)
	}
}
