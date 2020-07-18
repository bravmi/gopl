package main

import (
	"fmt"
)

func main() {
	const (
		_ = 1 << (10 * iota)
		KiB // 1024
		MiB // 1048576
		GiB // 1073741824
		TiB // 1099511627776				(exceeds 1 << 32)
		PiB // 1125899906842624
		EiB // 1152921504606846976
		ZiB // 1180591620717411303424		(exceeds 1 << 64)
		YiB // 1208925819614629174706176

		KB = 1000
		MB = 1000 * KB // 1e6
		GB = 1000 * MB // 1e9
		TB = 1000 * GB // 1e12
		PB = 1000 * TB // 1e15
		EB = 1000 * PB // 1e18
		ZB = 1000 * EB // 1e21, overflows int (max ~9e18)
		YB = 1000 * ZB // 1e24
	)

	fmt.Println(EB)
}
