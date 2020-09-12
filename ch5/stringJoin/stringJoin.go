package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf(stringJoin(" ", "Hi", "my", "name", "is", "enhao"))
}

func stringJoin(sep string, elems ... string) string{
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}
	n := len(sep) * (len(elems) - 1)	// space for seps
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])				// space for each elem
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(elems[0])
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}