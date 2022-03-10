package util

import (
	"fmt"
	"io"
)

type Sc interface {
	ReadInt() (int, error)
	ReadString() (string, error)
	Read(format string, a ...interface{}) error
}

type ReaderScanner struct {
	reader io.Reader
}

func NewReaderScanner(reader io.Reader) ReaderScanner {
	return ReaderScanner{reader}
}

func (rs *ReaderScanner) ReadInt() (res int, err error) {
	_, err = fmt.Fscan(rs.reader, &res)
	if err != nil {
		ExitWithMsg(1, "输入格式错误(int)")
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
		ExitWithMsg(1, "输入格式错误(string)")
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
			ExitWithMsg(1, "输入格式错误")
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
