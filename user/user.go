package user

import (
	"common/mysql"
	"errors"
	"log"
	"sync"
	"time"
	"encoding/json"
)

const (
	UserStatusNormal    = 0
	UserStatusForbidden = 1
)

type RightType int

const (
	RIGHT_FULL_ADMINISTRATOR RightType = iota
	RIGHT_NORMAL_ADMINISTRATOR
	RIGHT_END
)

type User struct {
	ID            int    `json:"id"`
	PfID          int    `json:"pfid"`
	PfName        string `json:"pfname"`
	Name          string `json:"name"`
	Password      string `json:"-"`
	Right         int    `json:"right"`
	RightName     string `json:"rightname"`
	Address       string `json:"address"`
	Status        int    `json:"status"`
	LastLoginTime int64  `json:"lastlogintime"`
}

type UserManager struct {
	users map[int]*User
	lock  sync.Mutex
}

var UserMgr UserManager

func LoadUser() bool {
	ch := make(chan error)
	mysql.AddRequest(&LoadUserReq{users: &UserMgr.users, ch: ch})
	err := <-ch
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func UserLogin(name, password string) (string, User, error) {
	UserMgr.lock.Lock()
	defer UserMgr.lock.Unlock()

	var user User

	for _, v := range UserMgr.users {
		if v.Name == name {
			if v.Password != password {
				return "", user, errors.New("password wrong")
			} else {
				sessionID, err := UserSessionMgr.addSession(v.ID)
				if err != nil {
					return "", user, err
				}
				ch := make(chan bool)
				now := time.Now().Unix()
				mysql.AddRequest(&UpdateUserLastLoginTimeReq{id: v.ID, lastLoginTime: now, ch: ch})
				ret := <-ch
				if !ret {
					return "", user, errors.New("update login time failed")
				}
				v.LastLoginTime = now
				user = *v
				return sessionID, user, nil
			}
		}
	}
	return "", user, errors.New("username wrong")
}

func GetAllUser() (*[]byte, error) {
	UserMgr.lock.Lock()
	defer UserMgr.lock.Unlock()

	users := make([]User, len(UserMgr.users))

	var i int = 0
	for _, v := range UserMgr.users {
		users[i] = *v
		i++
	}

	j, err := json.Marshal(users)
	if err != nil {
		log.Println(err.Error())
	}
	return &j, err
}

func init() {
	UserMgr.users = make(map[int]*User)
}

