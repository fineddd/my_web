package user
import (
	"log"
	"sync"
	"errors"
	"encoding/json"
	"strconv"
)

const (
	UserRightGlobalScope                = 1 << 0  //全服权限
	UserRightModifyPlatform             = 1 << 1  //修改平台
	UserRightModifyServer               = 1 << 2  //修改服务器
	UserRightViewUser                   = 1 << 3  //查看用户
	UserRightModifyUser                 = 1 << 4  //修改用户
	UserRightViewRight                  = 1 << 5  //查看权限组
	UserRightModifyRight                = 1 << 6  //修改权限组
	UserRightGMCmd                      = 1 << 7  //使用ｇｍ命令
	UserRightViewGMHelp                 = 1 << 8  //查看ｇｍ帮助
	UserRightModifyGMHelp               = 1 << 9  //修改ｇｍ帮助
	UserRightPlayerBaseInfo             = 1 << 10 //玩家基础信息
	UserRightPlayerItemInfo             = 1 << 11 //玩家物品信息
	UserRightPlayerResInfo              = 1 << 12 //玩家资源信息
	UserRightViewCompensation           = 1 << 13 //查看玩家补偿
	UserRightModifyCompensation         = 1 << 14 //修改玩家补偿
	UserRightViewNotice                 = 1 << 15 //查看运营公告
	UserRightModifyNotice               = 1 << 16 //修改运营公告
	UserRightOnlineData                 = 1 << 17 //在线报表
	UserRightOutflow                    = 1 << 18 //流失率
	UserRightGradeDistributed           = 1 << 19 //等级分布
	UserRightTaskDistributed            = 1 << 20 //任务分布
	UserRightSingleServer               = 1 << 21 //单服报表
	UserRightAllServer                  = 1 << 22 //全服报表
	UserRightPlatformDistributed        = 1 << 23 //平台报表
	UserRightCDKey                      = 1 << 24 //CDK功能
	UserRightRealTimeOnline             = 1 << 25 //实时在线
	UserRightRealTimeNewPlayer          = 1 << 26 //实时新增
	UserRightRealTimePay                = 1 << 27 //实时营收
	UserRightPlayerActionPlayerGuide    = 1 << 28 //新手引导
	UserRightPlayerActionTimeDistribute = 1 << 29 //时长分布
	UserRightPlayerActionPlayerExist    = 1 << 30 //用户留存
	UserRightPlayerActionPayDistribute  = 1 << 32 //充值分布
	UserRightPlayerActionConsume        = 1 << 33 //消费分布
	UserRightGmForbidden                = 1 << 34 //封号
	UserRightGmForbiddenTalk            = 1 << 35 //禁言
	UserRightRealTimeDiamondOutput      = 1 << 36 //钻石产出
	UserRightOperatorActive             = 1 << 37 //运营活动
	UserRightPlayerActionAccumulatePay  = 1 << 38 //用户行为累计充值
	UserRightPlayerActionVipQuery       = 1 << 39 //vip查询
	UserRightPlayerActionVipCountQuery  = 1 << 40 //vip数量查询
	UserRightOrderListQuery             = 1 << 41 //订单列表查询
	UserRightOrderUserQuery             = 1 << 42 //订单用户查询
)

type RightType int

const (
	RIGHT_FULL_ADMINISTRATOR RightType = 1 << iota
	RIGHT_NORMAL_ADMINISTRATOR
	RIGHT_END
)

type Right struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Val  uint64 `json:"-"`
	Val1 uint64 `json:"val1"`
	Val2 uint64 `json:"val2"`
}

type RightManager struct {
	rights map[int]*Right
	lock   sync.Mutex
}

var RightMgr RightManager

func GetAllRight() (*[]byte, error) {
	RightMgr.lock.Lock()
	defer RightMgr.lock.Unlock()

	rights := make([]Right, len(RightMgr.rights))

	var i int = 0
	for _, v := range RightMgr.rights {
		rights[i] = *v
		rights[i].Val1 = v.Val & 0xFFFFFFFF
		rights[i].Val2 = v.Val >> 32
		i++
	}
	j, err := json.Marshal(rights)
	if err != nil {
		log.Println(err.Error())
	}
	return &j, err
}

func GetRightName(id int) (string, error) {
	RightMgr.lock.Lock()
	defer RightMgr.lock.Unlock()
	right, ok := RightMgr.rights[id]
	if !ok {
		return "N/A", errors.New("not exist right:" + strconv.Itoa(id))
	}
	return right.Name, nil
}

func GetRightValue(id int) (uint64, error) {
	RightMgr.lock.Lock()
	defer RightMgr.lock.Unlock()
	right, ok := RightMgr.rights[id]
	if !ok {
		return 0, errors.New("not exist right:" + strconv.Itoa(id))
	}
	return right.Val, nil
}

func init() {
	RightMgr.rights = make(map[int]*Right)
	(RightMgr.rights)[0] = &Right{ID: 0, Name: "FullAdminstrator", Val: uint64(RIGHT_FULL_ADMINISTRATOR)}
	(RightMgr.rights)[1] = &Right{ID: 1, Name: "Adminstrator", Val: uint64(RIGHT_NORMAL_ADMINISTRATOR)}
}