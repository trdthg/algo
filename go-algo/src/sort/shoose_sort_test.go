package sort

import (
	"fmt"
	"testing"
)

func TestChooseSort(t *testing.T) {
	var arr = append(Arr, 1)
	fmt.Println(arr)
	// ChooseSort(arr)
	ChooseSortV2(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
