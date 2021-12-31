package main

import "fmt"

func main() {
	checkAge := map[string]int{
		"Long": 21,
		"Loi":  22,
	}
	if age, isExist := checkAge["Loi"]; isExist {
		fmt.Println(age)
	}

	// switch case
	number := 3
	switch {
	case number <= 3:
		fmt.Println("<=3")
		fallthrough // allow continue
	case number <= 5:
		fmt.Println("<=5")
	default:
		fmt.Sprintln("Error")
	}

	var i interface{} = 1.5
	switch i.(type) {
	case int:
		fmt.Println("data type is int")
	case float64:
		fmt.Println("data type is float64")
	default:
		fmt.Println("data type is NaN")
	}
}
