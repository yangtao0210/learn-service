package examination

import (
	"fmt"
	"strings"
)

func LoadBalanceApiCluster() {
	var n int
	fmt.Scan(&n)
	urls := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&urls[i])
	}
	var level int
	var keyword string
	fmt.Scan(&level, &keyword)

	frequency := 0
	for _, url := range urls {
		urlLevelKeys := strings.Split(url, "/")
		if len(urlLevelKeys) > level && strings.EqualFold(urlLevelKeys[level], keyword) {
			frequency++
		}
	}
	fmt.Println(frequency)
}
