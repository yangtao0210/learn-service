package main

import (
	"bufio"
	"fmt"
	"learn_code/algrithm/od"
	"os"
	"strconv"
	"strings"
)

func main() {
	test_min_diff()
}

func test_min_diff() {
	read := bufio.NewReader(os.Stdin)
	//输入10个元素
	bytes, _, _ := read.ReadLine()
	nums := strings.Split(string(bytes), " ")
	scores := make([]int, 10)
	for i, num := range nums {
		number, _ := strconv.Atoi(num)
		if i < 10 {
			scores[i] = number
		}
	}
	fmt.Println(od.GetMinDiff(scores))
}

func test_min_eat() {
	read := bufio.NewReader(os.Stdin)
	//输入n个桃子数量
	bytes, _, _ := read.ReadLine()
	nums := strings.Split(string(bytes), " ")
	//输入小时h
	nBytes, _, _ := read.ReadLine()
	h, _ := strconv.Atoi(string(nBytes))
	piles := make([]int, 0)
	for _, num := range nums {
		number, _ := strconv.Atoi(num)
		piles = append(piles, number)
	}
	fmt.Println(od.MinEatingSpeed(piles, h))
}

func metaSubstrMaxLenTest() {
	var flaw int
	var str string
	fmt.Scan(&flaw)
	fmt.Scan(&str)
	fmt.Println(od.MetaSubstrMaxLength(str, flaw))
}

func test_friends_positions() {
	read := bufio.NewReader(os.Stdin)
	//输入n
	nBytes, _, _ := read.ReadLine()
	n, _ := strconv.Atoi(string(nBytes))
	//输入n个身高
	bytes, _, _ := read.ReadLine()
	nums := strings.Split(string(bytes), " ")
	heights := make([]int, n)
	for i, num := range nums {
		number, _ := strconv.Atoi(num)
		heights[i] = number
	}
	fmt.Println(od.FriendPosition(heights))
}
