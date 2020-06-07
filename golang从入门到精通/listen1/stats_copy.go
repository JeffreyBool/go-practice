/**
* Author: JeffreyBool
* Date: 2020/6/5
* Time: 20:08
* Software: GoLand
 */

package main

import (
	"fmt"
	"sync"
)

type Stats struct {
	sync.Mutex
	counters map[string]int
}

// bad
// Snapshot 返回当前状态
//func (s *Stats) Snapshot() map[string]int {
//	s.Lock()
//	defer s.Unlock()
//	return s.counters
//}

func (s *Stats) Snapshot() map[string]int {
	s.Lock()
	defer s.Unlock()

	result := make(map[string]int, len(s.counters))
	for k,v := range s.counters {
		result[k] = v
	}
	return result
}

func main() {
	stats := Stats{counters: make(map[string]int)}
	snapshot := stats.Snapshot()
	fmt.Println(snapshot)
}
