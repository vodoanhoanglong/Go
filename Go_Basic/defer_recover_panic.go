package main

import "fmt"

func main() {
	fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	// => 1, 3, 2 -> refer push value on stack and excute when end main

	// Note:
	a := 12
	defer fmt.Println(a)
	a = 24
	// => a = 12 not 24

	// defer - panic - recover
	defer func() { // first
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}() // call func
	panic("Hello")             // second
	fmt.Println("Hello world") // not excute because error

	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

func panicker() {
	fmt.Println("Create panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
			panic(err) // throw panic for main handle. Note at line 41
		}
	}()
	panic("Hello")
	// result: Start -> Create panic -> Error: Hello -> End
	// if throw for main -> Start -> Create panic -> Error: Hello
}
