package utils

import "fmt"

type MyError struct {
	Message1 string // 函数返回错误
	Message2 string // 自定义错误
}

func (e *MyError) Error() string {
	return fmt.Sprintf("my error: %s", e.Message1)
}

func (e *MyError) Print() {
	fmt.Println("错误1: " + e.Message1)
	fmt.Println("错误2: " + e.Message2)
}

func (e *MyError) Nil() bool {
	if e != nil && len(e.Message1) > 0 && len(e.Message2) > 0 {
		return false
	}
	return true
}

func (e *MyError) String() string {
	return fmt.Sprintf("first error: %s, second error: %s", e.Message1, e.Message2)
}
