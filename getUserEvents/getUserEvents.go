package getUserEvents

import (
	"encoding/json"
	"fmt"
	"gitlab_tapd/common"
	"gitlab_tapd/logs"
	"gitlab_tapd/utils"
	"sync"
)

var (
	Wg       sync.WaitGroup
	Pids     map[int64]int
	mapMutex = sync.RWMutex{}
)

// 获取用户从开始时间到结束时间的事件数量
func GetEvents(after, before string, userId int64) {
	res := utils.Get(fmt.Sprintf("https://git.guangl.cn/api/v4/users/%d/events?after=%s&before=%s", userId, after, before))
	pids := make([]*PIds, 0)
	if err := json.Unmarshal(res, &pids); err != nil {
		logs.MyError("GET Events From Gitlab Failed", err.Error)
		panic(err.Error())
	}
	pidsMap := make(map[int64]int, 0)
	for _, v := range pids {
		pidsMap[v.ProjectId] = 0
		mapMutex.Lock()
		Pids[v.ProjectId] = 0
		mapMutex.Unlock()
	}
	for k := range pidsMap {
		common.FUser[userId].ProjectInfos[k] = &common.ProjectInfo{k, "", 0, 0}
	}
	defer func() { Wg.Done() }()
}

type PIds struct {
	ProjectId int64 `json:"project_id"`
}
