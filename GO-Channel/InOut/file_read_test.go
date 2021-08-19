package InOut

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	path:="../Source-Text/"
	fileName:="*"
	fmt.Printf("%x\n", path)
	fmt.Printf("%x\n", path+fileName)
	sb := strings.Builder{}
	sb.WriteString(path)
	sb.WriteString(fileName)
	fmt.Println(sb.String())
	fmt.Printf("%x", sb.String())
}

func TestReadFile2(t *testing.T) {
	path:="../../Source-Text/"
	n, _ := os.Open(path+"hamlet.txt")
	scanner := bufio.NewScanner(n)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
		break
	}
}

func TestGlobReadFile(t *testing.T) {
	n, _ := filepath.Glob("../../Source-Text/"+"hamlet.txt")
	fmt.Println(n)
}

func TestTotalCountCode(t *testing.T) {
	files := &Files{}
	err := files.makeFileList("../../Source-Text/hamlet.txt")
	if err != nil {
		return
	}
	for _, v := range files.fileList{
		v.findContainLine("this")
		files.Answer += v.AnswerByFile
		fmt.Println(v.AnswerByFile)
	}
}

func TestMakeFileList(t *testing.T) {
	files := &Files{}
	_ = files.makeFileList("../../Source-Text/hamlet.txt")
	fmt.Println(files.fileList[0].lineList[0])
}
