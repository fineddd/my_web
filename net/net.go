package net

import (
	"common/net_mgr"
	"log"
	"time"
	"proto"
)

type pfNetMgr struct {
	net_mgr.BaseTcpNetMgr
	centerServSession uint64
	handle            *GmMsgEventHdl
}

var __pfNetMgr *pfNetMgr

func (netMgr *pfNetMgr) SetCenterServSession(session uint64) {
	netMgr.centerServSession = session
}

func (netMgr *pfNetMgr) GetCenterServSession() uint64 {
	return netMgr.centerServSession
}

func GetInstance() *pfNetMgr {
	return __pfNetMgr
}

func GetHandle() *GmMsgEventHdl {
	return __pfNetMgr.handle
}

func Update() {
	for {
		time.Sleep(time.Second)
		__pfNetMgr.handle.CheckTimeoutPackage()
	}
}

func HeartBeat() {
	var hb proto.HeartBeat
	for {
		time.Sleep(time.Second*30)
		__pfNetMgr.SendMessage(__pfNetMgr.centerServSession, proto.MsgType_MSG_HEART_BEAT_, 0, &hb)
	}
}

func init() {
	__pfNetMgr = new(pfNetMgr)
	if b := __pfNetMgr.CommonInit(); b == false {
		log.Fatal("init net failed")
	}
	__pfNetMgr.handle = NewGmMsgEventHdl()
	__pfNetMgr.SetEventHdl(__pfNetMgr.handle)
	go Update()
//	go HeartBeat()
}
