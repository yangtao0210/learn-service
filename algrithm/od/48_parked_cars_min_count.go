package od

import (
	"fmt"
	"strings"
)

func ParkedCarsMinCount() {
	var str string
	fmt.Scan(&str)
	cars := strings.Split(str, ",")
	slots := strings.Split(strings.Join(cars, ""), "0")
	count := 0
	for _, slot := range slots {
		slotLen := len(slot)
		if slotLen != 0 {
			if slotLen%3 != 0 {
				count += 1
			}
			count += slotLen / 3
		}
	}
	fmt.Println(count)
}
