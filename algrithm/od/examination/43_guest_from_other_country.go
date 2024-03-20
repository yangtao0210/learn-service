package examination

import "fmt"

func GuestFromOtherCountry() {
	var k, n, m int
	fmt.Scan(&k, &n, &m)
	fmt.Println(getCountOfMBaseLuckyNumber(k, n, m))
}

func getCountOfMBaseLuckyNumber(k, n, m int) int {
	count := 0
	//判断输入是否合法
	if k < 0 || n < 0 || m <= 1 || n >= m {
		return 0
	}
	for k > 0 {
		if k%m == n {
			count++
		}
		k = k / m
	}
	return count
}
