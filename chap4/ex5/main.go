package main

import "fmt"

func squash(strs []string) []string {
	i := 0
	for _, s := range strs[1:] {
		if strs[i] == s {
			continue
		}
		i++
		strs[i] = s
	}
	return strs[:i+1]
}

func main() {
	strs := []string{"a", "a", "b", "c", "c", "c", "d", "d", "e"}
	fmt.Println("squash:")
	fmt.Println(squash(strs))
}
