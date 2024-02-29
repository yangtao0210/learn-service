package od

import (
	"fmt"
)

func RichestLittleFamily() {
	var n int
	//输入成员总数
	fmt.Scan(&n)
	//输入每个节点的价值 & 初始化每个小家庭的财富和（当前节点的财富）
	wealth := make([]int, n)
	famWealth := make([]int, n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&wealth[i-1])
		famWealth[i-1] = wealth[i-1]
	}
	//存储每个小家庭的财富和
	maxWealth := 0
	//输入n-1行节点对关系 & 计算并更新最富裕和
	for i := 0; i < n-1; i++ {
		var n1, n2 int
		fmt.Scan(&n1, &n2)
		//n1的财富 + 子节点的财富
		famWealth[n1-1] += wealth[n2-1]
		//更新最大财富和
		maxWealth = max(maxWealth, famWealth[n1-1])
	}
	fmt.Println(famWealth)
	fmt.Println(maxWealth)
}
