package InOut

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
	findTotalContainLines(fileChan chan *File) int
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

func (f *File) findContainLine(word string){
	f.AnswerByFile = 0
	for _, v := range f.lineList{
		if strings.Contains(v, word){
			f.AnswerByFile += 1
		}
	}
}

func GetInputs(files *Files) error{
	var path string
	var fileName string
	_, err := fmt.Scanf("%s %s", &path, &fileName)
	if err != nil{
		return err
	}
	err2:= files.makeFileList(path, fileName)

	if err2!=nil{
		return err2
	}
	return nil
}

func (f *Files) makeFileList(path string, fileName string) error{
	filesets, err:= filepath.Glob(path+fileName)
	if err!=nil{
		return err
	}

	f.fileList = make([]*File, len(filesets))
	for i, v := range filesets {
		f.fileList[i] = &File{}
		err2 := f.fileList[i].ReadFile(path, v)
		if err2 != nil {
			return err2
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
		file.findContainLine(word)
		f.Answer += file.AnswerByFile
	}
	wg.Done()
}
