package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(1, 1, 2, 3, 4))
}

func max(v1 int, vals ...int) int {
	if len(vals) < 1 {
		return v1
	}
	max := v1
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}