package sort

func QuickSort_v0(arr []int) {
	QuickSort(arr, 0, len(arr)-1)
}

func QuickSort(arr []int, begin, end int) {
	if begin < end {
		loc := partition_3(arr, begin, end)
		QuickSort(arr, begin, loc-1)
		QuickSort(arr, loc+1, end)
	}
}

func QuickSort_v2(arr []int, begin, end int) {
	if begin < end {
		if end-begin <= 4 {
			InsertSort3V2(arr, begin, end)
			return
		}
		loc := partition_1(arr, begin, end)
		QuickSort(arr, begin, loc-1)
		QuickSort(arr, loc+1, end)
	}
}

func QuickSort_v4(arr []int, begin, end int) {
	if begin < end {
		loc := partition_4(arr, begin, end)
		QuickSort(arr, begin, loc-1)
		QuickSort(arr, loc+1, end)
	}
}

func QuickSort_v5(arr []int, begin, end int) {
	if begin < end {
		ll, lr := partition_5(arr, begin, end)
		QuickSort(arr, begin, ll-1)
		QuickSort(arr, lr+1, end)
	}
}

// 交换法(依次把大于pivot的扔到最后面)
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
	pivot := arr[begin]

	l, r := begin, end
	for l < r {
		// 从右向左找到小的，让小的填上左边的坑位，r的位置是新的坑位
		for l < r && arr[r] >= pivot {
			r -= 1
		}
		arr[l] = arr[r]
		// 从左向右找大的，填不右边的坑位
		for l < r && arr[l] < pivot {
			l += 1
		}
		arr[r] = arr[l]
	}
	arr[l] = pivot
	return l
}

// 真交换法，左右同时推进
// if l > pivot && r < pivot { swap(l, r) }
func partition_3(arr []int, begin, end int) int {
	// arr[begin]是基准
	pivot := arr[begin]

	l, r := begin, end
	for l < r {
		// 从右向左找小的
		for l < r && arr[r] >= pivot {
			r -= 1
		}
		// 从左向右找大的
		for l < r && arr[l] <= pivot {
			l += 1
		}
		if l >= r {
			break
		}
		// 大小左右换位
		arr[l], arr[r] = arr[r], arr[l]
	}
	// 哨兵和l换位
	arr[begin], arr[l] = arr[l], arr[begin]
	return l
}

// 三向取中, 优化了寻找pivot的方法
func partition_4(arr []int, begin, end int) int {
	// arr[begin]是基准
	mid := (begin + end) / 2
	l, r := begin, end
	if arr[l] > arr[r] {
		arr[l], arr[r] = arr[r], arr[l]
	}
	if arr[mid] > arr[r] {
		arr[mid], arr[r] = arr[r], arr[mid]
	}
	if arr[mid] > arr[l] {
		arr[mid], arr[l] = arr[l], arr[mid]
	}
	pivot := arr[l]
	for l < r {
		// 从右向左找小的
		for l < r && arr[r] >= pivot {
			r -= 1
		}
		// 从左向右找大的
		for l < r && arr[l] < pivot {
			l += 1
		}
		if l > r {
			break
		}
		// 大小左右换位
		arr[l], arr[r] = arr[r], arr[l]
	}
	// 哨兵和l换位
	arr[r], arr[l] = arr[l], arr[r]
	return l
}

// 三向切分, 类似第一种，大的全部放到右边，小的全部放到左边，等于的就mid++
// 但是左右边界指针不移动，等于的都聚在中间，这样中间的就不用重新排序
func partition_5(arr []int, begin, end int) (int, int) {
	l, r := begin, end
	mid := begin + 1
	pivot := arr[begin]
	// 以中间坐标为准
	for mid <= r {
		if arr[mid] > pivot { // 大于基准数，那么交换，右指针左移
			arr[mid], arr[r] = arr[r], arr[mid]
			r--
		} else if arr[mid] < pivot { // 小于基准数，那么交换，左指针右移
			arr[mid], arr[l] = arr[l], arr[mid]
			l++
			mid++
		} else {
			mid++
		}
	}

	return l, r
}
