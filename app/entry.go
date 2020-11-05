package main

import (
	"fmt"
	"Go_practice/app/intset"
)

func main()  {
	var x intset.IntSet
	x.Add(1)
	fmt.Println(x.String())
    var words = x.Words()
    words[0] = 0
	fmt.Println(x.String())
}
