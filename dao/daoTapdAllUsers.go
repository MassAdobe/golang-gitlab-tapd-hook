package dao

import (
	"fmt"
	"gitlab_tapd/logs"
	"gitlab_tapd/utils"
)

// 插入Tapd用户
func InsertIntoUsers(users *TapdUsers) {
	if create := Ggorm.Table(users.TableName()).Create(&users); create.RowsAffected == 0 || create.Error != nil {
		logs.MyError(fmt.Sprintf("Insert Tapd User:%s failed", users.TapdUserName), create.Error)
	}
}

// 更新Tapd用户
func UpdateIntoUsers(users *TapdUsers) {
	users.UpdatedTm = utils.GetTime()
	mainId := users.TapdUserId
	users.TapdUserId = 0
	if update := Ggorm.Table(users.TableName()).Where("Tapd_user_id = ?", mainId).Update(&users); update.Error != nil || update.RowsAffected == 0 {
		logs.MyError(fmt.Sprintf("Update Tapd User:%s failed", users.TapdUserName), update.Error)
	}
}
