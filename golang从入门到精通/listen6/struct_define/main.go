package main

import (
	"fmt"
)

type User struct {
	Username string
	Sex      string
	Age      int32
	Avatar   string
}

func main() {
	var user User
	user.Age = 18
	user.Avatar = "http://baidu.com/image/xxx.jpg"
	user.Sex = "男"
	user.Username = "user01"

	fmt.Printf("user.username=%s age=%d sex=%s avatar=%s\n", user.Username, user.Age, user.Sex, user.Avatar)

	user2 := User{
		Username: "user02",
		Sex:      "man",
		Age:      18,
		Avatar:   "http://baidu.com/image/xxx.jpg",
	}

	fmt.Printf("user2=%#v\n", user2)

	var user3 User
	fmt.Printf("user3=%#v\n", user3)
}
