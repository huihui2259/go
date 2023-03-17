package main

import (
	_ "goDemo/conf"
	"log"

	"goDemo/router"
)

func main() {
	r := router.InitRouter()
	if err := r.Run(":8089"); err != nil {
		log.Fatal("服务器启动失败...")
	}

}
