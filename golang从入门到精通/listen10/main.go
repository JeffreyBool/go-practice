package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	reflectExample(x)
	reflectValue(x)
	reflectSetValue(&x)
	fmt.Printf("x value is %v\n", x)
	var b *int = new(int)
	*b = 100
	reflectSetValue(*b)
}

func reflectExample(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("type of a is:%v\n", t)
	switch t.Kind() {
	case reflect.Int64:
		fmt.Printf("a is Int64\n")
	case reflect.Float64:
		fmt.Printf("a is Float64\n")
	case reflect.Ptr:
		fmt.Printf("a is Ptr\n")
	}
}

func reflectValue(a interface{}) {
	v := reflect.ValueOf(a)
	switch v.Kind() {
	case reflect.Int64:
		fmt.Printf("a is int64, store value is:%d\n", v.Int())
	case reflect.Float64:
		fmt.Printf("a is float64, store value is:%f\n", v.Float())
	}
}

func reflectSetValue(a interface{})  {
	v := reflect.ValueOf(a)
	switch v.Kind() {
	case reflect.Int64:
		v.SetInt(100)
		fmt.Printf("a is int64, store value is:%d\n", v.Int())
	case reflect.Float64:
		v.SetFloat(6.8)
		fmt.Printf("a is float64, store value is:%f\n", v.Float())
	case reflect.Ptr:
		fmt.Printf("set a to 6.8\n")
		v.Elem().SetFloat(6.8)
	default:
		fmt.Printf("default switch\n")
	}
}
