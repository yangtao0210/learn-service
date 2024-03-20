package interview

//题目描述
//在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。

func CanCompleteCircuit(gas []int, cost []int) int {
	//起点下标,区间汽油余量（之前区间剩余量+当前区间的gas-cost），总汽油余量
	start, curGasMargin, totalGasMargin := 0, 0, 0
	for i := range gas {
		margin := gas[i] - cost[i]
		curGasMargin += margin
		totalGasMargin += margin
		//区间汽油余量小于0，表示跑不到下一个加油站
		if curGasMargin < 0 {
			start = i + 1
			curGasMargin = 0
		}
	}
	if totalGasMargin < 0 {
		return -1
	}
	return start
}
