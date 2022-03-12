package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := Arr
	fmt.Println(arr)
	QuickSort_v3(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
