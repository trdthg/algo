package util

import (
	"fmt"
	"os"
)

func Exit(code int) {
	os.Exit(code)
}

func ExitWithMsg(code int, msg string) {
	fmt.Println(msg)
	os.Exit(code)
}
