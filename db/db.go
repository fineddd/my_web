package db

import (
	"common/mysql"
	"my_web/conf"
	"log"
	"strconv"
	"time"
)

func startDb() {
	for {
		time.Sleep(time.Millisecond)
		mysql.Update()
	}
}

func init() {
	host := "tcp(" + conf.DbHost + ":" + strconv.Itoa(conf.DbPort) + ")"
	b := mysql.Open(host, conf.DbUser, conf.DbPassword, conf.DbName, "utf8", 1000)
	if !b {
		log.Println("can't open %s!", conf.DbName)
		return
	}
	log.Println("db init ok")
	go startDb()
}
