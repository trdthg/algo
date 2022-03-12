package sort

// i=1; 每次把第i个元素插入到[0..i-1]里; i++
// 缺点: 最慢的情况下就是反着的冒泡
func InsertSort(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		// 待排数 i
		sort_num := arr[i]
		// 最后一个数 i-1
		var j = i - 1
		if sort_num >= arr[j] {
			continue
		}
		// 把第i个元素插入到[0..i-1]里
		for ; j >= 0 && sort_num < arr[j]; j-- {
			// j元素后移一位
			arr[j+1] = arr[j]
		}
		// i插入到j的位置
		arr[j+1] = sort_num
	}
}

func InsertSort2(arr []int) {
	len := len(arr)
	for i := 1; i < len; i++ {
		sort_num := arr[i]
		var j = i - 1
		if sort_num >= arr[j] {
			continue
		}
		for ; j >= 0 && sort_num < arr[j]; j-- {
			// j元素和j+1交换位置
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

func InsertSort3(arr []int) {
	for i := 1; i < len(arr); i++ {
		// 对某一组进行插入排序
		// i = 1, j = 0 0和1比较 j-1 结束
		// i = 2, j = 1 1和2比较 j-1 0和1比较 ... 结束
		// ...  , j = 2 2和3比较 ... 1和2比较 ... 0和1比较 ... 结束
		// ...
		for j := i - 1; j >= 0 && arr[j+1] < arr[j]; j -= 1 {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

func InsertSort3V2(arr []int, begin, end int) {
	if begin > end || end >= len(arr) {
		return
	}
	for i := begin + 1; i < end+1; i++ {
		// 对某一组进行插入排序
		// i = 1, j = 0 0和1比较 j-1 结束
		// i = 2, j = 1 1和2比较 j-1 0和1比较 ... 结束
		// ...  , j = 2 2和3比较 ... 1和2比较 ... 0和1比较 ... 结束
		// ...
		for j := i - 1; j >= begin && arr[j+1] < arr[j]; j -= 1 {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}
