package main

import (
	"github.com/trdthg/algo/go-algo/src/util"
)

func roundRobinSchedule(a [][]int, n int) {
	if n == 0 {
		return
	}
	if n == 2 {
		a[0][0] = 1
		a[0][1] = 2
		a[1][0] = 2
		a[1][1] = 1
	} else {
		roundRobinSchedule(a, n/2)
		// 右下
		for i := n / 2; i < n; i++ {
			for j := n / 2; j < n; j++ {
				a[i][j] = a[i-n/2][j-n/2]
			}
		}
		// 右
		for i := 0; i < n/2; i++ {
			for j := n / 2; j < n; j++ {
				a[i][j] = a[i][j-n/2] + n/2
			}
		}
		// 下
		for i := n / 2; i < n; i++ {
			for j := 0; j < n/2; j++ {
				a[i][j] = a[i-n/2][j+n/2]
			}
		}
	}
}

func main() {
	n, _ := util.NewReaderScanner(*util.NewReader()).ReadIntWithMsg("请输入运动员数量: ")
	i := n
	if i != 1 || i == 0 {
		util.ExitWithMsg(1, "输入数量需要是2^n(n>=1)")
	}
	for {
		if i%2 == 0 {
			i /= 2
		} else {
			break
		}
	}
	if i != 1 {
		util.ExitWithMsg(1, "输入数量需要是2^n(n>=1)")
	}
	a := util.NewMatrix(n, n)
	roundRobinSchedule(a, n)
	util.PrintMatrix(a)
}
