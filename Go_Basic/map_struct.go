package main

import (
	"fmt"
	"reflect"
)

type People struct { // People is public
	age      int // age is private, if outside package wanna to access -> Age -> public
	name     string
	isMale   bool
	subjects []string
}

type Animal struct {
	People  // extend from People struct
	species string
	phone   int `Maximum can't over 100`
}

func main() {
	// declare struct map[key]data_type {"key": value, ...}
	// constructor Map struct no value:
	peopleNameAgeMap := make(map[string]int)
	peopleNameAgeMap = map[string]int{
		"Khoa Pug":    40,
		"Vuong Pham":  20,
		"Johnny Dang": 45,
	}

	// m := map[[]int]int{} => error key of map not a slice
	// m := map[[3]int]int{} => not error

	fmt.Println(peopleNameAgeMap)
	fmt.Println(peopleNameAgeMap["Khoa Pug"]) // access data of the key name "Khoa Pug"
	peopleNameAgeMap["Khoa Pug"] = 30         // update value of the key name "Khoa Pug"
	peopleNameAgeMap["Long Vo"] = 21          // insert new value with key name "Long Vo"
	delete(peopleNameAgeMap, "Long Vo")       // delete value of the key name "Long Vo"
	_, isExist := peopleNameAgeMap["Long Vo"] // param 1: value, param 2: bool
	fmt.Println(isExist)                      // false
	fmt.Println(len(peopleNameAgeMap))

	copyMap := peopleNameAgeMap
	copyMap["Khoa Pug"] = 25 // same the memory with peopleNameAgeMap

	// struct
	studentLong := People{
		age:      18,
		name:     "Long",
		isMale:   true,
		subjects: []string{"BackEnd", "FrontEnd", "Design"},
	}

	fmt.Println(studentLong.subjects)

	peoplePlayKhoaPug := People{}

	peoplePlayKhoaPug.age = 25
	peoplePlayKhoaPug.isMale = true
	peoplePlayKhoaPug.name = "Khoa Pug"
	peoplePlayKhoaPug.subjects = []string{"Math", "English"}
	fmt.Println(peoplePlayKhoaPug)

	// quick declare
	studentNoName := struct{ name string }{name: "Happy"}
	copyStruct := studentNoName // not same the memory
	copyStruct.name = "Happy Birth Day"
	// if wanna same the memory -> copyStruct := &studentNoName
	fmt.Println(studentNoName) // Happy
	fmt.Println(copyStruct)    // Happy Birth Day

	// extend
	studentAnimal := Animal{}
	studentAnimal.name = "Dog"
	studentAnimal.isMale = true
	studentAnimal.age = 5
	studentAnimal.subjects = []string{"Cat", "Turtle", "Bird"}
	studentAnimal.species = "Wolf"
	studentAnimal.phone = 101
	fmt.Println(studentAnimal)

	// validate field
	validate := reflect.TypeOf(studentAnimal)
	field, _ := validate.FieldByName("phone") // value, error message
	fmt.Println(field)

}
