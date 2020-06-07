/**
* Author: JeffreyBool
* Date: 2020/6/5
* Time: 19:07
* Software: GoLand
 */

package main

import (
	"sync"
)

type F interface {
	f()
}

type S1 struct{
	sync.RWMutex
}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

// S export.
type S struct{ data string }

func (s S) Read() string      { return s.data }
func (s *S) Write(str string) { s.data = str }

func main() {
	//a := make(map[int]S)
	//a[1].Read()
	//
	//b :=  make(map[int]*S)
	//b[1].Write("B")
	//b[1].Read()

	s1Val := S1{}
	s1Val.f()
	s1Ptr := &S1{}
	s1Ptr.f()
	s2Val := S2{}
	s2Val.f()
	s2Ptr := &S2{}
	s2Ptr.f()

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Ptr
	// i = s2Val
	i.f()
}
