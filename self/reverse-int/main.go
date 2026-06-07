package reverseint

func Reverse(in int) (res int) {
	for in != 0 {
		res *= 10
		res += in % 10
		in /= 10
	}
	return
}
