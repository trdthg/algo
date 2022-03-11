package sort

// 合并两个有序数组
func MergeTwoSortedArray(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l += 1
		} else {
			result = append(result, right[r])
			r += 1
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

// 合并n个有序数组
func MergeSortedArray(arrays ...[]int) (result []int) {
	switch len(arrays) {
	case 0:
		result = []int{}
	case 1:
		result = arrays[0]
	case 2:
		result = MergeTwoSortedArray(arrays[0], arrays[1])
	default:
		len := len(arrays)
		result = MergeSortedArray(MergeSortedArray(arrays[:len/2]...), MergeSortedArray(arrays[len/2:]...))
	}
	return
}
