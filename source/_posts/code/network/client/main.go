package main

import (
	"runtime"
	"sync"
)

func main() {
	var mu sync.Mutex
	go func() {
		mu.Lock()

		mu.Unlock()
	}()
	mu.Unlock()

	runtime.GC()
}
