package main

import (
	"fmt"
)

func main() {
	var a map[string]int = map[string]int{
		"stu01": 100,
		"stu02": 2000,
		"stu03": 300,
	}

	fmt.Println(a)
	a["stu01"] = 88888
	a["stu04"] = 38333
	fmt.Println(a)

	var key string = "stu04"
	fmt.Printf("the value of  key[%s] is :%d\n", key, a[key])
}
