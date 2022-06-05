package main

import "fmt"

func squashSame(strs []string) []string {
	k := 0
	for _, s := range strs[1:] {
		if strs[k] != s {
			k++
			strs[k] = s
		}
	}
	return strs[:k+1]
}

func main() {
	strs := []string{"a", "a", "b", "c", "c", "c", "d", "d", "e"}
	fmt.Println("squash:")
	fmt.Println(squashSame(strs))
}
