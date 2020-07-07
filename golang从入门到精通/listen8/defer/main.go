package main

import (
	"fmt"
)

func testA() int  {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func testB() (x int) {
	defer func() {
		x += 1
	}()
	return 5
}

func testC() (y int) {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func testD() (x int) {
	defer func(x int) {
		x += 1
	}(x)
	return 5
}

func main() {
	fmt.Println(testA())
	fmt.Println(testB())
	fmt.Println(testC())
	fmt.Println(testD())
}
