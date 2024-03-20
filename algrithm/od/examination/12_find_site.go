package examination

import "fmt"

func TestFindSite() {
	var str string
	fmt.Scan(&str)
	fmt.Println(FindSite([]byte(str)))
}

func FindSite(sites []byte) int {
	count := 0
	for i := 1; i < len(sites); i++ {
		if sites[i] == '0' && (i == 0 || sites[i-1] == '0') && (i == len(sites)-1 || sites[i+1] == '0') {
			count += 1
			sites[i] = '1'
		}
	}
	fmt.Println(string(sites))
	return count
}
