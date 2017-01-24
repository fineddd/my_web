package user

import (
	"common/mysql"
	"errors"
	"log"
)

type LoadUserReq struct {
	users *map[int]*User
	ch    chan error
}

func (req *LoadUserReq) OnExecute(dq *mysql.DBQuery) bool {
	if !dq.Prepare("select `id`,`name`,`password`,`right`,`address`,`lastlogintime`,`status`,`pfid` from `user`") {
		req.ch <- errors.New("prepare failed")
		return false
	} else if !dq.Query() {
		req.ch <- errors.New("query failed")
		return false
	}
	var id int
	var name string
	var password string
	var right int
	var address string
	var lastLoginTime int64
	var status int
	var pfID int
	var rightName string
	var platformName string

	for dq.NextRecord(&id, &name, &password, &right, &address, &lastLoginTime, &status, &pfID) {
		rightName, _= GetRightName(right)
		platformName, _= GetPlatformName(pfID)
		(*req.users)[id] = &User{ID: id, PfID: pfID, Name: name, Password: password,
			Right: right, RightName:rightName, PfName:platformName, Address: address, Status: status, LastLoginTime: lastLoginTime}
	}
	req.ch <- nil
	return false
}

func (req *LoadUserReq) SyncResultHdl() bool {
	return true
}

type AddUserReq struct {
	id            *int
	name          string
	password      string
	right         int
	address       string
	lastLoginTime int64
	status        int
	pfID          int
	ch            chan error
}

func (req *AddUserReq) OnExecute(dq *mysql.DBQuery) bool {
	if !dq.Prepare("insert into `user`(`name`,`password`,`right`,`address`,`lastlogintime`,`status`,`pfid`) values(?,?,?,?,?,?,?)") {
		req.ch <- errors.New("prepair failed")
		return false
	} else if !dq.Exec(req.name, req.password, req.right, req.address, req.lastLoginTime, req.status, req.pfID) {
		req.ch <- errors.New("exec failed")
		return false
	}
	if !dq.Prepare("select last_insert_id()") {
		req.ch <- errors.New("prepare failed")
		return false
	} else if !dq.Query() {
		req.ch <- errors.New("query failed")
		return false
	} else if !dq.NextRecord(req.id) {
		req.ch <- errors.New("get id failed")
		return false
	}
	req.ch <- nil
	return false
}

func (req *AddUserReq) SyncResultHdl() bool {
	return true
}

type DelUserReq struct {
	id int
	ch chan error
}

func (req *DelUserReq) OnExecute(dq *mysql.DBQuery) bool {
	if !dq.Prepare("delete from `user` where `id`=?") {
		req.ch <- errors.New("prepare failed")
		return false
	} else if !dq.Exec(req.id) {
		req.ch <- errors.New("exec failed")
		return false
	}
	req.ch <- nil
	return false
}

func (req *DelUserReq) SyncResultHdl() bool {
	return true
}

type DelUserByPfIDReq struct {
	id int
	ch chan error
}

func (req *DelUserByPfIDReq) OnExecute(dq *mysql.DBQuery) bool {
	if !dq.Prepare("delete from `user` where `pfid`=?") {
		req.ch <- errors.New("prepare failed")
		return false
	} else if !dq.Exec(req.id) {
		req.ch <- errors.New("exec failed")
		return false
	}
	req.ch <- nil
	return false
}

func (req *DelUserByPfIDReq) SyncResultHdl() bool {
	return true
}

type UpdateUserReq struct {
	id       int
	name     string
	password string
	right    int
	address  string
	status   int
	pfID     int
	ch       chan error
}

func (req *UpdateUserReq) OnExecute(dq *mysql.DBQuery) bool {
	if !dq.Prepare("update `user` set `name`=?,`password`=?,`right`=?,`address`=?,`status`=?,`pfid`=? where `id`=?") {
		req.ch <- errors.New("prepare failed")
		return false
	} else if !dq.Exec(req.name, req.password, req.right, req.address, req.status, req.pfID, req.id) {
		req.ch <- errors.New("exec failed")
		return false
	}
	req.ch <- nil
	return false
}

func (req *UpdateUserReq) SyncResultHdl() bool {
	return true
}

type UpdateUserPasswordReq struct {
	id       int
	password string
	ch       chan error
}

func (req *UpdateUserPasswordReq) OnExecute(dq *mysql.DBQuery) bool {
	if !dq.Prepare("update `user` set `password`=? where `id`=?") {
		req.ch <- errors.New("prepare failed")
		return false
	} else if !dq.Exec(req.password, req.id) {
		req.ch <- errors.New("exec failed")
		return false
	}
	req.ch <- nil
	return false
}

func (req *UpdateUserPasswordReq) SyncResultHdl() bool {
	return true
}

type UpdateUserLastLoginTimeReq struct {
	id            int
	lastLoginTime int64
	ch            chan bool
}

func (req *UpdateUserLastLoginTimeReq) OnExecute(dq *mysql.DBQuery) bool {
	if !dq.Prepare("update `user` set `lastlogintime`=? where `id`=?") {
		log.Println("prepare failed")
		req.ch <- false
		return false
	} else if !dq.Exec(req.lastLoginTime, req.id) {
		log.Println("exec failed")
		req.ch <- false
		return false
	}
	req.ch <- true
	return false
}

func (req *UpdateUserLastLoginTimeReq) SyncResultHdl() bool {
	return true
}