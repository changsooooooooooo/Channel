package main

import (
	"fmt"
	"testing"
)

func TestInput(t *testing.T){
	var input string
	_, err := fmt.Scanln(&input)
	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println(input)
}
