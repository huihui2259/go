package log1

import "fmt"

var Log *Logger

func init() {
	Log = CreateLogger()
}

// 创建日志器
func CreateLogger() *Logger {
	// 创建日志器
	l := NewLogger()
	// 创建命令行写入器
	cw := NewConsoleWriter()
	//注册命令行写入器到日志器中
	l.RegisterWriter(cw)
	// 创建文件写入器
	fw := NewFileWriter()
	// 设置文件名
	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}
	// 注册文件写入器到日志器中
	l.RegisterWriter(fw)
	return l
}
