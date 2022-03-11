package sort

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := append(Arr, 111)
	BubbleSort(arr)
	fmt.Println(arr)
	_, err := strconv.Atoi("1")
	if err != nil {
		t.Error("aa")
	}
}
