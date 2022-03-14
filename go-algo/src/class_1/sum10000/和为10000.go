package main

import (
	"fmt"

	"github.com/trdthg/algo/go-algo/src/util"
)

func main() {
	sc := util.NewReaderScanner(*util.NewReader())
	len1, _ := sc.ReadInt()
	arr1 := make([]int, len1)
	var n int
	for i := 0; i < len1; i++ {
		n, _ = sc.ReadInt()
		arr1[i] = n
	}
	len2, _ := sc.ReadInt()
	arr2 := make([]int, len2)
	for i := 0; i < len2; i++ {
		n, _ = sc.ReadInt()
		arr2[i] = n
	}
	fmt.Println(arr1, arr2)
	l1, l2 := HasSum(arr1, arr2, 5)
	fmt.Println(">>> ", l1, l2)
}

func HasSum(arr1, arr2 []int, sum int) (l1 int, l2 int) {
	// arr1升序
	l1, l2 = 0, 0
	len1, len2 := len(arr1), len(arr2)
	for l1 < len1 && l2 < len2 {
		tmp := arr1[l1] + arr2[l2]
		if tmp == sum {
			return l1, l2
		} else if tmp > sum {
			l2 += 1
		} else {
			l1 += 1
		}
	}
	return -1, -1
}
