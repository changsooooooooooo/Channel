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
	readFile(filePath string) error
	FindContainLine(word string)
	makeLineList(file *os.File)
}

type ProcessFileList interface{
	TaskNum() int
	makeFileList(path string) error
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

func (f *File) readFile(filePath string) error{
	file, err := os.Open(filePath)
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
	fullString := path+fileName
	err2:=files.makeFileList(fullString)
	if err2!=nil{
		return err2
	}
	return nil
}

func (f *Files) makeFileList(path string) error{
	filenames, err:=filepath.Glob(path)
	if err!=nil{
		return err
	}
	f.fileList = make([]*File, len(filenames))
	for i, v := range filenames {
		f.fileList[i] = &File{}
		err2 := f.fileList[i].readFile(v)
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
