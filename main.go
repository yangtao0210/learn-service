package main

import (
	"fmt"
	excel "github.com/xuri/excelize/v2"
	"strconv"
	"time"
)

var loc, _ = time.LoadLocation("Asia/Shanghai")

//	func main() {
//		//api := db.Apis{
//		//	ActionName:  "CreateEventNotify1",
//		//	ServiceName: "service",
//		//}
//		//if err := db.DeleteApiRecord(&api); err != nil {
//		//	log.Fatalln("删除失败")
//		//}
//		//log.Println("删除成功")
//	}

func main() {
	f, _ := excel.OpenFile("txsql_params.xlsx")
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index := 2
	for {
		key, _ := f.GetCellValue("Sheet1", fmt.Sprintf("A%d", index))
		value, _ := f.GetCellValue("Sheet1", fmt.Sprintf("B%d", index))
		index += 1
		_, err := strconv.Atoi(value)
		if err != nil {
			// 不是纯数字
			fmt.Printf("mysqld[\"%v\"] = \"%v\"\n", key, value)
		} else {
			fmt.Printf("mysqld[\"%v\"] = %v\n", key, value)
		}
		if index == 37 {
			break
		}
	}
}
