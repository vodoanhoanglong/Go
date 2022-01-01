package main

import "fmt"

func generatePipeline(numbers []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func fanIn(inputChannel ...<-chan int) <-chan int {
	in := make(chan int)
	go func() {
		for _, c := range inputChannel {
			for n := range c {
				in <- n
			}
		}
	}()
	return in
}

func main() {
	randomNumbers := []int{}
	for i := 1; i < 1000; i++ {
		randomNumbers = append(randomNumbers, i)
	}
	// generate pipeline
	inputChan := generatePipeline(randomNumbers)

	// fan-out
	// use more cpu but faster
	cpu1 := fanOut(inputChan)
	cpu2 := fanOut(inputChan)
	cpu3 := fanOut(inputChan)
	cpu4 := fanOut(inputChan)

	// fan-in
	cpu := fanIn(cpu1, cpu2, cpu3, cpu4)
	sum := 0
	for i := 0; i < len(randomNumbers); i++ {
		sum += <-cpu
	}
	fmt.Printf("Total sum of square: %d", sum)
}
