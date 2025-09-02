package main

import (
	"fmt"
	"sync"
	"time"
)

type NewMap struct {
	mp  map[string]string
	mut sync.RWMutex
}

func NewMapInit() *NewMap {
	return &NewMap{mp: make(map[string]string)}
}

func (c *NewMap) Get(key string) (string, bool) {
	c.mut.RLock()
	result, ok := c.mp[key]
	defer c.mut.RUnlock()
	return result, ok
}

func (c *NewMap) Set(key string, value string) {
	c.mut.Lock()
	c.mp[key] = value
	defer c.mut.Unlock()
}

func main() {
	m := NewMapInit()
	var wg sync.WaitGroup

	// 5 горутин пишут
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				key := fmt.Sprintf("key-%d", j)
				val := fmt.Sprintf("writer-%d", id)
				m.Set(key, val)
			}
		}(i)
	}

	// 5 горутин читают
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				key := fmt.Sprintf("key-%d", j)
				if val, ok := m.Get(key); ok {
					_ = val
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("done")
	time.Sleep(100 * time.Millisecond)
}
