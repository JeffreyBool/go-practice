/**
* Author: JeffreyBool
* Date: 2020/6/27
* Time: 01:50
* Software: GoLand
 */

package main

import (
	"fmt"
)

func main() {
	var a []int
	a = make([]int, 5, 10)
	a[0] = 10
	//a[1] = 20
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	a = append(a, 11)
	fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))

	for i := 0; i < 8; i++ {
		a = append(a, i)
		fmt.Printf("a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
	}
	//观察切片的扩容操作，扩容的策略是翻倍扩容
	a = append(a, 1000)
	fmt.Printf("扩容之后的地址：a=%v addr:%p len:%d cap:%d\n", a, a, len(a), cap(a))
}

