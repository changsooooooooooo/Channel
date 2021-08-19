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
		return
	}
	taskNum:=files.TaskNum()
	wg.Add(taskNum)

	go files.FindTotalContainLines(&wg, fileChan, "this")
	//task 지정 어떻게 하면 되는지!
	files.AddTask(fileChan)

	close(fileChan)
	wg.Wait()

	fmt.Println("Answer : ", files.Answer)
}
