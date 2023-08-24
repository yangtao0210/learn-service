package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

}

func dealFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content := scanner.Text()
		fmt.Println(content)
		lines = append(lines, strings.TrimSpace(strings.Split(content, "|")[1]))
	}

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
		return
	}

	result := strings.Join(lines, ",")
	fmt.Println(result)
}

// 两数之和
func twoSum(nums []int, target int) []int {
	hs := make(map[int]int)
	for index, num := range nums {
		if p, ok := hs[target-num]; ok {
			return []int{p, index}
		}
		hs[num] = index
	}
	return nil
}

// 二分查找
func search(nums []int, target int) int {
	if target < nums[0] || target > nums[len(nums)-1] {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
