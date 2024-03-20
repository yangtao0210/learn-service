package examination

import (
	"fmt"
)

func ComputeMaxReturns() {
	//处理输入
	var M, N, X int
	fmt.Scan(&M, &N, &X)
	returns, risks, maxInvests := make([]int, M), make([]int, M), make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&returns[i])
	}
	for i := 0; i < M; i++ {
		fmt.Scan(&risks[i])
	}
	for i := 0; i < M; i++ {
		fmt.Scan(&maxInvests[i])
	}

	maxReturn := 0
	bestInvestMents := make([]int, M)
	//计算风险范围内的最大回报投资方案
	for i := 0; i < M; i++ {
		//选取第一个投资项目：判断当前项目的风险及投资回报
		if risks[i] <= X {
			investI := min(N, maxInvests[i])
			//回报：投资额 * 回报率
			currentReturn := investI * returns[i]
			if currentReturn > maxReturn {
				maxReturn = currentReturn
				bestInvestMents = make([]int, M)
				bestInvestMents[i] = investI
			}
		}
		for j := i + 1; j < M; j++ {
			if risks[i]+risks[j] <= X {
				investI, investJ := maxInvests[i], maxInvests[j]
				//优先投资回报率高的项目：投资额
				if returns[i] > returns[j] {
					investI = min(N, investI)
					investJ = min(N-investI, investJ)
				} else {
					investJ = min(N, investJ)
					investI = min(N-investJ, investI)
				}

				//计算投资两个项目的回报
				currentReturn := investI*returns[i] + investJ*returns[j]
				if currentReturn > maxReturn {
					maxReturn = currentReturn
					bestInvestMents = make([]int, M)
					bestInvestMents[i], bestInvestMents[j] = investI, investJ
				}
			}
		}
	}
	for _, ment := range bestInvestMents {
		fmt.Printf("%v ", ment)
	}

}
