package main

import (
	"fmt"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/tealeg/xlsx"
	"log"
	"strings"
)

func main() {
	instanceNames := []string{"实例", "instance", "实例_test_1", "_实例-1", "instance_test", "实例-", "-"}
	for _, instance := range instanceNames {
		log.Printf("'%v' is matched or not : %v", instance, regulerMatch(instance))
	}
}
func regulerMatch(instanceName string) bool {
	return gregex.IsMatchString(`[\p{Han}_a-zA-Z0-9\-]{1,60}`, instanceName)
}
func dealParamStrValue(val string) string {
	target := val
	if strings.Contains(val, "'") {
		target = strings.Split(val, "'")[1]
	}
	log.Println(target)
	return target
}

func Excel2SQL() {
	xlFile, err := xlsx.OpenFile("proxy_params.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	insertStmt := "INSERT into default_param_config(default_template_id, component, param_path,param_name, param_type, default_value,enum_value, enum_split, max_value,min_value, created_at, updated_at) values "

	// 获取第一个工作表
	sheet := xlFile.Sheets[0]
	// 遍历每一行
	for i, row := range sheet.Rows {
		// 跳过表头行
		if i == 0 {
			continue
		}
		// 生成单条插入语句的值部分
		values := "('ptpl-9m4jn9f5', 'proxy',"
		// 遍历每一列
		for index, cell := range row.Cells {
			text := cell.String()
			if index == 0 {
				values += fmt.Sprintf("'GateWay.%s', ", text)
			} else {
				values += fmt.Sprintf("'%s', ", text) // 将每个单元格的值添加到插入语句中
			}
		}
		//// 去除最后一个逗号和空格
		//values = values[:len(values)-2]
		// 添加单条插入语句的值部分到批量插入语句中
		insertStmt += values + "now(),now()), "
	}

	// 去除最后一个逗号和空格
	insertStmt = insertStmt[:len(insertStmt)-2]
	fmt.Println(insertStmt + ";")
}
