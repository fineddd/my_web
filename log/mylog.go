package log

import (
	"fmt"
	"log"
	"os"
	"time"
	"runtime"
	"strconv"
)

type LogType int

const (
	LOG_Debug LogType = iota
	LOG_Error
	LOG_Fatal
)
var mapLog map[LogType]*log.Logger

func Println(strPlayerID string, typeLevel LogType, strMsg string) {
	pc, _, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	mapLog[typeLevel].Println(strPlayerID + "|" + f.Name() + "|" + strconv.Itoa(line) + "|" + strMsg)
	return
}
func Print(strPlayerID string, typeLevel LogType, strMsg string) {
	pc, _, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	mapLog[typeLevel].Print(strPlayerID + "|" + f.Name() + "|" + strconv.Itoa(line) + "|" + strMsg)
	return
}

func Printf(strPlayerID string, typeLevel LogType, format string, strMsg string) {
	pc, _, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	mapLog[typeLevel].Printf(format, strPlayerID + "|" + f.Name() + "|" + strconv.Itoa(line) + "|" + strMsg)
	return
}

func init() {
	now := time.Now().Format("2006010215")
	mapLog = make(map[LogType]*log.Logger)

	debugFile, err := os.OpenFile("debug_"+now+".log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("create debug log file failed!")
		return
	}
	mapLog[LOG_Debug] = log.New(debugFile, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)

	errorFile, err := os.OpenFile("error_"+now+".log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("create error log file failed!")
		return
	}
	mapLog[LOG_Error] = log.New(errorFile, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)

	fatalFile, err := os.OpenFile("fatal_"+now+".log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("create fatal log file failed!")
		return
	}
	mapLog[LOG_Fatal] = log.New(fatalFile, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)
	log.Println("log init complete")
}
