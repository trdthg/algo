package sort

import (
	"fmt"
	"testing"
)

// 全部sort
func TestMergeSortTtob(t *testing.T) {
	arr := Arr
	fmt.Println(arr)
	MergeSortTtob(arr, 0, len(arr))
	fmt.Println(arr)
}

func TestMergeSortBtot(t *testing.T) {
	arr := Arr
	fmt.Println(arr)
	MergeSortBtot(arr, 0, len(arr))
	fmt.Println(arr)
}

// 只sort一部分
func TestMergeSortPartial(t *testing.T) {
	arr := Arr
	fmt.Println(arr)
	MergeSortByBlockReverse(arr, 0, len(arr))
	fmt.Println(arr)
}
