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
	n := 0
	n = sc.NextInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = sc.NextInt()
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr[0])
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
