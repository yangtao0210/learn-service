package examination

import (
	"fmt"
	"sort"
)

type SeqFrequent struct {
	Seq int
	Fre int
}

func HotPageCount() {
	//先输入个数n
	var n, threshold int
	fmt.Scan(&n)
	used := make([]int, n)
	//输入n个整数
	for i := 0; i < n; i++ {
		fmt.Scan(&used[i])
	}
	fmt.Scan(&threshold)
	//统计每个内存页的访问次数
	frequentMap := make(map[int]int)
	for _, page := range used {
		if _, ok := frequentMap[page]; ok {
			frequentMap[page] += 1
		} else {
			frequentMap[page] = 1
		}
	}
	//输出结果
	if len(frequentMap) > 0 {
		//将map数据存储到结构体数组，便于排序
		seqFres := make([]SeqFrequent, 0)
		for k, v := range frequentMap {
			if v >= threshold {
				seqFres = append(seqFres, SeqFrequent{
					k, v,
				})
			}
		}
		//排序
		sort.Slice(seqFres, func(i, j int) bool {
			//频次降序
			if seqFres[i].Fre > seqFres[j].Fre {
				return true
			} else if seqFres[i].Fre == seqFres[j].Fre {
				//频次相等时,序号升序
				return seqFres[i].Seq < seqFres[j].Seq
			}
			return false
		})
		fmt.Println(len(seqFres))
		for _, seq := range seqFres {
			fmt.Printf("%v ", seq.Seq)
		}
	} else {
		fmt.Println(0)
	}
}
