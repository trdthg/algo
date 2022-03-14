package main

import "fmt"

type Major struct {
	id   int
	name string
}

func NewMajor(id int, name string) *Major {
	return &Major{id, name}
}

type Student struct {
	id               string `zh:"学号"`
	name             string `zh:"姓名"`
	tree_planted_num uint32 `zh:"植树量"`
	major            *Major `zh:"专业"`
}

func NewStudent(id, name string, tree_num uint32, m *Major) *Student {
	return &Student{id, name, tree_num, m}
}

type Group struct {
	major    *Major     `zh:"专业"`
	num      uint32     `zh:"专业总人数"`
	students []*Student `zh:"成员"`
}

func NewGroup(major *Major, num uint32, students []*Student) *Group {
	return &Group{major: major, num: num, students: students}
}

// type By func(i, j Student) bool

// func (b By) Sort(o *Org) {
// 	util.Sort(o, util.BubbleSort)
// }

type Org struct {
	groups   []*Group
	students students
}

type students []*Student

func (ss students) FindByTreePlanted(n int) (index int, cnt int) {
	len := ss.Len()
	l, mid, r := 0, len/2, len
	index, cnt = -1, 0
	for l < r {
		fmt.Println(l, r)
		cnt += 1
		tmp := int(ss[mid].tree_planted_num)
		if n < tmp {
			r = mid - 1
		} else if n == tmp {
			index = mid
			return
		} else {
			l = mid + 1
		}
		mid = (l + r) / 2
	}
	return
}

func (ss students) Len() int {
	return len(ss)
}
func (ss students) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
func (ss students) Less(i, j int) bool {
	return ss[i].tree_planted_num > ss[j].tree_planted_num
}

type ByTreePlanted struct {
	students
}

func (b ByTreePlanted) Less(i, j int) bool {
	return b.students[i].tree_planted_num < b.students[j].tree_planted_num
}
