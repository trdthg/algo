package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	var arr = append(Arr, 123)
	fmt.Println(arr)
	// InsertSort3(arr)
	InsertSort3V2(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
