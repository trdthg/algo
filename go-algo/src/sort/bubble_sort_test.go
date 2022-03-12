package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := append(Arr, 111)
	fmt.Println(arr)
	BubbleSortV2(arr, 0, len(arr)-1)
	// BubbleSort_v2(arr)
	fmt.Println(arr)
}
