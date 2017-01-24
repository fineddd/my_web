package net

import (
	p "github.com/golang/protobuf/proto"
	"common/net_base"
	"container/list"
	"log"
	"proto"
	"sync"
	"time"
)

const MAX_NET_MSG_COUNT = (1 << 7)

type LogicPackage net_base.LogicPackage

type IRecvPackageHdl interface {
	OnRecvPackage(pack *LogicPackage)
	OnTimeout()
}

type RecvPackageTimeout struct {
	timeout int64
	sync    int
}

type GmMsgEventHdl struct {
	typeHandles map[proto.MsgType]func(pack *LogicPackage)
	handles     map[int]IRecvPackageHdl
	timeoutList list.List
	sync        int
	lock        sync.Mutex
}

func (hdl *GmMsgEventHdl) RegisterMsgTypeHdl(msgType proto.MsgType, msgHdl func(pack *LogicPackage)) {
	hdl.lock.Lock()
	defer hdl.lock.Unlock()
	hdl.typeHandles[msgType] = msgHdl
}

func (hdl *GmMsgEventHdl) Send(session uint64, msgCode proto.MsgType, msg p.Message, packHdl IRecvPackageHdl) bool {
	hdl.lock.Lock()
	defer hdl.lock.Unlock()

	if len(hdl.handles) >= MAX_NET_MSG_COUNT {
		return false
	}

	for {
		hdl.sync++
		if hdl.sync >= MAX_NET_MSG_COUNT {
			hdl.sync = 0
		}
		if _, ok := hdl.handles[hdl.sync]; !ok {
			break
		}
	}

	var timeout RecvPackageTimeout
	timeout.timeout = time.Now().Unix() + 3
	timeout.sync = hdl.sync

	hdl.timeoutList.PushBack(timeout)

	hdl.handles[hdl.sync] = packHdl
	__pfNetMgr.SendMessage(session, msgCode, hdl.sync, msg)
	return true
}

func (hdl *GmMsgEventHdl) SendToGameServNoReply(pfID, servID int, session uint64, msgCode proto.MsgType, msg p.Message) bool {
	var inf proto.MsgTrasferInf
	var goal int32 = 1
	var platformID int32 = int32(pfID)
	var serverID int32 = int32(servID)
	inf.Goal = &goal
	inf.PfID = &platformID
	inf.ServID = &serverID

	data, err := p.Marshal(msg)
	if err != nil {
		log.Println(err)
		return false
	}
	var d *[]byte
	if d, err = net_base.EncodePackage(data, msgCode, 0); err != nil {
		log.Println(err)
		return false
	}
	inf.Msg = *d
	__pfNetMgr.SendMessage(session, proto.MsgType_MSG_X2X_SYS_TRASFER_INF_, 0, &inf)
	return true
}

func (hdl *GmMsgEventHdl) SendToGameServ(pfID, servID int, session uint64, msgCode proto.MsgType, msg p.Message, packHdl IRecvPackageHdl) bool {
	var inf proto.MsgTrasferInf
	var goal int32 = 1
	var platformID int32 = int32(pfID)
	var areaID int32 = int32(servID)
	inf.Goal = &goal
	inf.PfID = &platformID
	inf.ServID = &areaID

	hdl.lock.Lock()
	defer hdl.lock.Unlock()

	if len(hdl.handles) >= MAX_NET_MSG_COUNT {
		return false
	}
	for {
		hdl.sync++
		if hdl.sync >= MAX_NET_MSG_COUNT {
			hdl.sync = 0
		}
		if _, ok := hdl.handles[hdl.sync]; !ok {
			break
		}
	}
	data, err := p.Marshal(msg)
	if err != nil {
		log.Println(err)
		return false
	}
	var d *[]byte
	if d, err = net_base.EncodePackage(data, msgCode, uint32(hdl.sync)); err != nil {
		log.Println(err)
		return false
	}
	inf.Msg = *d

	var timeout RecvPackageTimeout
	timeout.timeout = time.Now().Unix() + 3
	timeout.sync = hdl.sync

	hdl.timeoutList.PushBack(timeout)

	hdl.handles[hdl.sync] = packHdl
	__pfNetMgr.SendMessage(session, proto.MsgType_MSG_X2X_SYS_TRASFER_INF_, 0, &inf)
	return true
}

func (hdl *GmMsgEventHdl) OnRecvPackage(pack *net_base.LogicPackage) {
	if pack == nil {
		log.Println("pack is null")
		return
	}
	hdl.lock.Lock()
	defer hdl.lock.Unlock()

	if packHdl, ok := hdl.typeHandles[pack.Head.Type]; ok {
		packHdl((*LogicPackage)(pack))
	} else {
		packHdl, ok := hdl.handles[pack.Head.Sync]
		if !ok {
			log.Println("sync is not found")
			return
		}
		packHdl.OnRecvPackage((*LogicPackage)(pack))

		for e := hdl.timeoutList.Front(); e != nil; e = e.Next() {
			if e.Value.(RecvPackageTimeout).sync == pack.Head.Sync {
				hdl.timeoutList.Remove(e)
				break
			}
		}
		delete(hdl.handles, pack.Head.Sync)
	}

}

func (hdl *GmMsgEventHdl) CheckTimeoutPackage() {
	hdl.lock.Lock()
	defer hdl.lock.Unlock()
	nowTime := time.Now().Unix()
	var next *list.Element
	for e := hdl.timeoutList.Front(); e != nil; {
		if nowTime >= e.Value.(RecvPackageTimeout).timeout {
			packHdl, ok := hdl.handles[e.Value.(RecvPackageTimeout).sync]
			if !ok {
				log.Println("sync is not found")
			} else {
				packHdl.OnTimeout()
				delete(hdl.handles, e.Value.(RecvPackageTimeout).sync)
			}
			next = e.Next()
			hdl.timeoutList.Remove(e)
			e = next
		} else {
			break
		}
	}
}

func (hdl *GmMsgEventHdl) OnAccept(session uint64) {
	log.Println("new session ", session)
}

func (hdl *GmMsgEventHdl) OnDisconnect(session uint64) {
	log.Println("disconnect ", session)
	if session == GetInstance().GetCenterServSession() {
		GetInstance().SetCenterServSession(0)
	}
}

func NewGmMsgEventHdl() *GmMsgEventHdl {
	hdl := new(GmMsgEventHdl)
	hdl.sync = 0
	hdl.handles = make(map[int]IRecvPackageHdl)
	hdl.typeHandles = make(map[proto.MsgType]func(pack *LogicPackage))
	return hdl
}
