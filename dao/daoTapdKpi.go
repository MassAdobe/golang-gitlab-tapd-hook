package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gitlab_tapd/logs"
	"gitlab_tapd/utils"
)

// 插入Rela表
func InsertIntoRela(bugId, projectId, relaType, tapdUserId int64) {
	newRela := &TapdBugRela{BugId: bugId, ProjectId: projectId, RelaType: relaType, TapdUsersId: tapdUserId, IsDeleted: 1}
	if create := Ggorm.Table(newRela.TableName()).Create(&newRela); create.RowsAffected == 0 || create.Error != nil {
		logs.MyError(fmt.Sprintf("Insert Tapd Kpi Rela:%d failed", bugId), create.Error)
	}
	logs.MyInfo(fmt.Sprintf("TAPD RELA bugID: %d; relaType: %d; tapdUserId: %d", bugId, relaType, tapdUserId))
	//fmt.Println("---TAPD RELA--- ", bugId, "  ", relaType, "  ", tapdUserId)
}

// 插入Kpi表
func InsertIntoKpi(kpi *TapdBugKpi) {
	if create := Ggorm.Table(kpi.TableName()).Create(&kpi); create.RowsAffected == 0 || create.Error != nil {
		logs.MyError(fmt.Sprintf("Insert Tapd Kpi:%d failed", kpi.BugId), create.Error)
	}
	//fmt.Println("---TAPD KPI---", kpi.BugId, "  ", kpi.BugTitle)
	logs.MyInfo(fmt.Sprintf("TAPD KPI bugID: %d; bugTitle: %s", kpi.BugId, kpi.BugTitle))
}

// 更新KPI表
func UpdateIntoKpi(kpi *TapdBugKpi) {
	kpi.UpdatedTm = utils.GetTime()
	mainId := kpi.BugId
	kpi.BugId = 0
	if update := Ggorm.Table(kpi.TableName()).Where("bug_id = ?", mainId).Update(&kpi); update.Error != nil || update.RowsAffected == 0 {
		logs.MyError(fmt.Sprintf("Update Tapd Kpi:%d failed", kpi.BugId), update.Error)
	}
	logs.MyInfo(fmt.Sprintf("TAPD KPI UPDATE bugId: %d", kpi.BugId))
}

// 进入后应该是存在的情况，所以先删除所有人，在下面的方法中恢复
func DeleteAllRela(bugId, projectId, relaType int64) {
	rela := &TapdBugRela{IsDeleted: 2, UpdatedTm: utils.GetTime()}
	if update := Ggorm.Table(rela.TableName()).Where("bug_id = ? and project_id = ? and rela_type = ?", bugId, projectId, relaType).Update(&rela); update.Error != nil {
		logs.MyError(fmt.Sprintf("Select Then Update Tapd Kpi:%d failed", rela.BugId), update.Error)
	}
}

// 查询Rela表中是否存在，存在则不管，不存在则插入，没有了则删除
func SelectThenDMLRela(bugId, projectId, relaType, tapdUserId int64) {
	rela := &TapdBugRela{BugId: bugId, ProjectId: projectId, RelaType: relaType, TapdUsersId: tapdUserId, IsDeleted: 1}
	relaFind := new(TapdBugRela)
	if find := Ggorm.Table(relaFind.TableName()).Where("bug_id = ? and project_id = ? and rela_type = ? and tapd_users_id = ?", bugId, projectId, relaType, tapdUserId); find.Error != nil && find.Error != gorm.ErrRecordNotFound {
		if find.Error != gorm.ErrRecordNotFound || find.RowsAffected == 0 { // 不存在 插入
			if create := Ggorm.Table(rela.TableName()).Create(&rela); create.RowsAffected == 0 || create.Error != nil {
				logs.MyError(fmt.Sprintf("Select Then Insert Tapd Kpi Rela:%d failed", bugId), create.Error)
			}
			return
		}
		logs.MyError("Select Tapd Kpi Failed", find.Error)
		return
	} else { // 已经存在了 直接更新
		if update := Ggorm.Table(rela.TableName()).Where("bug_rela_id = ?", relaFind.BugRelaId).Update(&rela); update.Error != nil || update.RowsAffected == 0 {
			if update.RowsAffected == 0 {
				logs.MyInfo(fmt.Sprintf("Select Then Update Tapd Has No Rows To Updated. BugId: %d", bugId))
			} else {
				logs.MyError(fmt.Sprintf("Select Then Update Tapd Kpi(BugId):%d failed", rela.BugId), update.Error)
			}
		}
		logs.MyInfo(fmt.Sprintf("TAPD KPI SELECT UPDATE. BugId: %d", bugId))
	}
}
