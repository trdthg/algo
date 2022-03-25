package main

import "fmt"

// 6
// 3 2 4 5 6 1

// 6
func main() {
	var n int
	var i int
	fmt.Scanf("%d", &n)
	var a = make([]int, n)
	for i = 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	cnt := 0
	for i = 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] > a[j] {
				cnt += 1
			}
		}
	}
	fmt.Println(cnt)

}
