package main

import (
	"builder_example/builder"
	"fmt"
)

func main() {
	normalBuilder := builder.GetBuilder("normal")
	iglooBuilder := builder.GetBuilder("igloo")

	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.GetDoorType())
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.GetWindowType())
	fmt.Printf("Normal House Floor Type: %d\n", normalHouse.GetFloorType())

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()
	fmt.Printf("Igloo House Door Type: %s\n", iglooHouse.GetDoorType())
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.GetWindowType())
	fmt.Printf("Igloo House Floor Type: %d\n", iglooHouse.GetFloorType())
}
