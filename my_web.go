package main

import (
	"runtime"
	//"my_web/log"
	"my_web/web"
	_ "my_web/conf"
	_ "my_web/db"
	"log"
)

func main() {
	runtime.GOMAXPROCS(100)
	log.Println("main start")
	web.StartHttp()
	log.Println("main end")
}
