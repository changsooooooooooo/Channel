package InOut

import (
	"fmt"
	"io/ioutil"
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
	path:="../Source-Text/"
	n, err := ioutil.ReadFile(path)
	fmt.Println(err)
	fmt.Println(n)
}
