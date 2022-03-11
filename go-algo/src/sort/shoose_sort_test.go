package sort

import (
	"fmt"
	"testing"
)

func TestChooseSort(t *testing.T) {
	var arr = append(Arr, 1)
	ChooseSort(arr)
	fmt.Println(arr)
}
