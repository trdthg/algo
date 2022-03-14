package main

import (
	"fmt"

	"github.com/trdthg/algo/go-algo/src/util"
)

func main() {
	n, _ := util.NewReaderScanner(*util.NewReader()).ReadInt()
	a := make([]int, n<<1+2)
	// 初始化
	for i := 0; i < n; i++ {
		a[i] = 1
	}
	for i := n; i < 2*n; i++ {
		a[i] = 2
	}
	a[n<<1] = 0
	a[n<<1+1] = 0
	fmt.Println(a)
	pos := 2 * n
	mv(a, n-1, &pos)
}

// k  : 当前位置
// pos: 空礼盒位置
func move(arr []int, k int, pos *int) {
	arr[*pos] = arr[k]
	arr[k] = 0
	arr[*pos+1] = arr[k+1]
	arr[k+1] = 0
	*pos = k
	fmt.Println(arr)
}

func mv(arr []int, n int, pos *int) {
	if n == 3 {
		move(arr, 3, pos)
		move(arr, 7, pos)
		move(arr, 1, pos)
		move(arr, 6, pos)
		move(arr, 0, pos)
	} else {
		move(arr, n, pos)
		move(arr, 2*n-1, pos)
		mv(arr, n-1, pos)
	}
}
