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

```go
package main

import (
	"fmt"
	"math/rand"
)

func testInterface() {
	var a interface{}
	var b int = 100
	var c float32 = 1.2
	var d string = "hello"

	a = b
	fmt.Printf("a=%v\n", a)

	a = c
	fmt.Printf("a=%v\n", a)

	a = d
	fmt.Printf("a=%v\n", a)
}

func studentStore() {
	var stuMap map[int]map[string]interface{}
	stuMap = make(map[int]map[string]interface{}, 16)
	//插入学生id=1，姓名=stu01, 分数=78.2, 年龄= 18
	var id = 1
	var name = "stu01"
	var score = 78.2
	var age = 18

	value, ok := stuMap[id]
	if !ok {
		value = make(map[string]interface{}, 8)
	}

	value["name"] = name
	value["id"] = id
	value["score"] = score
	value["age"] = age
	stuMap[id] = value

	fmt.Printf("stuMap:%#v\n", stuMap)

	for i := 0; i < 10; i++ {
		value, ok := stuMap[i]
		if !ok {
			value = make(map[string]interface{}, 8)
		}

		value["name"] = fmt.Sprintf("stu%d", i)
		value["id"] = i
		value["score"] = rand.Float32() * 100.0
		value["age"] = rand.Intn(100)
		stuMap[i] = value
	}

	fmt.Println()
	for k, v := range stuMap {
		fmt.Printf("id=%d stu info=%#v\n", k, v)
	}
}

func main() {
	//testInterface()
	studentStore()
}

```