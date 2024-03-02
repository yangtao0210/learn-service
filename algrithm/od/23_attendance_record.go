package od

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AttendanceRecord() {
	//先输入个数n
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
	}
	//再输入n行字符串
	for n > 0 && scanner.Scan() {
		fmt.Println(adjustGetAward(strings.Split(scanner.Text(), " ")))
		n--
	}
}

// 根据考勤记录判断是否可获奖
func adjustGetAward(records []string) bool {
	absentCount := 0
	for i := 0; i < len(records); i++ {
		//缺勤超过1次
		if strings.EqualFold(records[i], "absent") {
			absentCount++
			if absentCount > 1 {
				return false
			}
		}
		//连续迟到或者早退
		if strings.EqualFold(records[i], "leaveearly") || strings.EqualFold(records[i], "late") {
			if i > 0 && (strings.EqualFold(records[i-1], "leaveearly") || strings.EqualFold(records[i-1], "late")) {
				return false
			}
		}
		//任意连续七次考勤，非正常上班状态超过3次
		if i >= 6 {
			unPresentCount := 0
			for j := i; j >= j-6; j-- {
				if !strings.EqualFold(records[i], "present") {
					unPresentCount++
					if unPresentCount > 3 {
						return false
					}
				}
			}
			return true
		}
	}
	return true
}
