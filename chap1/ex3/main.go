package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func echo1() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo2() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3() string {
	return strings.Join(os.Args[1:], " ")
}

func timeit(f func() string, n int) {
	fmt.Printf("[timing %v func]: ", GetFuncName(f))
	start := time.Now()
	for k := 0; k < n; k++ {
		_ = f()
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func GetFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func main() {
	fmt.Printf("os.Args = %v\n\n", os.Args)
	var n int = 1e7
	timeit(echo1, n)
	timeit(echo2, n)
	timeit(echo3, n)
}
