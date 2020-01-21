package getUserProjectInfo

import (
	"encoding/json"
	"fmt"
	"gitlab_tapd/common"
	"gitlab_tapd/utils"
	"strings"
	"sync"
)

var (
	Wg                  sync.WaitGroup
	ProjectInfo         map[int64]*PInfo
	mapMutex            = sync.RWMutex{}
	CommitColumnWgCount int
)

// 获取所有项目信息，并且赋值给FUser
func GetCurUserProjectName(pid int64) {
	res := utils.Get(fmt.Sprintf("https://git.guangl.cn/api/v4/projects/%d", pid))
	pinfo := new(PInfo)
	if err := json.Unmarshal(res, &pinfo); err != nil {
		panic(err.Error())
	}
	mapMutex.Lock()
	ProjectInfo[pid] = pinfo
	mapMutex.Unlock()
	defer func() { Wg.Done() }()
}

// 赋值给FUser
func BlankFUserAboutProject() {
	CommitColumnWgCount = 0
	for _, v := range common.FUser {
		if len(v.ProjectInfos) == 0 {
			continue
		}
		for k1, v1 := range v.ProjectInfos {
			if strings.ReplaceAll(ProjectInfo[k1].NameSpace.Name, " ", "") == "publish" || ProjectInfo[k1].NameSpace.ParentId == 0 {
				// || strings.ReplaceAll(ProjectInfo[k1].NameSpace.Kind, " ", "") == "group"
				delete(v.ProjectInfos, k1)
				continue
			}
			v1.ProjectName = ProjectInfo[k1].ProjectName
			CommitColumnWgCount++
		}
	}
}

type PInfo struct {
	ProjectName string     `json:"name"`
	NameSpace   pNameSpace `json:"namespace"`
}

type pNameSpace struct {
	Name     string `json:"name"`
	Kind     string `json:"kind"`
	ParentId int64  `json:"parent_id"`
}
