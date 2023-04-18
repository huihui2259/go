package log2

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Info  *log.Logger // 重要的信息
	Warn  *log.Logger // 警告信息
	Error *log.Logger // 错误信息
)

func init() {
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}

	Info = log.New(io.MultiWriter(file, ioutil.Discard), "Info: ", log.Ldate|log.Ltime|log.Lshortfile)

	Warn = log.New(os.Stdout, "Warn: ", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// func Info(s interface{}) {
// 	info.Println(s)
// }

// func Warn(s interface{}) {
// 	warn.Println(s)
// }

// func Error(s interface{}) {
// 	myerror.Println(s)
// }
