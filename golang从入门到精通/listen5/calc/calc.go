/**
* Author: JeffreyBool
* Date: 2020/6/27
* Time: 17:53
* Software: GoLand
 */

package calc

import (
	"fmt"
)

var (
	Sum int
	//sub int
)

func Add(a, b int) int {
	return a + b
}

func init() {
	fmt.Println("init func in calc")
}
