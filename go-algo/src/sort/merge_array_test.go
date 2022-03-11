package sort

import (
	"fmt"
	"testing"
)

var arrays = [][]int{
	{1, 2, 3},
	{1, 2, 3},
	{1, 2, 3},
	{1, 2, 3},
	{1, 2, 3},
}

// 合并两个有序数组
func TestMergeTwoSortedArray(t *testing.T) {
	a := MergeTwoSortedArray(arrays[0], arrays[1])
	fmt.Println(a)
}

func TestMergeSortedArray(t *testing.T) {
	a := MergeSortedArray(arrays...)
	fmt.Println(a)
}
