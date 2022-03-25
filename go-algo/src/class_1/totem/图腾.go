package main

import (
	"fmt"

)

func main() {
	a := NewCharMatrix(1000, 1000)
	h, w := 2, 4
	// a[0][0] = ' '
	// a[0][3] = ' '
	// a[0][1] = '/'
	// a[0][2] = '\\'
	// a[1][0] = '/'
	// a[1][1] = '_'
	// a[1][2] = '_'
	// a[1][3] = '\\'
	a[0][0] = '\\'
	a[0][1] = '_'
	a[0][2] = '_'
	a[0][3] = '/'
	a[1][0] = ' '
	a[1][1] = '\\'
	a[1][2] = '/'
	a[1][3] = ' '
	var i, j, k int
	n, _ := NewReaderScanner(*NewReader()).ReadIntWithMsg("请输入图腾个数: ")
	// 从左下开始复制
	for i = 1; i < n; i++ {
		//向下和向右
		// fmt.Println(n, w, h)
		// 右
		for j = 0; j < h; j++ {
			for k = w; k < w<<1; k++ {
				a[j][k] = a[j][k-w]
			}
		}
		// 向下
		for j = h; j < h<<1; j++ {
			for k = w / 2; k < w*3/2; k++ {
				// fmt.Println(k)
				a[j][k] = a[j-h][k-w/2]
			}
		}
		w *= 2
		h *= 2
		//刷新完成一次
	}
	for i = 0; i < h+1; i++ {
		for j = 0; j < w+1; j++ {
			// fmt.Printf("%c", a[h-i][w-j])
			fmt.Printf("%c", a[i][j])
		}
		fmt.Println()
	}
	for i = 0; i < h+1; i++ {
		for j = 0; j < w+1; j++ {
			fmt.Printf("%c", a[h-i][w-j])
		}
		fmt.Println()
	}
}
