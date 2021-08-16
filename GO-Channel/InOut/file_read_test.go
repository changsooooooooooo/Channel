package InOut

import (
	"fmt"
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
