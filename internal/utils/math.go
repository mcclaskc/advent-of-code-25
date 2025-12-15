package utils

func AbsInt(n int) int {
	if n < 0 {
		return -1 * n
	} else {
		return n
	}
}
