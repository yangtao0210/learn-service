package interview

import "fmt"

//题目：给定n，打印分母小于等于n的最简分式【分子分母公约数为1】数组

func SimplifiedFractions(n int) []string {
	if n < 2 {
		return []string{}
	}
	res := make([]string, 0)
	// 解法1：使用map存储分数值和形式
	// set := make(map[float64]string)
	// for i := 2; i <= n; i++ {
	// 	for j := 1; j < i; j++ {
	// 		key := float64(j) / float64(i)
	// 		if _, ok := set[key]; !ok {
	// 			set[key] = fmt.Sprintf("%v/%v", j, i)
	// 		}
	// 	}
	// }
	// for _,v := range set{
	//     res = append(res,v)
	// }
	// 解法2：使用数学理论：求分子，分母最大公约数为1的分数，即为最简分式
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) == 1 {
				res = append(res, fmt.Sprintf("%v/%v", j, i))
			}
		}
	}
	return res

}

// 求两个数的最大公约数 ===扩展： 最小公倍数 = 两数乘积/最大公约数
func gcd(a, b int) int {
	//递归出口:余数为0，输出另一个
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
