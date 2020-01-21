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
	TapdSysUsers map[string]*dao.TapdUsers
)

// 获取所有的用户信息
func GainAllUsers() {
	res := utils.TapdGet(fmt.Sprintf("https://api.tapd.cn/workspaces/users?workspace_id=%d&fields=users", common.TAPD_WORKSPACE_ID))
	users := new(UsersAns)
	users.Date = make([]*UserWorkspace, 0)
	if err := json.Unmarshal(res, &users); err != nil {
		logs.MyError("GET All Usersa From Tapd Failed", err.Error)
		panic(err.Error())
	}
	TapdSysUsers = make(map[string]*dao.TapdUsers, 0)
	for _, v := range users.Date {
		userId, _ := strconv.ParseInt(v.UserWs.UserId, 10, 64)
		id, _ := strconv.ParseInt(v.UserWs.Id, 10, 64)
		status, _ := strconv.ParseInt(v.UserWs.Status, 10, 64)
		workspaceId, _ := strconv.ParseInt(v.UserWs.WorkspaceId, 10, 64)
		if utils.CompareTmSame(v.UserWs.ModifiedTm) { // 如果新增日期和昨天相等，插入
			dao.InsertIntoUsers(&dao.TapdUsers{TapdUserId: userId, TapdId: id, TapdUserName: v.UserWs.UserName, TapdUserStatus: status, TapdUserPosition: v.UserWs.Position, TapdUserWorkspaceId: workspaceId, SourceCreatedTm: v.UserWs.CreatedTm, SourceUpdatedTm: v.UserWs.ModifiedTm})
		} else if utils.CompareTmSame(v.UserWs.CreatedTm) { // 如果修改日期和昨天相等，更新
			dao.UpdateIntoUsers(&dao.TapdUsers{TapdUserId: userId, TapdId: id, TapdUserName: v.UserWs.UserName, TapdUserStatus: status, TapdUserPosition: v.UserWs.Position, TapdUserWorkspaceId: workspaceId, SourceCreatedTm: v.UserWs.CreatedTm, SourceUpdatedTm: v.UserWs.ModifiedTm})
		}
		TapdSysUsers[v.UserWs.UserName] = &dao.TapdUsers{TapdUserId: userId, TapdId: id, TapdUserName: v.UserWs.UserName, TapdUserStatus: status, TapdUserPosition: v.UserWs.Position, TapdUserWorkspaceId: workspaceId, SourceCreatedTm: v.UserWs.CreatedTm, SourceUpdatedTm: v.UserWs.ModifiedTm}
	}
}

type UsersAns struct {
	Date []*UserWorkspace `json:"data"`
}

type UserWorkspace struct {
	UserWs *common.TapdUsers `json:"UserWorkspace"`
}
