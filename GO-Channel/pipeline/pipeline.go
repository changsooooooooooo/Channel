package pipeline

import (
	"GO-Channel/InOut"
	"fmt"
)

type GetTarget interface{
	GetTargetWord() error
}

type GetResult interface{
	DoSequential() int
	DoAsync() int
}

type Pipe struct{
	targetWord string
	files *InOut.Files
}

func (p *Pipe) GetPath() error{
	p.files = &InOut.Files{}
	err:=InOut.GetInputs(p.files)
	if err!=nil{
		return err
	}
	return nil
}

func (p *Pipe) GetTargetWord() error{
	var input string
	_, err := fmt.Scanln(&input)
	if err!=nil{
		return err
	}
	p.targetWord = input
	return nil
}

func (p *Pipe) DoSequential() int{
	return p.files.FindTotalContainLinesWithoutChannel(p.targetWord)
}

func (p *Pipe) DoAsync() int{
	return p.files.DoChannelTask(p.targetWord)
}

