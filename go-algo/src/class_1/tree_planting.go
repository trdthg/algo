package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/trdthg/algo/go-algo/src/util"
)

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

type Org struct {
	groups   []*Group
	students []*Student
}

// 要排序的数组，用来比较的字段，比较规则，顺序
func Sort(array []*interface{}, cmp func(interface{}) int) {
	arrayCopy := make([]*interface{}, len(array))
	copy(arrayCopy, array)
	sort.Slice(arrayCopy, func(i, j int) bool {
		
	})
}

func solution(reader io.Reader) {
	sc := util.NewReaderScanner(reader)
	var org = Org{groups: []*Group{}, students: []*Student{}}
	major_num, _ := sc.ReadIntWithMsg("请输入专业数量: ")
	for i := 0; i < major_num; i++ {
		// 专业信息
		majorName, _ := sc.ReadStringWithMsg(fmt.Sprintf("请输入第%d个专业名称: ", i+1))
		studnetNum, _ := sc.ReadIntWithMsg(fmt.Sprintf("请输入%s专业学生数量: ", majorName))
		major := NewMajor(i, majorName)
		var students []*Student
		for i := 0; i < studnetNum; i++ {
			// 学生信息
			var student = &Student{}
			_ = sc.ReadWithMsg(fmt.Sprintf("请输入第%d位学生信息[学号 姓名 植树量]: ", i+1), &student.id, &student.name, &student.tree_planted_num)
			student.major = major
			// 加入组织
			org.students = append(org.students, student)
			students = append(students, student)
		}
		org.groups = append(org.groups, NewGroup(major, uint32(studnetNum), students))
	}
	fmt.Println(org.groups[0].students)
	util.ExitWithMsg(0, "正常退出")
}

func main() {
	filePath := flag.String("f", "test.txt", "test input's file path")
	flag.Parse()
	str, err := ioutil.ReadFile(*filePath)
	var reader io.Reader
	if err != nil {
		fmt.Printf("[WARN] 没有找到测试文件 '%s', 请用户手动输入\n", *filePath)
		reader = os.Stdin
	} else {
		reader = strings.NewReader(string(str))
	}
	solution(reader)
}
