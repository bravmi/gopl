// Covers all of ex3, ex4, ex5
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountTableLoop(x uint64) int {
	n := 0
	for i := 0; i < 8; i++ {
		n += int(pc[byte(x>>(i*8))])
	}
	return n
}

func PopCountShift(x uint64) int {
	n := 0
	for i := 0; i < 64; i++ {
		if x&1 != 0 {
			n++
		}
		x >>= 1
	}
	return n
}

func PopCountClear(x uint64) int {
	n := 0
	for ; x != 0; x = x & (x - 1) {
		n++
	}
	return n
}
