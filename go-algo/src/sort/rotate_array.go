package sort

func Rotation(arr []int, begin, mid, end int) {
	reverse(arr, begin, mid)
	reverse(arr, mid, end)
	reverse(arr, begin, end)
}

func reverse(arr []int, l, r int) {
	r -= 1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l += 1
		r -= 1
	}
}
