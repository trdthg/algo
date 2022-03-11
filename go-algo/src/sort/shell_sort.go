package sort

// 取 d = 6 对 [5 x x x x x 6 x x x x x] 进行直接插入排序，没有变化。
// 取 d = 3 对 [5 x x 6 x x 6 x x 4 x x] 进行直接插入排序，排完序后：[4 x x 5 x x 6 x x 6 x x]。
// 取 d = 1 对 [4 9 1 5 8 14 6 49 25 6 6 3] 进行直接插入排序，因为 d=1 完全就是直接插入排序了。

// 插入排序的优化版，思想在于分组
// 1 2 3 4 5 6 7 8 9
// 按照n的倍数进行分组，每组依次(其实是同时的)进行插入排序
func ShellSort(arr []int) {

	len := len(arr)
	// 每次步长都➗2
	for step := len / 2; step >= 1; step /= 2 {
		// 同时遍历所有组的元素，当step = 9 / 2 = 4 时
		// 1       5       9
		//   2       6
		//     3       7
		//       4      8
		for i := step; i < len; i++ {
			// 对某一组进行插入排序
			// j = 0 第一组1和5比较 结束
			// j = 1 第二组2和6比较 结束
			// j = 2 ...
			// j = 3 ...
			// j = 4 第一组5和9比较, j-step, 1和5比较
			// j = 5 ...
			for j := i - step; j >= 0 && arr[j+step] < arr[j]; j -= step {
				arr[j], arr[j+step] = arr[j+step], arr[j]
			}
		}
	}
}
