package main

import (
	"fmt"
	"sync"
)

var count int
var wg sync.WaitGroup
var mut sync.Mutex

func increment() {
    mut.Lock()
	count++
	mut.Unlock()
	wg.Done()
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}

	wg.Wait()
	fmt.Println("Final Count: ", count)
}

