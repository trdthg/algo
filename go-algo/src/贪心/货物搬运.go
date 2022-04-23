package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := NewScanner()
	_, m := sc.GetVec()
	sum := 0
	solve(m, &sum)
	println(sum)
}

func solve(arr []int, sum *int) {
	if len(arr) == 0 {
		return
	}
	if len(arr) == 1 {
		fmt.Println("good", arr[0], *sum)
		*sum += arr[0]
		return
	}
	cut_point, min := find_min_index(arr)
	fmt.Println(">>", min, len(cut_point), *sum)
	*sum += min * len(cut_point)
	for i := range arr {
		arr[i] -= min
	}
	cut_point = append(cut_point, len(arr))
	for _, v := range cut_point {
		fmt.Println("|>", v)
	}
	l := 0
	for _, r := range cut_point {
		next_arr := arr[l:r]
		solve(next_arr, sum)
		l = r + 1
	}

	return
}

func find_min_index(arr []int) ([]int, int) {
	min := 1000000
	min_i := []int{}
	for i, v := range arr {
		if v < min {
			min_i = min_i[0:0]
			min = v
			min_i = append(min_i, i)
		} else if v == min {
			min_i = append(min_i, i)
		}
	}
	return min_i, min
}

type Scanner struct {
	reader bufio.Reader
	buf    []string
	index  int
	buff   []byte
	temp   string
}

func NewScanner() Scanner {
	return Scanner{
		reader: *bufio.NewReader(os.Stdin),
		buf:    []string{},
		index:  0,
		buff:   []byte{},
	}
}

func (s *Scanner) next() *Scanner {
	for {
		if len(s.buf) != 0 {
			next := s.buf[s.index]
			s.index += 1
			if s.index == len(s.buf) {
				s.buf = s.buf[:0]
			}
			s.temp = next
			return s
		} else {
			s.buff, _, _ = s.reader.ReadLine()
			s.buf = strings.Fields(string(s.buff))
			s.buff = s.buff[:0]
			s.index = 0
		}
	}
}

func (s *Scanner) NextStr() string {
	return s.next().temp
}

func (s *Scanner) NextInt() int {
	res, _ := strconv.Atoi(s.next().temp)
	return res
}

func (s *Scanner) GetVec() (int, []int) {
	n := 0
	n = s.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = s.NextInt()
	}
	return n, arr
}

func (s *Scanner) GetMatrix() (int, [][]int) {
	n := 0
	n = s.NextInt()
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			arr[i][j] = s.NextInt()
		}
	}
	return n, arr
}
