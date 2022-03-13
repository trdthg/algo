package sort

import (
	"fmt"
	"testing"
)

func TestRotation(t *testing.T) {
	arr := Arr
	fmt.Println(arr)
	Rotation(arr, 0, len(arr)/2, len(arr))
	fmt.Println(arr)
}
