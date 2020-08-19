package main

import "fmt"

// not exactly single pass but I like it
func rotateRight(a []int, k int) {
	n := len(a)
	k = k % n
	rotateLeft(a, n-k)
}

func rotateLeft(a []int, k int) {
	b := a[k:]
	b = append(b, a[:k]...)
	copy(a, b)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	rotateLeft(a, 2)
	fmt.Println("a:", a)
}
