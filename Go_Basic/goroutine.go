package main

import (
	"fmt"
	"sync"
)

func count(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(name, i)
	}
}

func main() {
	var wg = sync.WaitGroup{}
	wg.Add(2)

	// go: goroutine -> parallel
	go func() {
		count("sheep")
		wg.Done() // => Wait()
	}()
	go func() {
		count("zzz")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done") // last run

	// If you want to run it sync, use sync.RWMutex{}
}
