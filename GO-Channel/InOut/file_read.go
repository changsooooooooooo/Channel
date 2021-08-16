package InOut

import (
	"container/list"
	"fmt"
	"io/ioutil"
)

type ProcessFile interface{
	ReadFile(path string, fileName string)
	MakeLineList()
	FindContainLine() int
}

type ProcessFileList interface{
	MakeFileList() *list.List
	FindTotalContainLines() int
}

type File struct{
	lineList *list.List
	AnswerByFile int
}

type Files struct{
	fileList *list.List
	Answer int
}

func (f *File) ReadFile(path string, fileName string){
	n, err := ioutil.ReadFile(path+fileName)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(n)
}
