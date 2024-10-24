package mathutil

func Abs(n int) int {
	if n < 0 {
		return n * -1.0
	}
	return n
}
