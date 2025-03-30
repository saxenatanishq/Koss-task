package main
import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func main() {
	myCh := make(chan string)

	wg.Add(2)
	
	go func() {
		fmt.Println(<-myCh)
		wg.Done()
	}()
	
	go func() {
		myCh <- "Yo wassup?"
		wg.Done()
	}()
	wg.Wait()
	close(myCh)
}


