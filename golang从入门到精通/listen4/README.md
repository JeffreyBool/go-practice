## 课后练习

1. 写一个程序，统计一个字符串每个单词出现的次数。比如： s = "how do you do", 输出 how = 1 do = 2 you = 1

```go
package main

import (
	"fmt"
	"strings"
)

func statWordCount(str string) (result map[string]int) {
	words := strings.Split(str, " ")
	result = make(map[string]int, len(words))
	for _, v := range words {
		count, ok := result[v]
		if !ok {
			result[v] = 1
		} else {
			result[v] = count + 1
		}
	}
	return result
}

func main() {
	var str = "how do you do ? do you like me ?"
	result := statWordCount(str)
	fmt.Printf("result:%#v\n", result)
}
```

2. 写一个，实现学生信息的存储，学生有 ID 、年龄、分数等信息。需要非常方便的用过 ID 查找到对应的学生的信息。
