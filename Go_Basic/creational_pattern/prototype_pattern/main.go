package main

import (
	"fmt"
	p "prototype_example/prototype"
)

func main() {
	file1 := &p.File{Name: "File1"}
	file2 := &p.File{Name: "File2"}
	file3 := &p.File{Name: "File3"}
	folder1 := &p.Folder{
		Childrenses: []p.INode{file1},
		Name:        "Folder 1",
	}
	folder2 := &p.Folder{
		Childrenses: []p.INode{folder1, file2, file3},
		Name:        "Folder 2",
	}
	fmt.Println("\n Printing for Folder 2")
	folder2.Print("  ")
	cloneFolder := folder2.Clone()
	fmt.Println("\n Printing for clone folder 2")
	cloneFolder.Print("  ")
}
