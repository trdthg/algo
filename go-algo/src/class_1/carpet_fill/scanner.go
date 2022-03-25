package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Sc interface {
	ReadInt() (int, error)
	ReadString() (string, error)
	Read(format string, a ...interface{}) error
}

type ReaderScanner struct {
	reader io.Reader
}

func NewReaderScanner(reader io.Reader) *ReaderScanner {
	return &ReaderScanner{reader}
}

func NewReader() *io.Reader {
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
	return &reader
}

func (rs *ReaderScanner) ReadInt() (res int, err error) {
	_, err = fmt.Fscan(rs.reader, &res)
	if err != nil {
		// ExitWithMsg(1, "输入格式错误(int)")
		fmt.Print("输入格式错误(int), 请重新输入: ")
		return rs.ReadInt()
	}
	return
}

func (rs *ReaderScanner) ReadIntWithMsg(msg string) (res int, err error) {
	fmt.Print(msg)
	return rs.ReadInt()
}

func (rs *ReaderScanner) ReadString() (res string, err error) {
	_, err = fmt.Fscan(rs.reader, &res)
	if err != nil {
		// ExitWithMsg(1, "输入格式错误(string)")
		fmt.Print("输入格式错误(string), 请重新输入: ")
		return rs.ReadString()
	}
	return
}

func (rs *ReaderScanner) ReadStringWithMsg(msg string) (res string, err error) {
	fmt.Print(msg)
	return rs.ReadString()
}

func (rs *ReaderScanner) read(a []interface{}) (err error) {
	for _, v := range a {
		switch v := v.(type) {
		case (*int):
			_, err = fmt.Fscan(rs.reader, v)
		case (*uint32):
			_, err = fmt.Fscan(rs.reader, v)
		case (*string):
			_, err = fmt.Fscan(rs.reader, v)
		default:
			fmt.Printf("unsupported type: %V\n", v)
		}
		if err != nil {
			// ExitWithMsg(1, "输入格式错误")
			fmt.Print("输入格式错误, 请重新输入: ")
			return rs.read(a)
		}
	}
	return
}
func (rs *ReaderScanner) Read(a ...interface{}) (err error) {
	return rs.read(a)
}
func (rs *ReaderScanner) ReadWithMsg(msg string, a ...interface{}) (err error) {
	fmt.Print(msg)
	return rs.read(a)
}
