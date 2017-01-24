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

type User struct {
	ID            int    `json:"id"`
	PfID          int    `json:"pfid"`
	PfName        string `json:"pfname"`
	Name          string `json:"name"`
	Password      string `json:"-"`
	Right         int    `json:"right"`
	RightName     string `json:"rightname"`
	Address       string `json:address`
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

func AddUser(pfID int, name, password string, right int) (int, error) {

	if !IsPlatformExist(pfID) {
		return 0, errors.New("platform is not exist")
	}

	UserMgr.lock.Lock()
	defer UserMgr.lock.Unlock()

	for _, v := range UserMgr.users {
		if v.Name == name {
			return 0, errors.New("user is exist")
		}
	}
	var lastLoginTime int64 = 0

	ch := make(chan error)
	var id int
	mysql.AddRequest(&AddUserReq{id: &id, name: name, password: password, right: right, lastLoginTime: lastLoginTime, pfID: pfID, ch: ch})
	err := <-ch
	if err != nil {
		return 0, err
	}

	var user User
	user.ID = id
	user.Name = name
	user.Password = password
	user.Right = right
	user.LastLoginTime = lastLoginTime
	user.Status = 0
	user.PfID = pfID
	UserMgr.users[user.ID] = &user
	return user.ID, nil
}

func DelUser(id int, pfID int) error {
	UserMgr.lock.Lock()
	defer UserMgr.lock.Unlock()

	var usr *User
	var ok bool

	if usr, ok = UserMgr.users[id]; !ok {
		return errors.New("user not exist！")
	}
	if usr.PfID != pfID {
		return errors.New("pf id is wrong!")
	}

	ch := make(chan error)
	mysql.AddRequest(&DelUserReq{id: id, ch: ch})
	err := <-ch
	if err != nil {
		log.Println(err)
		return err
	}
	delete(UserMgr.users, id)
	UserSessionMgr.DelSession(id)
	return nil
}

func UpdateUser(id int, name string, password *string, right int, pfID int) error {

	UserMgr.lock.Lock()
	defer UserMgr.lock.Unlock()

	user, ok := UserMgr.users[id]
	if !ok {
		return errors.New("user not exist！")
	}

	if user.PfID != pfID {
		if !IsPlatformExist(pfID) {
			return errors.New("platform is not exist!")
		}
	}

	if user.Name != name {
		for _, v := range UserMgr.users {
			if v.Name == name {
				return errors.New("user is exist")
			}
		}
	}

	ch := make(chan error)
	if password == nil {
		password = &(user.Password)
	}

	mysql.AddRequest(&UpdateUserReq{id: id, name: name, password: *password, right: right, pfID: pfID, ch: ch})
	err := <-ch
	if err != nil {
		return err
	}

	user.Name = name
	user.Right = right
	user.PfID = pfID
	user.Password = *password

	UserSessionMgr.DelSession(id)
	return nil
}

func init() {
	UserMgr.users = make(map[int]*User)
}

