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
	test_get_min_diff()
}

func test_get_min_diff() {
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

func test_prices_ss() {
	reader := bufio.NewReader(os.Stdin)
	strByte, _, _ := reader.ReadLine()
	nums := make([]int, 0)
	for _, s := range strings.Split(string(strByte), " ") {
		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}
	fmt.Println(od.PricesOfShousi(nums))
}

func test_find_site() {
	var str string
	fmt.Scan(&str)
	fmt.Println(od.FindSite([]byte(str)))
}

func test_max_contains_o() {
	var str string
	fmt.Scan(&str)
	fmt.Println(od.MaxLengthContainsO(str))
}

func test_min_diff() {
	var n int
	fmt.Scan(&n)
	scores := make([][2]int, n)
	for i, _ := range scores {
		var index, score int
		fmt.Scan(&index, &score)
		scores[i][0] = index
		scores[i][1] = score
	}
	res := od.MinDiffOfScore(scores)
	for _, r := range res {
		fmt.Printf("%v %v\n", r[0], r[1])
	}
}

func test_validate_pwd() {
	for {
		var str string
		fmt.Scan(&str)
		if strings.EqualFold(str, "-1") {
			return
		}
		fmt.Println(od.ValidatePassword(str))
	}
}

func test_unused_char() {
	for {
		var str string
		fmt.Scan(&str)
		if strings.EqualFold(str, " ") {
			return
		}
		res := od.StatisticUnusedChars(str)
		resStr := ""
		for k, v := range res {
			resStr += string(k) + ":" + strconv.Itoa(int(v)) + ","
		}
		fmt.Println(resStr[:len(resStr)-1])
	}
}

func test_number_str() {
	var s string
	var n int
	fmt.Scan(&s)
	fmt.Scan(&n)
	fmt.Println(od.NumberOfSpecialStr(s, n))
}

func test_express_nature() {
	for {
		var n int
		fmt.Scan(&n)
		if n == -1 {
			break
		}
		od.ExpressIntegerWithNature(n)
	}
}

func test_amount_count() {
	for {
		reader := bufio.NewReader(os.Stdin)
		lineBytes, _, _ := reader.ReadLine()
		strs := strings.Split(string(lineBytes), " ")
		if len(strs) <= 2 {
			break
		}
		var nums []int
		for _, str := range strs {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}
		fmt.Println(od.AmountCount(nums))
	}
}

func test_max_value(weight, value []int, bagweight int) int {
	n := len(weight)
	//dp[j]:背包容量为j时，可装物品的最大价值和

	dp := make([]int, bagweight+1)
	for i := 0; i < n; i++ {
		for j := weight[i]; j <= bagweight; j++ {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[bagweight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
