package od

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 学生成绩信息
type StuScore struct {
	Name   string
	Scores []int
}

func OrderOfStudents() {
	var n, m int
	fmt.Scan(&n, &m)

	//输入科目
	subjects := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&subjects[i])
	}

	//输入n行姓名+成绩
	scores := make([][]string, n)
	for i := 0; i < n; i++ {
		scores[i] = make([]string, m+1)
		for j := 0; j <= m; j++ {
			fmt.Scan(&scores[i][j])
		}
	}

	//输入要排名的科目
	var s string
	fmt.Scan(&s)

	//判断排名科目是否存在，并返回科目下标
	index := m
	for i, subject := range subjects {
		if strings.EqualFold(subject, s) {
			index = i
			break
		}
	}

	subjectScores := getStuScoreByScores(scores)
	//排序：(科目/总)成绩降序，若成绩相同，则姓名升序
	sort.Slice(subjectScores, func(i, j int) bool {
		if subjectScores[i].Scores[index] > subjectScores[j].Scores[index] {
			return true
		} else if subjectScores[i].Scores[index] == subjectScores[j].Scores[index] {
			return subjectScores[i].Name < subjectScores[j].Name
		}
		return false
	})

	//打印结果
	for _, ss := range subjectScores {
		fmt.Printf("%v ", ss.Name)
	}
}

func getStuScoreByScores(scores [][]string) []StuScore {
	//存储学生成绩信息
	subjectScores := make([]StuScore, 0)
	for i := 0; i < len(scores); i++ {
		total := 0
		stuscores := make([]int, 0)
		for j := 1; j < len(scores[i]); j++ {
			//记录各科成绩
			score, _ := strconv.Atoi(scores[i][j])
			stuscores = append(stuscores, score)
			total += score
		}
		//总分加入成绩的最后一列
		stuscores = append(stuscores, total)
		subjectScores = append(subjectScores, StuScore{
			Name:   scores[i][0],
			Scores: stuscores,
		})
	}
	return subjectScores
}
