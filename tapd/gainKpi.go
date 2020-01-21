package tapd

import (
	"encoding/json"
	"fmt"
	"gitlab_tapd/common"
	"gitlab_tapd/dao"
	"gitlab_tapd/logs"
	"gitlab_tapd/utils"
	"strconv"
	"strings"
	"time"
)

func GainTapdKpiInsert(project *dao.TapdProjects, selectTime string) {
	// 先增加
	for i := 1; i < 1000; i++ {
		res := utils.TapdGet(fmt.Sprintf("https://api.tapd.cn/bugs?workspace_id=%d&limit=200&page=%d&created=%s", project.ProjectId, i, selectTime))
		kpis := new(KpiAns)
		kpis.Data = make([]*Bug, 0)
		time.Sleep(1 * time.Second)
		if err := json.Unmarshal(res, &kpis); err != nil {
			fmt.Println(string(res))
			logs.MyError("GET All Kpi From Tapd Failed", err.Error)
			panic(err.Error())
		}
		if len(kpis.Data) == 0 {
			break
		}
		for _, v := range kpis.Data {
			bugId, _ := strconv.ParseInt(v.BugKpi.Id, 10, 64)
			currentOwner := strings.Split(v.BugKpi.CurrentOwner, common.SPLIT_SEMICOLON_MARK)
			// 当前处理人
			if len(currentOwner) != 0 {
				for _, vc := range currentOwner {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 1, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 创建人
			reporter := strings.Split(v.BugKpi.Reporter, common.SPLIT_SEMICOLON_MARK)
			if len(reporter) != 0 {
				for _, vc := range reporter {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 2, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 测试人员
			te := strings.Split(v.BugKpi.Te, common.SPLIT_SEMICOLON_MARK)
			if len(te) != 0 {
				for _, vc := range te {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 3, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 开发人员
			de := strings.Split(v.BugKpi.De, common.SPLIT_SEMICOLON_MARK)
			if len(te) != 0 {
				for _, vc := range de {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 4, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 审核人员
			auditer := strings.Split(v.BugKpi.Auditer, common.SPLIT_SEMICOLON_MARK)
			if len(auditer) != 0 {
				for _, vc := range auditer {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 5, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 验证人员
			confirmer := strings.Split(v.BugKpi.Confimer, common.SPLIT_SEMICOLON_MARK)
			if len(confirmer) != 0 {
				for _, vc := range confirmer {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 6, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 关闭人员
			closer := strings.Split(v.BugKpi.Closer, common.SPLIT_SEMICOLON_MARK)
			if len(closer) != 0 {
				for _, vc := range closer {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 7, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			insert := new(dao.TapdBugKpi)
			insert.BugId = bugId
			if len(v.BugKpi.Title) != 0 {
				insert.BugTitle = v.BugKpi.Title
			}
			if len(v.BugKpi.Priority) != 0 {
				insert.BugPriority = v.BugKpi.Priority
			}
			if len(v.BugKpi.Severity) != 0 {
				insert.BugSeverity = v.BugKpi.Severity
			}
			if len(v.BugKpi.Status) != 0 {
				insert.BugStatus = v.BugKpi.Status
			}
			if _, ok := TapdSysUsers[v.BugKpi.LastModify]; ok {
				insert.BugLastModifiedId = TapdSysUsers[v.BugKpi.LastModify].TapdUserId
			}
			if len(v.BugKpi.LastModify) != 0 {
				insert.BugLastModifiedName = v.BugKpi.LastModify
			}
			if len(v.BugKpi.Created) != 0 {
				insert.SourceCreatedTm = v.BugKpi.Created
			}
			if len(v.BugKpi.InProgressTime) != 0 {
				insert.SourceInProgressTm = v.BugKpi.InProgressTime
			}
			if len(v.BugKpi.Resolved) != 0 {
				insert.SourceResolvedTm = v.BugKpi.Resolved
			}
			if len(v.BugKpi.VerifyTime) != 0 {
				insert.SourceVerifyTm = v.BugKpi.VerifyTime
			}
			if len(v.BugKpi.Closed) != 0 {
				insert.SourceClosedTm = v.BugKpi.Closed
			}
			if len(v.BugKpi.RejectTime) != 0 {
				insert.SourceRejectTm = v.BugKpi.RejectTime
			}
			if len(v.BugKpi.Modified) != 0 {
				insert.SourceModifiedTm = v.BugKpi.Modified
			}
			if len(v.BugKpi.Begin) != 0 {
				insert.SourceBeginTm = v.BugKpi.Begin
			}
			if len(v.BugKpi.Due) != 0 {
				insert.SourceDueTm = v.BugKpi.Due
			}
			if len(v.BugKpi.Deadline) != 0 {
				insert.SourceDeadlineTm = v.BugKpi.Deadline
			}
			if len(v.BugKpi.Description) != 0 {
				insert.BugDescription = v.BugKpi.Description
			}
			insert.ProjectId = project.ProjectId
			dao.InsertIntoKpi(insert)
		}
	}
}

func GainTapdKpiUpdate(project *dao.TapdProjects, selectTime string) {
	// 再更新
	for i := 1; i < 1000; i++ {
		res := utils.TapdGet(fmt.Sprintf("https://api.tapd.cn/bugs?workspace_id=%d&limit=200&page=%d&modified=%s", project.ProjectId, i, selectTime))
		kpis := new(KpiAns)
		kpis.Data = make([]*Bug, 0)
		time.Sleep(1 * time.Second)
		if err := json.Unmarshal(res, &kpis); err != nil {
			fmt.Println(string(res))
			logs.MyError("GET All Kpi From Tapd Failed", err.Error)
			panic(err.Error())
		}
		if len(kpis.Data) == 0 {
			break
		}
		for _, v := range kpis.Data {
			if v.BugKpi.Created == v.BugKpi.Modified {
				continue
			}
			bugId, _ := strconv.ParseInt(v.BugKpi.Id, 10, 64)
			currentOwner := strings.Split(v.BugKpi.CurrentOwner, common.SPLIT_SEMICOLON_MARK)
			// 当前处理人
			if len(currentOwner) != 0 {
				dao.DeleteAllRela(bugId, project.ProjectId, 1)
				for _, vc := range currentOwner {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.SelectThenDMLRela(bugId, project.ProjectId, 1, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 创建人
			reporter := strings.Split(v.BugKpi.Reporter, common.SPLIT_SEMICOLON_MARK)
			if len(reporter) != 0 {
				dao.DeleteAllRela(bugId, project.ProjectId, 2)
				for _, vc := range reporter {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.SelectThenDMLRela(bugId, project.ProjectId, 2, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// TODO 测试人员
			te := strings.Split(v.BugKpi.Te, common.SPLIT_SEMICOLON_MARK)
			if len(te) != 0 {
				dao.DeleteAllRela(bugId, project.ProjectId, 3)
				for _, vc := range te {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.SelectThenDMLRela(bugId, project.ProjectId, 3, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 开发人员
			de := strings.Split(v.BugKpi.De, common.SPLIT_SEMICOLON_MARK)
			if len(te) != 0 {
				dao.DeleteAllRela(bugId, project.ProjectId, 4)
				for _, vc := range de {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.SelectThenDMLRela(bugId, project.ProjectId, 4, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 审核人员
			auditer := strings.Split(v.BugKpi.Auditer, common.SPLIT_SEMICOLON_MARK)
			if len(auditer) != 0 {
				dao.DeleteAllRela(bugId, project.ProjectId, 5)
				for _, vc := range auditer {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.SelectThenDMLRela(bugId, project.ProjectId, 5, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 验证人员
			confirmer := strings.Split(v.BugKpi.Confimer, common.SPLIT_SEMICOLON_MARK)
			if len(confirmer) != 0 {
				dao.DeleteAllRela(bugId, project.ProjectId, 6)
				for _, vc := range confirmer {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.SelectThenDMLRela(bugId, project.ProjectId, 6, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 关闭人员
			closer := strings.Split(v.BugKpi.Closer, common.SPLIT_SEMICOLON_MARK)
			if len(closer) != 0 {
				dao.DeleteAllRela(bugId, project.ProjectId, 7)
				for _, vc := range closer {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.SelectThenDMLRela(bugId, project.ProjectId, 7, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			update := new(dao.TapdBugKpi)
			update.BugId = bugId
			if len(v.BugKpi.Title) != 0 {
				update.BugTitle = v.BugKpi.Title
			}
			if len(v.BugKpi.Priority) != 0 {
				update.BugPriority = v.BugKpi.Priority
			}
			if len(v.BugKpi.Severity) != 0 {
				update.BugSeverity = v.BugKpi.Severity
			}
			if len(v.BugKpi.Status) != 0 {
				update.BugStatus = v.BugKpi.Status
			}
			if _, ok := TapdSysUsers[v.BugKpi.LastModify]; ok {
				update.BugLastModifiedId = TapdSysUsers[v.BugKpi.LastModify].TapdUserId
			}
			if len(v.BugKpi.LastModify) != 0 {
				update.BugLastModifiedName = v.BugKpi.LastModify
			}
			if len(v.BugKpi.Created) != 0 {
				update.SourceCreatedTm = v.BugKpi.Created
			}
			if len(v.BugKpi.InProgressTime) != 0 {
				update.SourceInProgressTm = v.BugKpi.InProgressTime
			}
			if len(v.BugKpi.Resolved) != 0 {
				update.SourceResolvedTm = v.BugKpi.Resolved
			}
			if len(v.BugKpi.VerifyTime) != 0 {
				update.SourceVerifyTm = v.BugKpi.VerifyTime
			}
			if len(v.BugKpi.Closed) != 0 {
				update.SourceClosedTm = v.BugKpi.Closed
			}
			if len(v.BugKpi.RejectTime) != 0 {
				update.SourceRejectTm = v.BugKpi.RejectTime
			}
			if len(v.BugKpi.Modified) != 0 {
				update.SourceModifiedTm = v.BugKpi.Modified
			}
			if len(v.BugKpi.Begin) != 0 {
				update.SourceBeginTm = v.BugKpi.Begin
			}
			if len(v.BugKpi.Due) != 0 {
				update.SourceDueTm = v.BugKpi.Due
			}
			if len(v.BugKpi.Deadline) != 0 {
				update.SourceDeadlineTm = v.BugKpi.Deadline
			}
			if len(v.BugKpi.Description) != 0 {
				update.BugDescription = v.BugKpi.Description
			}
			update.ProjectId = project.ProjectId
			dao.UpdateIntoKpi(update)
		}
	}
}

// 第一次批量使用
func InsertTapdKpi(project *dao.TapdProjects) {
	for i := 1; i < 1000; i++ {
		res := utils.TapdGet(fmt.Sprintf("https://api.tapd.cn/bugs?workspace_id=%d&limit=200&page=%d", project.ProjectId, i))
		kpis := new(KpiAns)
		kpis.Data = make([]*Bug, 0)
		time.Sleep(1 * time.Second)
		if err := json.Unmarshal(res, &kpis); err != nil {
			fmt.Println(string(res))
			logs.MyError("GET All Kpi From Tapd Failed", err.Error)
			panic(err.Error())
		}
		if len(kpis.Data) == 0 {
			break
		}
		for _, v := range kpis.Data {
			bugId, _ := strconv.ParseInt(v.BugKpi.Id, 10, 64)
			currentOwner := strings.Split(v.BugKpi.CurrentOwner, common.SPLIT_SEMICOLON_MARK)
			// 当前处理人
			if len(currentOwner) != 0 {
				for _, vc := range currentOwner {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 1, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 创建人
			reporter := strings.Split(v.BugKpi.Reporter, common.SPLIT_SEMICOLON_MARK)
			if len(reporter) != 0 {
				for _, vc := range reporter {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 2, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 测试人员
			te := strings.Split(v.BugKpi.Te, common.SPLIT_SEMICOLON_MARK)
			if len(te) != 0 {
				for _, vc := range te {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 3, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 开发人员
			de := strings.Split(v.BugKpi.De, common.SPLIT_SEMICOLON_MARK)
			if len(te) != 0 {
				for _, vc := range de {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 4, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 审核人员
			auditer := strings.Split(v.BugKpi.Auditer, common.SPLIT_SEMICOLON_MARK)
			if len(auditer) != 0 {
				for _, vc := range auditer {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 5, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 验证人员
			confirmer := strings.Split(v.BugKpi.Confimer, common.SPLIT_SEMICOLON_MARK)
			if len(confirmer) != 0 {
				for _, vc := range confirmer {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 6, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			// 关闭人员
			closer := strings.Split(v.BugKpi.Closer, common.SPLIT_SEMICOLON_MARK)
			if len(closer) != 0 {
				for _, vc := range closer {
					if _, ok := TapdSysUsers[vc]; ok {
						dao.InsertIntoRela(bugId, project.ProjectId, 7, TapdSysUsers[vc].TapdUserId)
					}
				}
			}
			insert := new(dao.TapdBugKpi)
			insert.BugId = bugId
			if len(v.BugKpi.Title) != 0 {
				insert.BugTitle = v.BugKpi.Title
			}
			if len(v.BugKpi.Priority) != 0 {
				insert.BugPriority = v.BugKpi.Priority
			}
			if len(v.BugKpi.Severity) != 0 {
				insert.BugSeverity = v.BugKpi.Severity
			}
			if len(v.BugKpi.Status) != 0 {
				insert.BugStatus = v.BugKpi.Status
			}
			if _, ok := TapdSysUsers[v.BugKpi.LastModify]; ok {
				insert.BugLastModifiedId = TapdSysUsers[v.BugKpi.LastModify].TapdUserId
			}
			if len(v.BugKpi.LastModify) != 0 {
				insert.BugLastModifiedName = v.BugKpi.LastModify
			}
			if len(v.BugKpi.Created) != 0 {
				insert.SourceCreatedTm = v.BugKpi.Created
			}
			if len(v.BugKpi.InProgressTime) != 0 {
				insert.SourceInProgressTm = v.BugKpi.InProgressTime
			}
			if len(v.BugKpi.Resolved) != 0 {
				insert.SourceResolvedTm = v.BugKpi.Resolved
			}
			if len(v.BugKpi.VerifyTime) != 0 {
				insert.SourceVerifyTm = v.BugKpi.VerifyTime
			}
			if len(v.BugKpi.Closed) != 0 {
				insert.SourceClosedTm = v.BugKpi.Closed
			}
			if len(v.BugKpi.RejectTime) != 0 {
				insert.SourceRejectTm = v.BugKpi.RejectTime
			}
			if len(v.BugKpi.Modified) != 0 {
				insert.SourceModifiedTm = v.BugKpi.Modified
			}
			if len(v.BugKpi.Begin) != 0 {
				insert.SourceBeginTm = v.BugKpi.Begin
			}
			if len(v.BugKpi.Due) != 0 {
				insert.SourceDueTm = v.BugKpi.Due
			}
			if len(v.BugKpi.Deadline) != 0 {
				insert.SourceDeadlineTm = v.BugKpi.Deadline
			}
			if len(v.BugKpi.Description) != 0 {
				insert.BugDescription = v.BugKpi.Description
			}
			insert.ProjectId = project.ProjectId
			dao.InsertIntoKpi(insert)
		}
	}
}

type KpiAns struct {
	Data []*Bug `json:"data"`
}

type Bug struct {
	BugKpi *common.TapdBug `json:"Bug"`
}
