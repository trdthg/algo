package sort

// i=0; 每次对第i位找到[i+1:len]中的最小值填上去(交换位置); i++
// 缺点: 找到最小元素可能很慢
func ChooseSort(array []int) {
	len := len(array)
	for i := 0; i < len; i++ {
		min := i
		// 找到它之后最小的和自己交换位置
		for j := i + 1; j < len; j++ {
			if array[min] > array[j] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
}
