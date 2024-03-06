package main

import (
	"fmt"
	"strings"
)

/*
운영 체제의 파일 시스템을 예시로 들어 복합체 패턴을 이해해 봅시다.
예시의 파일 시스템에는 파일과 폴더라는 두 가지 유형의 객체가 있으며
 때로는 파일과 폴더를 같은 방식으로 취급해야 하는 경우가 있는데 이럴 때 복합체 패턴이 유용합니다.

파일 시스템에서 특정 키워드에 대한 검색을 실행해야 한다고 상상해 보세요.
이 검색 작업은 파일과 폴더 모두에 적용됩니다. 파일의 경우 파일 내용만 살펴보나
폴더의 경우 검색된 키워드를 찾기 위해 해당 폴더의 모든 파일을 탐색합니다.
*/

//interface

type FileSystem interface {
	search(keyword string) []FileSystem
}

// leaf
type File struct {
	Content string
}

func (f *File) search(keyword string) []FileSystem {
	if strings.Contains(f.Content, keyword) {
		return []FileSystem{f}
	}

	return []FileSystem{}
}

// composite: 폴더 (안에 폴더, 파일 포함하고 있다)
type Folder struct {
	Children []FileSystem
	Name     string
}

func (f *Folder) search(keyword string) []FileSystem {
	result := []FileSystem{}
	for _, child := range f.Children {
		result = append(result, child.search(keyword)...)
	}

	if strings.Contains(f.Name, keyword) {
		result = append(result, f)
	}

	return result
}

func main() {

	file1 := File{Content: "hello world"}
	file2 := File{Content: "hello golang"}
	file3 := File{Content: "hello python"}

	file4 := File{Content: "hi java"}
	file5 := File{Content: "nice to see you c++"}

	folder1 := Folder{Name: "hello", Children: []FileSystem{&file1, &file2}}
	folder2 := Folder{Name: "hi", Children: []FileSystem{&file3, &file4, &file5, &folder1}}

	result1 := folder2.search("hello")
	fmt.Println(result1)

	result2 := folder2.search("hi")
	fmt.Println(result2)
}
