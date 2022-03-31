package main

import (
	"sync"
	"time"
)

/*
	1. 可用 sync.RWMutex 实现同步，避免读写操作同时进行
*/

func main() {
	var lock sync.RWMutex
	m := make(map[string]int)
	go func() {
		for {
			lock.Lock()
			m["a"] += 1
			lock.Unlock()
			time.Sleep(time.Microsecond)
		}
	}()
	go func() {
		lock.RLock()
		_ = m["b"]
		lock.RUnlock()
		time.Sleep(time.Microsecond)
	}()
	select {}
}
