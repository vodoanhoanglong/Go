package main

import "fmt"

func findMax(a, b *int) int {
	*a = 100
	if *a > *b {
		return *a
	}
	return *b
}

func findMin(nu1, nu2 int) *int {
	nu1 = 1
	if nu1 < nu2 {
		return &nu1
	}
	return &nu2
}

func findMax2(a, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return
}

func findSum(a ...int) (sum int) { // param is slice or arr
	// Note: (variable ...data_type) must be in the last parameter
	for _, num := range a {
		sum += num
	}
	return
}

func callDivide(a, b int) (int, error) { // Go allow value multiple return
	result := 0
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	result = a / b
	return result, nil
}

type greeter struct {
	greeting string
	name     string
}

// func (object) func_name(parameter) data_type_return
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "error" // different memory address -> not change
}

func (g *greeter) changeName() {
	g.name = "Yen"
}

func main() {
	a := 10
	b := 20
	fmt.Println(findMax(&a, &b))
	fmt.Println(a)

	min := findMin(a, b) // 2 memory is left and right -> not performance
	// solution: set return data type of findMin is pointer. EX: line 13
	fmt.Println(*min)

	max := findMax2(a, b)
	fmt.Println(max)

	sum := findSum(1, 2, 3, 4, 5, 6)
	fmt.Println(sum)

	res, err := callDivide(4, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	// anonymous function
	func() {
		fmt.Println("Hello")
	}() // call this func

	for i := 0; i <= 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	// struct func
	g := greeter{
		greeting: "Hello Long",
		name:     "Go",
	}
	g.greet()           // method of Greeter struct
	fmt.Println(g.name) // Go
	g.changeName()
	fmt.Println(g.name) // Yen
}
