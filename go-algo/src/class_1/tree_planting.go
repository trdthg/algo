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

type Major struct {
	name string
}

func NewMajors() []Major {
	return []Major{{"软件"}, {"计算机"}}
}

type Student struct {
	id               string `zh:"学号"`
	name             string `zh:"姓名"`
	tree_planted_num uint32 `zh:"植树量"`
}

func NewStudent(id, name string, tree_num uint32) Student {
	return Student{id, name, tree_num}
}

type Group struct {
	major    Major     `zh:"专业"`
	num      uint32    `zh:"总人数"`
	students []Student `zh:"成员"`
}

func solution(reader io.Reader) {
	sc := util.NewReaderScanner(reader)
	var groups []Group
	major_num, _ := sc.ReadIntWithMsg("请输入专业数量: ")
	for i := 0; i < major_num; i++ {
		majorName, _ := sc.ReadStringWithMsg(fmt.Sprintf("请输入第%d个专业名称: ", i+1))
		studnetNum, _ := sc.ReadIntWithMsg(fmt.Sprintf("请输入%s专业学生数量: ", majorName))
		var students []Student
		for i := 0; i < studnetNum; i++ {
			var student Student
			_ = sc.ReadWithMsg(fmt.Sprintf("请输入第%d位学生信息[学号 姓名 植树量]: ", i+1), &student.id, &student.name, &student.tree_planted_num)
			students = append(students, student)
		}
		group := Group{major: Major{majorName}, num: uint32(studnetNum), students: students}
		groups = append(groups, group)
	}
	fmt.Println()
	fmt.Println(groups)
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
