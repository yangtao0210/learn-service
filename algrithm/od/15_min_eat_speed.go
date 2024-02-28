package od

import "sort"

func MinEatingSpeed(piles []int, h int) int {
	if h < len(piles) {
		return 0
	}
	sort.Ints(piles)
	low, high := 1, piles[len(piles)-1]
	k := high
	for low < high {
		speed := low + (high-low)/2
		time := getEatTime(piles, speed)
		if time <= h {
			//花费时间小于所给时间：时间够用，减速
			k, high = speed, speed
		} else {
			//花费时间大于所给时间：时间不够用，提速
			low = speed + 1
		}
	}
	return k
}

// 计算以当前速度吃完所有桃子所花费的时间
func getEatTime(piles []int, speed int) int {
	time := 0
	for _, pile := range piles {
		time += (pile + speed - 1) / speed
	}
	return time
}
