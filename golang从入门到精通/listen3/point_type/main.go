/**
* Author: JeffreyBool
* Date: 2020/6/27
* Time: 02:21
* Software: GoLand
 */

package main

import (
	"fmt"
)

func TestPoint1() {
	var a int32
	a = 1000
	fmt.Printf("the addr of a :%p, a:%d\n", &a, a)

	var b *int32
	fmt.Printf("the addr of b: %p, b:%v\n", &b, b)
	if b == nil {
		fmt.Println("b is nil addr")
	}
	//*b = 100
	b = &a
	fmt.Printf("the addr of b: %p, b:%v\n", &b, b)
}

func testPoint2() {
	var a int = 200
	var b *int = &a
	fmt.Printf("b指向的地址存储的值为:%d\n", *b)
	*b = 1000
	fmt.Printf("b指向的地址储存的值为:%d\n", *b)
	fmt.Printf("a = %d\n", a)
}

func modify(a *int) {
	*a = 100
}

func testPoint3() {
	var b int = 10
	p := &b
	modify(p)
	fmt.Printf("b:%d\n", b)
}

func modify_arr(a *[3]int) {
	(*a)[0] = 100
}

func testPoint4() {
	var b [3]int = [3]int{1, 2, 3}
	modify_arr(&b)
	fmt.Printf("b:%v\n", b)
}

func testPoint5() {
	var a *int = new(int)
	*a = 100
	modify(a)
	fmt.Printf("*a=%d\n", *a)

	var b *[]int = new([]int)
	fmt.Printf("*b = %v\n", *b)
	(*b) = make([]int, 5, 100)
	(*b)[0] = 100
	(*b)[1] = 200
	fmt.Printf("*b = %v\n", *b)
}

func modifyInt(a *int) {
	*a = 100
}

func testPoint6() {
	var b = 10
	modifyInt(&b)
	fmt.Printf("b=%d\n", b)
}

func testPoint7() {
	var a int = 10
	var b *int = &a
	var c *int = b
	*c = 200
	fmt.Printf("*c=%d *b=%d a=%d\n", *c, *b, a)
}

func testPoint8() {
	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
}

func testPoint9() {
	b := 255
	a := &b
	fmt.Println("address os b is", a)
	fmt.Println("value of b is", *a)
	*a ++
	fmt.Println("new value of b is", b)
}

func main() {
	//TestPoint1()
	//testPoint2()
	//testPoint3()
	//testPoint4()
	//testPoint5()
	//testPoint6()
	//testPoint7()
	//testPoint8()
	testPoint9()
}
