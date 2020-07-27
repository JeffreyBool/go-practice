package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/JeffreyBool/go-practice/golang从入门到精通/listen9/logger"
)

func main() {
	fmt.Println(logger.PathExists("/Users/zhanggaoyuan/go/src/github.com/JeffreyBool/go-practice/golang从入门到精通/listen9/logger/logs/test.log"))
	fmt.Println(path.Base("/user/go/all.log"))

	fmt.Println(strings.Split(path.Dir("/user/go/all.log"),".log")[0])
}
