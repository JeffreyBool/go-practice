/**
* Author: JeffreyBool
* Date: 2020/6/27
* Time: 02:17
* Software: GoLand
 */

package main

import (
	"fmt"
)

func main() {
	var sa = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))
	}
	fmt.Println(sa)
}
