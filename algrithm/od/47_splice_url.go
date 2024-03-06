package od

import (
	"fmt"
	"strings"
)

func SpliceUrl() {
	var str string
	fmt.Scan(&str)
	urls := strings.Split(str, ",")
	first, end := urls[0], urls[1]
	if !strings.HasSuffix(first, "/") && !strings.HasPrefix(end, "/") {
		fmt.Printf("%v/%v", first, end)
	} else if strings.HasSuffix(first, "/") && strings.HasPrefix(end, "/") {
		fmt.Printf("%v%v", first[:len(first)-1], end)
	} else {
		fmt.Printf("%v%v", first, end)
	}
}
