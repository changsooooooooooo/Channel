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
	DoChannelTask() int
	findTotalContainLines(wg *sync.WaitGroup, fileChan chan *File, word string)
	FindTotalContainLinesWithoutChannel(word string) int
}

type File struct{
	lineList []string
	answerByFile int
}

type Files struct{
	fileList []*File
	answer int
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
	f.answerByFile = 0
	for _, v := range f.lineList{
		if strings.Contains(v, word){
			f.answerByFile += 1
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

func (f *Files) DoChannelTask(word string) int{
	var wg sync.WaitGroup
	length:=len(f.fileList)
	fileChan:=make(chan* File, length)
	//보관함이 없는 체널 안된다!
	f.answer = 0
	wg.Add(length)
	for _, v := range f.fileList{
		fileChan<-v
		go f.findTotalContainLines(&wg, fileChan, word)
	}
	wg.Wait()
	return f.answer
}

func (f *Files) findTotalContainLines(wg *sync.WaitGroup, fileChan chan *File, word string) {
	file:=<-fileChan
	file.findContainLine(word)
	f.answer += file.answerByFile
	wg.Done()
}

func (f *Files) FindTotalContainLinesWithoutChannel(word string) int {
	f.answer = 0
	for _, file := range f.fileList{
		file.findContainLine(word)
		f.answer += file.answerByFile
	}
	return f.answer
}
