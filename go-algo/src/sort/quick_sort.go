package sort

func QuickSort(arr []int, begin, end int) {
	if begin < end {
		loc := partition_1(arr, begin, end)
		QuickSort(arr, begin, loc-1)
		QuickSort(arr, loc+1, end)
	}
}

// 交换法(依次把大于guard的扔到最后面)
func partition_1(arr []int, begin, end int) int {
	// arr[begin]是基准
	l, r := begin+1, end
	for l < r {
		// 从左向右找大的，把大的放在右边，r--,
		if arr[l] > arr[begin] {
			arr[l], arr[r] = arr[r], arr[l]
			r -= 1
		} else {
			l += 1
		}
	}
	if arr[l] >= arr[begin] {
		l -= 1
	}
	arr[begin], arr[l] = arr[l], arr[begin]
	return l
}

// 填坑法(初始坑位是第一个)
func partition_2(arr []int, begin, end int) int {
	// arr[begin]是基准
	guard := arr[begin]

	l, r := begin, end
	for l < r {
		// 从右向左找到小的，让小的填上左边的坑位，r的位置是新的坑位
		for l < r && arr[r] >= guard {
			r -= 1
		}
		arr[l] = arr[r]
		// 从左向右找大的，填不右边的坑位
		for l < r && arr[l] < guard {
			l += 1
		}
		arr[r] = arr[l]
	}
	arr[l] = guard
	return l
}

// 真交换法
func partition_3(arr []int, begin, end int) int {
	// arr[begin]是基准
	guard := arr[begin]

	l, r := begin+1, end
	for l < r {
		// 从右向左找小的
		for l < r && arr[r] >= guard {
			r -= 1
		}
		// 从左向右找大的
		for l < r && arr[l] < guard {
			l += 1
		}
		if l > r {
			break
		}
		// 大小左右换位
		arr[l], arr[r] = arr[r], arr[l]
	}
	// 哨兵和l换位
	arr[begin], arr[l] = arr[l], arr[begin]
	return l
}
