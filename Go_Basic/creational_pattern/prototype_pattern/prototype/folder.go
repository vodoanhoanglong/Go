package prototype

import "fmt"

type Folder struct {
	Childrenses []INode
	Name        string
}

func (f *Folder) Print(s string) {
	fmt.Println(s + f.Name)
	for _, i := range f.Childrenses {
		i.Print(s + s)
	}
}

func (f *Folder) Clone() INode {
	cloneFolder := &Folder{Name: f.Name + "_Clone"}
	var tempChildrenses []INode
	for _, i := range f.Childrenses {
		copy := i.Clone()
		tempChildrenses = append(tempChildrenses, copy)
	}
	cloneFolder.Childrenses = tempChildrenses
	return cloneFolder
}
