package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	count int
}

func main() {
	c := Counter{
		count: 0,
	}

	var mu sync.Mutex

	for range 1000 {
		go func() {
			mu.Lock()
			c.count++ // атомарные операции
			mu.Unlock()
		}()
	}

	time.Sleep(200 * time.Millisecond)
	fmt.Println(c.count)
}
