package sort

// 全部sort
func MergeSort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	tmp := MergeTwoSortedArray(MergeSort(arr[:len/2]), MergeSort(arr[len/2:]))
	for i := 0; i < len; i++ {
		arr[i] = tmp[i]
	}
	return MergeTwoSortedArray(MergeSort(arr[:len/2]), MergeSort(arr[len/2:]))
}

// 只sort一部分
func MergeSortPartial(arr []int, left, right int) {

}
