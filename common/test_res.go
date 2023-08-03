package common

import (
	"fmt"
)

type TestRes struct {
	TaskName string
	PassCnt  int
	TotalCnt int
	ErrInfos []string
}

func (tr *TestRes) GetPassCnt() int {
	return tr.PassCnt
}

func (tr *TestRes) GetTotalCnt() int {
	return tr.TotalCnt
}

func (tr *TestRes) GetPassRate() float64 {
	if tr.TotalCnt == 0 {
		return 0.0
	}
	return float64(tr.PassCnt) / float64(tr.TotalCnt)
}

func (tr *TestRes) Print() {
	fmt.Println(tr.TaskName, ":")
	fmt.Printf("PassRate: %.2f  %d / %d\n", tr.GetPassRate(), tr.PassCnt, tr.TotalCnt)
	fmt.Println("ErrInfo:")
	for _, info := range tr.ErrInfos {
		fmt.Println(info)
	}
}
