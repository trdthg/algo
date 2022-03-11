package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	var arr = append(Arr, 123)
	fmt.Println(arr)
	InsertSort(arr)
	fmt.Println(arr)
}
