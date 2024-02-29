package od

import (
	"fmt"
	"strings"
)

func TestValidatePassword() {
	for {
		var str string
		fmt.Scan(&str)
		if strings.EqualFold(str, "-1") {
			return
		}
		fmt.Println(ValidatePassword(str))
	}
}

func ValidatePassword(str string) (string, bool) {
	res := ""
	for i := 0; i < len(str); i++ {
		if len(res) > 0 && str[i] == '<' {
			res = res[:len(res)-1]
			continue
		}
		res += string(str[i])
	}
	return res, isValid(res)
}

func isValid(str string) bool {
	upper, lower, num, isValid := 0, 0, 0, true
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			num++
		}
		if str[i] >= 'a' && str[i] <= 'z' {
			lower++
		}
		if str[i] >= 'A' && str[i] <= 'Z' {
			upper++
		}
	}
	fmt.Println(num, upper, lower, len(str), len(str)-num-upper-lower)
	//判断密码是否合法
	if len(str) < 8 || num < 1 || lower < 1 || upper < 1 || len(str)-num-lower-upper < 1 {
		isValid = false
	}
	return isValid
}
