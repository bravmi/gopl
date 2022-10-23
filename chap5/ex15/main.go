package main

import "fmt"

func min(res int, vals ...int) int {
	for _, val := range vals {
		if val < res {
			res = val
		}
	}
	return res
}

func max(res int, vals ...int) int {
	for _, val := range vals {
		if val > res {
			res = val
		}
	}
	return res
}

func main() {
	fmt.Println(min(1, 2, 3, 4))
	fmt.Println(max(1, 2, 3, 4))
}
