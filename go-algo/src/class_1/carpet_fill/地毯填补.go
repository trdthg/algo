package main

import (
	"fmt"
	"math"

	"github.com/trdthg/algo/go-algo/src/util"
)

func main() {
	sc := util.NewReaderScanner(*util.NewReader())
	n, _ := sc.ReadIntWithMsg("请输入迷宫大小k(0< k <=10): ")
	n = int(math.Pow(2, float64(n)))
	x, _ := sc.ReadIntWithMsg("请输入公主所在位置x, y: ")
	y, _ := sc.ReadInt()
	a := util.NewMatrix(n, n)
	fmt.Println(x, y, n)
	solve(a, x-1, y-1, 0, 0, n)
}

// m(matrix) => 矩阵
// x,y => 公主坐标
// a,b => 矩阵左上坐标
// l(len) => 矩阵宽度
func solve(m [][]int, x, y int, a, b int, l int) {
	if l == 1 {
		return
	}
	// 4个左上角
	x1, y1 := a, b
	x2, y2 := a, b+l/2
	x3, y3 := a+l/2, b
	x4, y4 := a+l/2, b+l/2
	// 中间4点
	xlt, ylt := a+l/2-1, b+l/2-1
	xrt, yrt := a+l/2-1, b+l/2
	xlb, ylb := a+l/2, b+l/2-1
	xrb, yrb := a+l/2, b+l/2
	// 左上
	if x-a < l/2 && y-b < l/2 {
		// 每次打印的都是中间的3个构成的
		fmt.Printf("%d %d 1\n", xrb+1, yrb+1)
		// 左上(自己)
		// fmt.Println("---------------------1", l)
		solve(m, x, y, x1, y1, l/2)
		// 中间的三个是假公主坐标[](https://cdn.luogu.com.cn/upload/image_hosting/j0zm0q2q.png)
		// 右上
		// fmt.Println("---------------------2", l, x2, y2)
		solve(m, xrt, yrt, x2, y2, l/2)
		// 左下
		// fmt.Println("---------------------3", l)
		solve(m, xlb, ylb, x3, y3, l/2)
		// 右下
		// fmt.Println("---------------------4", l)
		solve(m, xrb, yrb, x4, y4, l/2)
	} else if x-a < l/2 && y-b >= l/2 {
		fmt.Printf("%d %d 2\n", xlb+1, ylb+1)
		// 左上
		solve(m, xlt, ylt, x1, y1, l/2)
		// 右上(自己)
		solve(m, x, y, x2, y2, l/2)
		// 左下
		solve(m, xlb, ylb, x3, y3, l/2)
		// 右下
		solve(m, xrb, yrb, x4, y4, l/2)
	} else if x-a >= l/2 && y-b < l/2 {
		fmt.Printf("%d %d 3\n", xrt+1, yrt+1)
		// 左上
		solve(m, xlt, ylt, x1, y1, l/2)
		// 右上
		solve(m, xrt, yrt, x2, y2, l/2)
		// 左下(自己)
		solve(m, x, y, x3, y3, l/2)
		// 右下
		solve(m, xrb, yrb, x4, y4, l/2)
	} else if x-a >= l/2 && y-b >= l/2 {
		fmt.Printf("%d %d 4\n", xlt+1, ylt+1)
		// 左上
		solve(m, xlt, ylt, x1, y1, l/2)
		// 右上
		solve(m, xrt, yrt, x2, y2, l/2)
		// 左下
		solve(m, xlb, ylb, x3, y3, l/2)
		// 右下(自己)
		solve(m, x, y, x4, y4, l/2)
	}
}
