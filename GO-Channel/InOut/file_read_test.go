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
	path:="../../Source-Text/"
	n, _ := filepath.Glob(path+"*")
	fmt.Println(n)
}
