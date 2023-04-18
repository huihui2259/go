package main

import (
	_ "goDemo/conf"
	"log"

	"goDemo/kafka"
	"goDemo/router"
)

// git token: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
func main() {
	go kafka.Receive()
	r := router.InitRouter()
	if err := r.Run(":8089"); err != nil {
		log.Fatal("服务器启动失败...")
	}

}
