package main

import "fmt"

type INode interface {
	Print(string)
	Clone() INode
}

type Folder struct {
	Name     string
	Children []INode
}

func (fo *Folder) Print(indent string) {
	fmt.Println(indent, fo.Name)
	for _, ch := range fo.Children {
		ch.Print(indent + indent)
	}
}

func (fo *Folder) Clone() INode {
	newFolder := fo
	newFolder.Name = fo.Name + "_clone"

	var tempChildren []INode
	for _, i := range fo.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	newFolder.Children = tempChildren

	return newFolder
}

type File struct {
	Name string
}

func (fo *File) Print(indent string) {
	fmt.Println(indent, fo.Name)
}

func (fi *File) Clone() INode {
	newFile := fi
	newFile.Name = fi.Name + "_clone"
	return newFile
}

func main() {
	file1 := &File{Name: "File1"}
	file2 := &File{Name: "File2"}
	file3 := &File{Name: "File3"}

	folder1 := &Folder{
		Children: []INode{file1},
		Name:     "Folder1",
	}

	folder2 := &Folder{
		Children: []INode{folder1, file2, file3},
		Name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")

}
