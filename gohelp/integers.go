package gohelp

// ModPow function
func ModPow(x int, y int, n int) int {
	if y == 0 {
		return 1
	}
	z := ModPow(x, (y / 2), n)
	if y%2 == 0 {
		return (z * z) % n
	}
	return (x * (z * z)) % n
}
