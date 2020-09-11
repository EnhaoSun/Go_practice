// 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a := returnN()
	fmt.Println(a)
}

func returnN() (result int) {
	defer func() {
		if p := recover(); p != nil {
			switch t := p.(type) {
			case int:
				fmt.Println("p store int", t)
				result = t
			case string:
				fmt.Println("p store string", t)
				var err error
				result, err = strconv.Atoi(t)
				if err != nil {
					_, _ =fmt.Fprintf(os.Stderr, "Can not parse string: %s to integer\n", t)
				}
			}
		}
	}()
	panic("3aaa")
}

