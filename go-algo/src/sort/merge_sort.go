package sort

// 自顶向下, 递归
func MergeSortTtob(arr []int, begin, end int) {
	// 元素数量大于一开始递归
	if end-begin > 1 {
		// 1,2,3
		mid := (begin + end) / 2
		MergeSortTtob(arr, begin, mid)
		MergeSortTtob(arr, mid, end)
		mergeByAppend(arr, begin, mid, end)
	}
}

// 自底向上,非递归
func MergeSortBtot(arr []int, begin, end int) {
	step := 1
	for (end-begin)*2 > step {
		// 0 1 2 3 4 5 6 [0,6)
		// 1 | 0 1 2
		// 1 | 2 3 4
		// 1 | 4 5 6
		// 2 | 0 2 4
		// 2 | 4 6 8 -> 4 6 x-x (没有第二部分)
		// 4 | 0 4 8 -> 0-4 4-6 (第二部分没满)
		for i := begin; i < end; i += (step << 1) {
			l, mid, r := i, i+step, i+(step<<1)
			// 不用合并了
			if mid > end {
				continue
			}
			// 合并
			if r > end {
				r = end
			}
			mergeByAppend(arr, l, mid, r)
		}
		step <<= 1
	}
}

// 使用临时数组
func mergeByAppend(arr []int, begin, mid, end int) {
	l, r := begin, mid
	result := make([]int, 0, end-begin)
	for l < mid && r < end {
		if arr[l] <= arr[r] {
			result = append(result, arr[l])
			l += 1
		} else {
			result = append(result, arr[r])
			r += 1
		}
	}
	result = append(result, arr[l:mid]...)
	result = append(result, arr[r:end]...)
	for i := begin; i < end; i++ {
		arr[i] = result[i-begin]
	}
}

func MergeSortByBlockReverse(arr []int, begin, end int) {
	blockSize := 3
	// 分块进行插入排序
	l, r := begin, blockSize
	for r <= end {
		InsertSort(arr[l:r])
		l = r
		r += blockSize
	}
	InsertSort(arr[l:end])

	mid := begin + begin
	for mid < end {
		l = begin
		mid = l + blockSize
		r = l + blockSize<<1
		// 每次合并两个block
		for r <= end {
			mergeByReverse(arr, l, (l+r)/2, r)
			l = r
			r += blockSize << 1
		}
		// 3 + 1
		if r > end {
			mergeByReverse(arr, l, (l+r)/2, end)
		}
		blockSize <<= 1
	}
}

// 使用👋摇算法, 不需要临时数组
func mergeByReverse(arr []int, begin, mid, end int) {
	l, r := begin, mid
	for l < mid && r < end {
		step := 0
		for l < mid && arr[l] <= arr[r] {
			l += 1
		}
		for r < end && arr[r] <= arr[l] {
			r += 1
			step += 1
		}
		Rotation(arr, l, mid, r)
		l += step
		mid = r
	}
}
