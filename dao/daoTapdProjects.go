package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gitlab_tapd/logs"
	"gitlab_tapd/utils"
)

// 查询库里所有的项目
func SelectAllProject(fProjects *[]*TapdProjects) {
	projects := make([]*TapdProjects, 0)
	if find := Ggorm.Table(new(TapdProjects).TableName()).Find(&projects); find.Error != nil && find.Error != gorm.ErrRecordNotFound {
		logs.MyError("Select Tapd Projects Failed", find.Error)
		return
	}
	dbMap := make(map[int64]*TapdProjects, 0)
	for _, v := range projects {
		dbMap[v.ProjectId] = v
	}
	for _, v := range *fProjects {
		if _, ok := dbMap[v.ProjectId]; ok {
			if v.ProjectStatus != dbMap[v.ProjectId].ProjectStatus || v.ProjectMemberCount != dbMap[v.ProjectId].ProjectMemberCount || v.ProjectName != dbMap[v.ProjectId].ProjectName {
				updateIntoProjects(v)
			}
		} else {
			insertIntoProjects(v)
		}
	}
}

// 插入Tapd用户
func insertIntoProjects(project *TapdProjects) {
	if create := Ggorm.Table(project.TableName()).Create(&project); create.RowsAffected == 0 || create.Error != nil {
		logs.MyError(fmt.Sprintf("Insert Tapd Project:%s failed", project.ProjectName), create.Error)
	}
}

// 更新Tapd用户
func updateIntoProjects(project *TapdProjects) {
	project.UpdatedTm = utils.GetTime()
	mainId := project.ProjectId
	project.ProjectId = 0
	if update := Ggorm.Table(project.TableName()).Where("project_id = ?", mainId).Update(&project); update.Error != nil || update.RowsAffected == 0 {
		logs.MyError(fmt.Sprintf("Update Tapd Project:%s failed", project.ProjectName), update.Error)
	}
}
