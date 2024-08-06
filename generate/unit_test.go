package generate

import (
	"fmt"
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	n := 3
	var s []string
	GenerateCombinations("", n, &s)
	fmt.Println(strings.Join(s, ", "))
}
