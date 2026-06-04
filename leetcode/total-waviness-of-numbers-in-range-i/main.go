package main

import "fmt"

func totalWaviness(num1 int, num2 int) int {
	count := 0
	for i := num1; i <= num2; i++ {
		count += calWaviness(i)
	}
	return count
}

func calWaviness(num int) (count int) {
	if num < 100 {
		return
	}
	digits := digitsToSlice(num)
	for i := 1; i < len(digits)-1; i++ {
		prev := digits[i-1]
		cur := digits[i]
		next := digits[i+1]

		if (prev < cur && cur > next) || (prev > cur && cur < next) {
			count++
		}
	}
	return
}

func digitsToSlice(n int) (digits []int) {
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	return
}

func main() {
	fmt.Println(totalWaviness(120, 130))
	fmt.Println(totalWaviness(198, 202))
	fmt.Println(totalWaviness(4848, 4848))
}
