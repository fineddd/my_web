package user

import (
	"common/session"
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type UserSessionManager struct {
	sessionMgr *session.SessionManager
	users      map[int]string
	lock       sync.Mutex
}

var UserSessionMgr UserSessionManager

func (mgr *UserSessionManager) addSession(userID int) (string, error) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	if sessionID, ok := mgr.users[userID]; ok {
		mgr.sessionMgr.DelSession(sessionID)
	}

	now := time.Now().Unix()
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	sessionID := strconv.FormatInt(now, 10) + strconv.Itoa(randSource.Intn(1000)) + strconv.Itoa(userID)

	if err := mgr.sessionMgr.AddSession(sessionID); err != nil {
		return "", err
	}
	mgr.users[userID] = sessionID
	return sessionID, nil
}

func (mgr *UserSessionManager) DelSession(userID int) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	if sessionID, ok := mgr.users[userID]; ok {
		mgr.sessionMgr.DelSession(sessionID)
		delete(mgr.users, userID)
	}
}

func (mgr *UserSessionManager) GetValue(r *http.Request, key interface{}) interface{} {
	cookie, err := r.Cookie("usersession")
	if err == nil {
		mgr.lock.Lock()
		defer mgr.lock.Unlock()
		return mgr.sessionMgr.GetValue(cookie.Value, key)
	}
	return nil
}

func (mgr *UserSessionManager) SetValue(r *http.Request, key, value interface{}) error {
	if cookie, err := r.Cookie("usersession"); err == nil {
		mgr.lock.Lock()
		defer mgr.lock.Unlock()
		return mgr.sessionMgr.SetValue(cookie.Value, key, value)
	}
	return errors.New("can't find usersession")
}

func (mgr *UserSessionManager) SetValueByID(id string, key, value interface{}) error {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	return mgr.sessionMgr.SetValue(id, key, value)
}

func (mgr *UserSessionManager) DelValue(r *http.Request, key interface{}) {
	if cookie, err := r.Cookie("usersession"); err == nil {
		mgr.lock.Lock()
		defer mgr.lock.Unlock()
		mgr.sessionMgr.DelValue(cookie.Value, key)
	}
}

func init() {
	UserSessionMgr.sessionMgr = session.NewSessionManager(7200)
	UserSessionMgr.users = make(map[int]string)
}