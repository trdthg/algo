package sort

import (
	"fmt"
	"testing"
)

// 取 d = 6 对 [5 x x x x x 6 x x x x x] 进行直接插入排序，没有变化。
// 取 d = 3 对 [5 x x 6 x x 6 x x 4 x x] 进行直接插入排序，排完序后：[4 x x 5 x x 6 x x 6 x x]。
// 取 d = 1 对 [4 9 1 5 8 14 6 49 25 6 6 3] 进行直接插入排序，因为 d=1 完全就是直接插入排序了。

// 插入排序的优化版，思想在于分组
// 1 2 3 4 5 6 7 8 9
// 按照n的倍数进行分组，每组依次进行插入排序
// 比如
func TestShellSort(t *testing.T) {
	arr := Arr
	fmt.Println(arr)
	ShellSort(arr)
	fmt.Println(arr)
}
