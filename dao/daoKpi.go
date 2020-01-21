package dao

import (
	"fmt"
	"gitlab_tapd/common"
	"gitlab_tapd/logs"
)

// 插入KPI
func InsertKPI() {
	for _, v := range common.FUser {
		for _, v1 := range v.ProjectInfos {
			kpi := &GitlabEventKpi{GitlabUserId: v.Id, ProjectId: v1.ProjectId, ProjectName: v1.ProjectName, GitlabCommitRows: v1.CommitColums, GitlabCommitTmz: v1.CommitTms}
			if create := Ggorm.Table(kpi.TableName()).Create(&kpi); create.RowsAffected == 0 || create.Error != nil {
				logs.MyError(fmt.Sprintf("Insert Gitlab KPI UserId:%d failed", kpi.GitlabUserId), create.Error)
			}
		}
	}
}
