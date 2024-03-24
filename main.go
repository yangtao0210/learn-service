package main

import (
	"fmt"
	"reflect"
)

func main() {
	//od.TestSubstrLength()
	//od.TestBalancedStringSplit()
	//od.TestGetMinRobotEnergy()
	//od.TaxiBilling()
	//od.RichestLittleFamily()
	//od.OpenResourceHotSort()
	//od.AttendanceRecord()
	//od.DiffAbsOfHeight()
	//od.CoverMaxArea()
	//od.MaxWeightRemainSilver()
	//od.HotPageCount()
	//od.AbstractOfString()
	//od.SortedByLowestPosition()
	//od.ClearRepeatElem()
	//od.OrderOfStudents()
	//od.HeightAndWeightOrder()
	//od.ExchangeToMinStr()
	//od.GpuScheduleTime()
	//od.SumOfMaxAndMinN()
	//od.SearchPositionBinary()
	//od.MaxCountGems()
	//od.RsaEncode()
	//od.WordCountOfGrasp()
	//od.AssignCpuFromTwoGroups()
	//od.GetLuckyNumber()
	//od.GuestFromOtherCountry()
	//od.PathsOfVisited()
	//od.PrintWordMatchedWithInputPrefix()
	//od.PositionSmallestKAscii()
	//od.SpliceUrl()
	//od.ParkedCarsMinCount()
	//od.LoadBalanceApiCluster()
	//od.MaxDispatchTeamsCount()
	//od.SumOfContinuousInterval()
	//od.SplitAndTransferString()
	//od.ContinuousCharLength()
	//od.TransferAndCompulateExpress()
	//od.ComputeMaxReturns()
	//od.DrawingMachine()
	//od.AirportFlightScheduler()
	//od.Airs()
	//od.SearchGold()
	//od.HeightOfThreeTree()
	//examination.GamesOfABC()
	//interview.ArrangementString()
	//fmt.Println(interview.Top1Nums([]int{1, 2, 3, 2, 4, 2, 4, 4}))
	//fmt.Println(interview.IsConsecutiveSum(8))
	//fmt.Println(interview.AssignCandies([]int{2, 0, 1, 3, 4, 0, 2}))
	//fmt.Println(interview.GetPrimeNumber(4))
	//fmt.Println(interview.MinTime())
	//fmt.Println(interview.GetMaxSumOfSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(float64(1/2), reflect.TypeOf(float64(1/2)))
}

func simplifiedFractions(n int) []string {
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
