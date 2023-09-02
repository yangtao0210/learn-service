package algorithm

import "fmt"

// Rescuvie 递归
func Rescuvie(n int) int {
	if n == 0 {
		return 1
	}
	return n * Rescuvie(n-1)
}

// RescuvieTial 尾递归：将递归结果作为参数传入,避免递归链越来越长
func RescuvieTial(n, r int) int {
	if n == 1 {
		return r
	}
	return RescuvieTial(n-1, r*n)
}
func F(n, a1, a2 int) int {
	if n == 0 {
		return a1
	}
	fmt.Print(a2, " ")
	return F(n-1, a2, a1+a2)
}

// BinarySearch 二分查找
func BinarySearch(array []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	midNum := array[mid]
	if midNum == target {
		return mid
	} else if midNum > target {
		return BinarySearch(array, target, left, mid-1)
	} else {
		return BinarySearch(array, target, mid+1, right)
	}
}
