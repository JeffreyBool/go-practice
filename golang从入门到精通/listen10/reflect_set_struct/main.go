package main

import (
	"reflect"

	"fmt"
)

type Student struct {
	Name  string
	Sex   int
	Age   int
	Score float32
}

func main() {
	var s Student
	v := reflect.ValueOf(&s)
	v.Elem().Field(0).SetString("张高元")
	v.Elem().FieldByName("Sex").SetInt(1)
	v.Elem().FieldByName("Age").SetInt(18)
	v.Elem().FieldByName("Score").SetFloat(99.2)

	fmt.Printf("s：%#v\n", s)
}
