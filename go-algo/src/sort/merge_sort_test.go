package sort

import (
	"fmt"
	"testing"
)

// 全部sort
func TestMergeSort(t *testing.T) {
	arr := Arr
	fmt.Println(arr)
	MergeSort(arr)
	fmt.Println(arr)
}

// 只sort一部分
func TestMergeSortPartial(t *testing.T) {
	arr := Arr
	fmt.Println(arr)
	// mergeSort(arr)
	fmt.Println(arr)
}
