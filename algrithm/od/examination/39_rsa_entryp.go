package examination

import (
	"fmt"
	"math"
)

func RsaEncode() {
	var n int
	fmt.Scan(&n)
	first, second := -1, -1
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 && isPrime(i) && isPrime(n/i) {
			if i < n/i {
				first, second = i, n/i
			} else {
				first, second = n/i, i
			}
			break
		}
	}
	fmt.Printf("%v %v", first, second)
}

// 判断n是否为素数
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
