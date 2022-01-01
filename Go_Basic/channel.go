package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	ch := make(chan int, 50) // holds up to 50 requests

	wg.Add(2)
	go func(ch <-chan int) { // receive only
		// solution 1:
		for i := range ch { // wait until ch variable there is data to continue
			fmt.Println(i)
		}
		/* solution 2:
		for{
			if i, ok := <-ch; ok{
				fmt.Println(i)
			}
			else{
				break
			}
		}
		*/
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // send only
		ch <- 42 // back to line 14
		ch <- 42 // second requests
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()

	// select
	chan1 := make(chan int, 5)
	chan2 := make(chan int, 5)
	wg.Add(2)
	go func() {
		for {
			select {
			case i := <-chan1:
				fmt.Printf("Channel 1: %v\n", i)
			case j := <-chan2:
				fmt.Printf("Channel 2: %v\n", j)
			default:
				break
			}
		}
		wg.Done()
	}()
	go func() {
		chan1 <- 42
		close(chan1)
		chan2 <- 27
		close(chan2)
		wg.Done()
	}()
	wg.Wait()
}
