package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gitlab_tapd/getAllUsers"
	"gitlab_tapd/logs"
	"gitlab_tapd/utils"
)

// 获取库中的所有用户
func SelectAndUpdateThenInsertAllUsers() {
	gitlabUser := make([]*GitlabUser, 0)
	if find := Ggorm.Table(new(GitlabUser).TableName()).Find(&gitlabUser); find.Error != nil && find.Error != gorm.ErrRecordNotFound {
		logs.MyError("Select Gitlab User Failed", find.Error)
		return
	}
	dbMap := make(map[int64]*GitlabUser, 0)
	for _, v := range gitlabUser {
		dbMap[v.GitlabUserId] = &GitlabUser{v.GitlabUserId, v.GitlabUserName, v.GitlabUserUsername, v.GitlabUserState, v.GitlabUserEmail, v.CreatedTm, v.UpdatedTm}
	}
	for _, v := range getAllUsers.Users {
		if _, ok := dbMap[v.Id]; ok {
			if v.Name != dbMap[v.Id].GitlabUserName || v.Username != dbMap[v.Id].GitlabUserUsername || v.Email != dbMap[v.Id].GitlabUserEmail || v.State != dbMap[v.Id].GitlabUserState {
				updateGitUsers(&GitlabUser{GitlabUserId: v.Id, GitlabUserName: v.Name, GitlabUserUsername: v.Username, GitlabUserState: v.State, GitlabUserEmail: v.Email, UpdatedTm: utils.GetTime()})
			}
		} else {
			insertGitlabUsers(&GitlabUser{GitlabUserId: v.Id, GitlabUserName: v.Name, GitlabUserUsername: v.Username, GitlabUserState: v.State, GitlabUserEmail: v.Email})
		}
	}
}

// 插入用户
func insertGitlabUsers(user *GitlabUser) {
	if create := Ggorm.Table(user.TableName()).Create(&user); create.RowsAffected == 0 || create.Error != nil {
		logs.MyError(fmt.Sprintf("Insert Gitlab User:%s failed", user.GitlabUserName), create.Error)
	}
}

// 更新用户
func updateGitUsers(user *GitlabUser) {
	user.UpdatedTm = utils.GetTime()
	mainId := user.GitlabUserId
	user.GitlabUserId = 0
	if update := Ggorm.Table(user.TableName()).Where("gitlab_user_id = ?", mainId).Update(&user); update.Error != nil || update.RowsAffected == 0 {
		logs.MyError(fmt.Sprintf("Update Gitlab User:%s failed", user.GitlabUserName), update.Error)
	}
}
