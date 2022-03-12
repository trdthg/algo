package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/trdthg/algo/go-algo/src/util"
)

func solution(reader io.Reader) {
	org := getInput(reader)
	fmt.Println(org.students)
	util.Sort(ByTreePlanted{org.students}, util.BubbleSort)
	fmt.Println(org.students)
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

func getInput(reader io.Reader) *Org {
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
	return &org
}
