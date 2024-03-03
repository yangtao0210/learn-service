package od

import "fmt"

func GpuScheduleTime() {
	//1.处理输入
	var concurrent, taskLen int
	fmt.Scan(&concurrent)
	fmt.Scan(&taskLen)

	tasks := make([]int, taskLen)
	for i := 0; i < taskLen; i++ {
		fmt.Scan(&tasks[i])
	}

	//2.计算结果
	currentTasks, time, index := 0, 0, 0
	for currentTasks != 0 || index < taskLen {
		if index < taskLen {
			//当前任务数量计算
			currentTasks += tasks[index]
			index++
		}
		//运行一次GPU，任务数-concurrent & 执行时间+1
		currentTasks -= concurrent
		time++
		if currentTasks < 0 {
			currentTasks = 0
		}
	}

	//3.输出结果
	fmt.Println(time)
}
