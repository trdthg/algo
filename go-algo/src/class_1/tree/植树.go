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
	sc := util.NewReaderScanner(reader)
	org := getInput(sc)
	fmt.Println(org.students)
	// util.Sort(ByTreePlanted{org.students}, util.BubbleSort)
	util.Sort(ByTreePlanted{org.students}, util.QuickSort)
	// util.Sort(ByTreePlanted{org.students}, util.InsertSort)
	for _, s := range org.students {
		fmt.Println(s.name, s.tree_planted_num)
	}
	num, _ := sc.ReadIntWithMsg("请输入植树数量: ")
	index, cnt := org.students.FindByTreePlanted(num)
	if index != -1 {
		s := org.students[index]
		fmt.Printf("该生排名在第%d位, 查找%d次成功, %s %s %s %d", index+1, cnt, s.major.name, s.id, s.name, s.tree_planted_num)
	} else {
		fmt.Printf("查询失败, 没有种植数量为%d的学生", num)
	}
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

func getInput(sc *util.ReaderScanner) *Org {
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
