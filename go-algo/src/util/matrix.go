package util

import "fmt"

func NewMatrix(m, n int) [][]int {
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
	}
	return a
}

func NewCharMatrix(m, n int) [][]uint8 {
	a := make([][]uint8, m)
	for i := range a {
		a[i] = make([]uint8, n)
		for j := 0; j < n; j++ {
			a[i][j] = ' '
		}
	}
	return a
}

func PrintMatrix(a [][]int) {
	for _, v := range a {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
