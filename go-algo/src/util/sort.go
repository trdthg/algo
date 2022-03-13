package util

type Sortable interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

func Sort(s Sortable, f func(s Sortable, begin, end int)) {
	begin, end := 0, s.Len()
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
	if end > s.Len() {
		return false
	}
	return true
}

func BubbleSort(s Sortable, begin, end int) {
	swapped := false
	for i := begin; i < end; i++ {
		swapped = false
		for j := begin; j < end-i-1; j++ {
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

func InsertSort(s Sortable, begin, end int) {
	for i := begin + 1; i < end; i++ {
		for j := i - 1; j >= begin && s.Less(j+1, j); j-- {
			s.Swap(j+1, j)
		}
	}
}

func QuickSort(s Sortable, begin, end int) {
	if begin < end {
		ll, rl := partition()
	}
}

func partition(s Sortable, begin, end int) (int, int) {
	mid := (begin + end) / 2
	l, r := begin, end
	// find the biggest
	if s.Less(r, l) {
		s.Swap(l, r)
	}
	if s.Less(r, mid) {
		s.Swap(r, mid)
	}
	// sort the lower two
	if s.Less(l, mid) {
		s.Swap(l, mid)
	}
	

}
