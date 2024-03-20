package examination

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func TestGetMinRobotEnergy() {
	read := bufio.NewReader(os.Stdin)
	//输入砖块列表
	bytes, _, _ := read.ReadLine()
	nums := strings.Split(string(bytes), " ")
	bricks := make([]int, 0)
	for _, num := range nums {
		number, _ := strconv.Atoi(num)
		bricks = append(bricks, number)
	}
	fmt.Println(GetMinRobotEnergy(bricks))
}

func GetMinRobotEnergy(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	l, r, energy := 1, nums[n-1], nums[n-1]
	if n >= 8 {
		if n == 8 {
			return nums[n-1]
		}
		return -1
	}
	//二分法求解
	for l < r {
		tempEnery := l + (r-l)/2
		cost := getTime(nums, tempEnery)
		if cost <= 8 {
			//花费小于8：每次充能太大，适当减小
			r, energy = tempEnery, tempEnery

		} else {
			//花费大于8：充能太小，无法完成工作，适当加大
			l = energy + 1
		}
	}
	return energy
}

// 计算能量为energy时搬完需要的时间
func getTime(nums []int, energy int) int {
	time := 0
	for _, num := range nums {
		time += int(math.Ceil(float64(num / energy)))
	}
	return time
}
