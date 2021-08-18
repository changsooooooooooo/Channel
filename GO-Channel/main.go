package main

import (
	"GO-Channel/InOut"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fileChan:=make(chan* InOut.File)

	files:=&InOut.Files{}
	err:=InOut.GetInputs(files)
	if err!=nil{
		fmt.Println(err)
	}
	taskNum:=files.TaskNum()
	wg.Add(taskNum)

	go files.FindTotalContainLines(&wg, fileChan, "this")

	files.AddTask(fileChan)

	close(fileChan)
	wg.Wait()

	fmt.Println("Answer : ", files.Answer)
}
