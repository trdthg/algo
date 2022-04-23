// 参加越多的比赛，noip 就能考的越好（假的）。
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := NewScanner()
	n := sc.NextInt()
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, 2)
		m[i][0] = sc.NextInt()
		m[i][1] = sc.NextInt()
	}
	sort.Slice(m, func(i, j int) bool {
		return m[i][1] < m[j][1]
	})
	sum := 0
	last_end := 0
	// fmt.Println(m)
	for _, race := range m {
		if start := race[0]; start >= last_end {
			sum += 1
			last_end = race[1]
		}
	}
	fmt.Println(sum)
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
