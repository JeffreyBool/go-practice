/**
* Author: JeffreyBool
* Date: 2020/6/5
* Time: 20:22
* Software: GoLand
 */

package main

import (
	"fmt"
	"sync"
)

type smap struct {
	sync.Mutex // 仅适用于非导出类型
	data       map[string]string
}

func (m *smap) Get(k string) string {
	m.Lock()
	defer m.Unlock()

	return m.data[k]
}

func newSMap() *smap {
	return &smap{
		data: make(map[string]string),
	}
}

func main() {
	smap := newSMap()
	smap.data["A"] = "A"
	smap.data["B"] = "B"
	smap.data["C"] = "C"
	smap.data["D"] = "D"

	fmt.Println(smap.Get("A"))
	fmt.Println(smap.Get("B"))
	fmt.Println(smap.Get("C"))
	fmt.Println(smap.Get("D"))
}
