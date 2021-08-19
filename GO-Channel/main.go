package main

import (
	"GO-Channel/pipeline"
	"fmt"
	"time"
)

func main() {
	pipe:=&pipeline.Pipe{}
	err:=pipe.GetPath()
	if err!=nil{
		fmt.Println(err)
		return
	}
	err = pipe.GetTargetWord()
	if err != nil {
		fmt.Println(err)
		return
	}
	start:=time.Now()
	answer:=pipe.DoAsync()
	fmt.Printf("Result : %d, Elapsed time : %s\n", answer, time.Since(start))
	start=time.Now()
	answer=pipe.DoSequential()
	fmt.Printf("Result : %d, Elapsed time : %s\n", answer, time.Since(start))
}
