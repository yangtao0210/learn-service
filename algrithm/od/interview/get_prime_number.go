package interview

import "math"

//题目：给定整数n，返回所有小于n的素数的数量

func GetPrimeNumber(n int) int {
	count := 0
	if n < 2 {
		return count
	}
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func isPrime(num int) bool {
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
