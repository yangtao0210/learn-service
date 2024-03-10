package od

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func GamesOfABC() {
	winner := ""
	gameMap := make(map[string][]string)
	reader := bufio.NewScanner(os.Stdin)
	opers := make([]string, 0)
	for reader.Scan() {
		if reader.Text() == "" {
			break
		}
		strs := strings.Split(reader.Text(), " ")
		name, oper := strs[0], strs[1]
		if _, ok := gameMap[oper]; !ok {
			gameMap[oper] = []string{name}
			opers = append(opers, oper)
		} else {
			gameMap[oper] = append(gameMap[oper], name)
		}
	}
	fmt.Println(gameMap)
	if len(gameMap) != 2 {
		fmt.Println("NULL")
	} else {
		//求解
		if slices.Contains(opers, "A") && slices.Contains(opers, "B") {
			winner = "A"
		} else if slices.Contains(opers, "B") && slices.Contains(opers, "C") {
			winner = "B"
		} else if slices.Contains(opers, "C") && slices.Contains(opers, "A") {
			winner = "C"
		} else {
			fmt.Println("NULL")
		}
	}
	//对赢家列表排序并输出结果
	if !strings.EqualFold(winner, "") {
		sort.Strings(gameMap[winner])
		for _, name := range gameMap[winner] {
			fmt.Println(name)
		}
	}

}
