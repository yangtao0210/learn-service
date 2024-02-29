package od

import (
	"fmt"
	"strconv"
	"strings"
)

func TestStatisticUnusedChars() {
	for {
		var str string
		fmt.Scan(&str)
		if strings.EqualFold(str, " ") {
			return
		}
		res := StatisticUnusedChars(str)
		resStr := ""
		for k, v := range res {
			resStr += string(k) + ":" + strconv.Itoa(int(v)) + ","
		}
		fmt.Println(resStr[:len(resStr)-1])
	}
}

func StatisticUnusedChars(str string) map[byte]uint8 {
	strs := strings.Split(str, "@")
	//构造全集合
	charMap := make(map[byte]uint8, 0)
	for _, peer := range strings.Split(strs[0], ",") {
		peers := strings.Split(peer, ":")
		charMap[peers[0][0]] = peers[1][0] - '0'
	}
	//构造已占用集合
	usedMap := make(map[byte]uint8, 0)
	for _, peer := range strings.Split(strs[1], ",") {
		peers := strings.Split(peer, ":")
		usedMap[peers[0][0]] = peers[1][0] - '0'
	}
	//遍历usedMap,更新全集map
	for k, v := range usedMap {
		charMap[k] -= v
		if charMap[k] == 0 {
			delete(charMap, k)
		}
	}
	return charMap
}
