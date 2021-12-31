package main

import "fmt"

func main() {
	i, j := 1, 0
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println(i)
			fmt.Println(j)
		}
		i++
		j++
	}

	// 2 for with label
Loop:
	for l := 1; l <= 5; l++ {
		for k := 1; k <= 5; k++ {
			if l > 1 {
				break Loop // only break -> break child for; break Loop -> break all for start Loop
			}
			fmt.Println(l * k)
		}
	}

	// loop arr
	array := []int{2, 3, 4}
	for index, value := range array {
		fmt.Println(index, value)
	}

	// loop map
	mapping := map[string]int{
		"Long": 21,
		"Loi":  22,
	}
	for key, value := range mapping {
		fmt.Println(key, value) // index -> key, value
	}

	// loop string
	str := "Hello World"
	for index, value := range str {
		fmt.Println(index, string(value))
	}

	// get index or value
	// value
	for _, value := range str {
		fmt.Println(string(value))
	}
	// index: for index := range str {fmt.Println(index)}
}
