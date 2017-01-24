package user
import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"sync"
)

type Platform struct {
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Note    string          `json:"note"`
}

type PlatformManager struct {
	platforms map[int]*Platform
	lock      sync.Mutex
}
type PlatformSlice []Platform
func (p PlatformSlice) Len() int {
	return len(p)
}
func (p PlatformSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p PlatformSlice) Less(i, j int) bool {
	return p[i].ID < p[j].ID
}

var PlatformMgr PlatformManager

func GetAllPlatform() (*[]byte, error) {
	PlatformMgr.lock.Lock()
	defer PlatformMgr.lock.Unlock()

	platforms := make(PlatformSlice, len(PlatformMgr.platforms))

	var i int = 0
	for _, v := range PlatformMgr.platforms {
		platforms[i] = *v
		i++
	}

	j, err := json.Marshal(platforms)
	if err != nil {
		log.Println(err.Error())
	}
	return &j, err
}

func GetPlatformName(id int) (string, error) {
	PlatformMgr.lock.Lock()
	defer PlatformMgr.lock.Unlock()
	pf, ok := PlatformMgr.platforms[id]
	if !ok {
		return "N/A", errors.New("not exist platform:" + strconv.Itoa(id))
	}
	return pf.Name, nil
}

func init() {
	PlatformMgr.platforms = make(map[int]*Platform)
	(PlatformMgr.platforms)[0] = &Platform{ID:0, Name:"Tencent"}
	(PlatformMgr.platforms)[1] = &Platform{ID:1, Name:"Mokun"}
}