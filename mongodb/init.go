package mongodb

import (
	"common/mongo"
	"my_web/conf"
	"log"
	"time"
)

func CheckConnect() {
	for {
		s := mongo.GetSession()
		err := s.Ping()
		if err != nil {
			s.Close()
			if !mongo.Open(conf.Mongo.Addr, conf.Mongo.User, conf.Mongo.Password, conf.Mongo.DB, 10000) {
				log.Println("reconnect mongodb failed!")
			} else {
				log.Println("reconnect mongodb ok")
			}
		}
		// log.Println("mongodb is connecting")
		time.Sleep(time.Minute * 5)
	}
}

func init() {
	if !mongo.Open(conf.Mongo.Addr, conf.Mongo.User, conf.Mongo.Password, conf.Mongo.DB, 10000) {
		log.Println("db open failed!")
		return
	}
	log.Println("mongoDB init ok!")
	go CheckConnect()
}
