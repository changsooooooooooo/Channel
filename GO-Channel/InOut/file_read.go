package InOut

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type ProcessFile interface{
	ReadFile(path string, fileName string) error
	FindContainLine(word string)
	makeLineList(file *os.File)
}

type ProcessFileList interface{
	makeFileList() error
	TaskNum() int
	FindTotalContainLines(fileChan chan *File) int
}

type File struct{
	lineList []string
	AnswerByFile int
}

type Files struct{
	fileList []*File
	Answer int
}

func (f *File) ReadFile(path string, fileName string) error{
	file, err := os.Open(path+fileName)
	if err!=nil{
		return err
	}
	f.makeLineList(file)
	return nil
}

func (f *File) makeLineList(file *os.File){
	f.lineList = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		f.lineList = append(f.lineList, scanner.Text())
	}
}

func (f *File) FindContainLine(word string){
	f.AnswerByFile = 0
	for _, v := range f.lineList{
		if strings.Contains(v, word){
			f.AnswerByFile += 1
		}
	}
}

func GetInputs(files *Files) error{
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil{
		return err
	}
	tempList:=strings.Split(input, " ")
	path:=tempList[0]
	err2:= files.makeFileList(path, tempList)

	if err2!=nil{
		return err2
	}
	return nil
}

func (f *Files) makeFileList(path string, fileList []string) error{
	f.fileList=make([]*File, len(fileList)-1)
	for i:=0;i<len(fileList)-1;i++{
		f.fileList[i] = &File{}
		err := f.fileList[i].ReadFile(path, fileList[i+1])
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Files) TaskNum() int{
	return len(f.fileList)
}

func (f *Files) AddTask(fileChan chan* File){
	for _, v := range f.fileList{
		fileChan<-v
	}
}

func (f *Files) FindTotalContainLines(wg *sync.WaitGroup, fileChan chan *File, word string) {
	f.Answer = 0
	for file := range fileChan{
		file.FindContainLine(word)
		f.Answer += file.AnswerByFile
	}
	wg.Done()
}
