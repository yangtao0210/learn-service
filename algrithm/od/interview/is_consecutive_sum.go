package interview

func IsConsecutiveSum(num int) bool {
	if num <= 2 {
		return false
	}
	//滑动窗口解决
	start, end := 1, 2
	sum := start + end
	for end <= num/2+1 {
		if sum == num {
			return true
		} else if sum < num {
			//连续和小于num：左不变，右边界扩大（扩大区间），更新sum
			end++
			sum += end
		} else {
			//连续和大于num：右不变，左边界扩大（缩小区间）更新sum
			sum -= start
			start++
		}
	}
	return false
}
