package util

type Sortable interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

func Sort(s Sortable, f func(s Sortable, begin, end int)) {
	begin, end := 0, s.Len()-1
	SortPartial(s, begin, end, f)
}

func SortPartial(s Sortable, begin, end int, f func(s Sortable, begin, end int)) {
	if !preCheck(s, begin, end) {
		panic("begin, end不合法")
	}
	f(s, begin, end)
}

func preCheck(s Sortable, begin, end int) bool {
	if begin < 0 {
		return false
	}
	if begin > end {
		return false
	}
	if end >= s.Len() {
		return false
	}
	return true
}

func BubbleSort(s Sortable, begin, end int) {
	swapped := false
	for i := begin; i < end+1; i++ {
		swapped = false
		for j := begin; j < end-i; j++ {
			if s.Less(j, j+1) {
				s.Swap(j, j+1)
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
