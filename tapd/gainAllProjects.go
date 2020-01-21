package tapd

import (
	"encoding/json"
	"fmt"
	"gitlab_tapd/common"
	"gitlab_tapd/dao"
	"gitlab_tapd/logs"
	"gitlab_tapd/utils"
	"strconv"
)

var (
	TapdSysProjects map[int64]*dao.TapdProjects
)

func GainAllProjects() {
	res := utils.TapdGet(fmt.Sprintf("https://api.tapd.cn/workspaces/projects?company_id=%d", common.TAPD_WORKSPACE_ID))
	projects := new(ProjectAns)
	projects.Data = make([]*Workspace, 0)
	if err := json.Unmarshal(res, &projects); err != nil {
		logs.MyError("GET All Projects From Tapd Failed", err.Error)
		panic(err.Error)
	}
	frontList := make([]*dao.TapdProjects, 0)
	TapdSysProjects = make(map[int64]*dao.TapdProjects, 0)
	for _, v := range projects.Data {
		id, _ := strconv.ParseInt(v.Projects.Id, 10, 64)
		creator, _ := strconv.ParseInt(v.Projects.CreatorId, 10, 64)
		frontList = append(frontList, &dao.TapdProjects{ProjectId: id, ProjectName: v.Projects.Name, ProjectStatus: v.Projects.Status, ProjectCreatedTm: v.Projects.CreatedTm, ProjectCreatorId: creator, ProjectCreator: v.Projects.CreatorName, ProjectMemberCount: v.Projects.MemberCount})
		TapdSysProjects[id] = &dao.TapdProjects{ProjectId: id, ProjectName: v.Projects.Name, ProjectStatus: v.Projects.Status, ProjectCreatedTm: v.Projects.CreatedTm, ProjectCreatorId: creator, ProjectCreator: v.Projects.CreatorName, ProjectMemberCount: v.Projects.MemberCount}
	}
	dao.SelectAllProject(&frontList)
}

type ProjectAns struct {
	Data []*Workspace `json:"data"`
}

type Workspace struct {
	Projects *common.TapdProject `json:"Workspace"`
}
