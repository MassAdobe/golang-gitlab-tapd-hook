package countUserCommitColumn

import (
	"encoding/json"
	"fmt"
	"gitlab_tapd/common"
	"gitlab_tapd/utils"
	"math/rand"
	"net/url"
	"sync"
	"time"
)

var (
	Wg      sync.WaitGroup
	mapMutx = sync.RWMutex{}
)

// 获取项目CommitID
func GainCommitIDs(userID, pid int64, name, since, until string) {
	// time.Sleep(time.Duration(RandInt64(1000, 5000)) * time.Millisecond)
	res := utils.Get(fmt.Sprintf("https://git.guangl.cn/api/v4/projects/%d/repository/commits?author_name=%s&since=%s&until=%s", pid, url.QueryEscape(name), since, until))
	cIds := make([]*CId, 0)
	if err := json.Unmarshal(res, &cIds); err != nil {
		fmt.Println(err.Error())
		fmt.Println("500 Internal Server Error", userID, pid, name)
		GainCommitIDs(userID, pid, name, since, until) // 自循环调用，可能会死循环
	}
	for _, v := range cIds {
		gainCommitColums(userID, pid, v.Id)
	}
	defer func() { Wg.Done() }()
}

// 获取项目更新行数
func gainCommitColums(userId, pid int64, id string) {
	res := utils.Get(fmt.Sprintf("https://git.guangl.cn/api/v4/projects/%d/repository/commits/%s", pid, id))
	c := new(cInfo)
	if err := json.Unmarshal(res, &c); err != nil {
		panic(err.Error())
	}
	mapMutx.Lock()
	fmt.Println("Commit COLUMNS", userId, pid, id)
	common.FUser[userId].ProjectInfos[pid].CommitColums += c.Stats.Total
	common.FUser[userId].ProjectInfos[pid].CommitTms++
	mapMutx.Unlock()
}

type CId struct {
	Id string `json:"id"`
}

type cInfo struct {
	Stats cStatus `json:"stats"`
}

type cStatus struct {
	Total int64 `json:"total"`
}

// 区间随机数
func RandInt64(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
