package conf

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type MongoCnf struct {
	Addr     string
	DB       string
	User     string
	Password string
}

type Config struct {
	DbUser         string
	DbPassword     string
	DbHost         string
	DbPort         int
	DbName         string
	Port           int
	CenterServAddr string
	MDB            MongoCnf
}

var DbUser string
var DbPassword string
var DbHost string
var DbPort int
var DbName string
var Port int
var CenterServAddr string
var Mongo MongoCnf

func init() {
	fp, err := os.Open("config")
	defer fp.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var cnf Config
	buf := bufio.NewReader(fp)
	cfline, _ := buf.ReadString(0)
	bline := []byte(cfline)
	err = json.Unmarshal(bline, &cnf)
	if err != nil {
		fmt.Println(err)
		return
	}
	DbUser = cnf.DbUser
	DbPassword = cnf.DbPassword
	DbHost = cnf.DbHost
	DbPort = cnf.DbPort
	DbName = cnf.DbName
	Port = cnf.Port
	CenterServAddr = cnf.CenterServAddr
	Mongo = cnf.MDB
}
