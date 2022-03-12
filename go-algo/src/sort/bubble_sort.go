package sort

// 相邻的交换顺序
// 缺点: 必须比较一遍
func BubbleSort(array []int) {
	len := len(array)
	swapped := false
	for i := 0; i < len; i++ {
		swapped = false
		// 每一轮中最大的一定会被排到末尾，下一轮就能少比较一次
		for j := 0; j < len-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
		// 某一轮没有发生交换，就表示已经排序完成
		if !swapped {
			// return
		}
	}
}

func BubbleSortV2(array []int, begin, end int) {
	if begin > end || end >= len(array) {
		return
	}
	swapped := false
	for i := begin; i < end+1; i++ {
		swapped = false
		// 每一轮中最大的一定会被排到末尾，下一轮就能少比较一次
		for j := begin; j < end-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
		// 某一轮没有发生交换，就表示已经排序完成
		if !swapped {
			return
		}
	}

}
