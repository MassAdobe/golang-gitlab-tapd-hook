package dao

type GitlabUser struct {
	GitlabUserId       int64  `gorm:"primary_key;type:bigint;column:gitlab_user_id"`                    // 用户id（gitlab系统ID）
	GitlabUserName     string `gorm:"type:varchar(64);column:gitlab_user_name"`                         // gitlab系统用户名
	GitlabUserUsername string `gorm:"type:varchar(64);column:gitlab_user_username"`                     // gitlab系统内部用户名
	GitlabUserState    string `gorm:"type:varchar(64);column:gitlab_user_state"`                        // 系统用户的状态
	GitlabUserEmail    string `gorm:"type:varchar(256);column:gitlab_user_email"`                       // 用户的邮箱地址
	CreatedTm          string `gorm:"type:timestamp(6);column:created_tm;default:CURRENT_TIMESTAMP(6)"` // 创建时间
	UpdatedTm          string `gorm:"type:timestamp(6);column:updated_tm;default:CURRENT_TIMESTAMP(6)"` // 更新时间
}

func (*GitlabUser) TableName() string {
	return "gitlab.gitlab_users"
}

type GitlabEventKpi struct {
	KpiId            int64  `gorm:"primary_key;AUTO_INCREMENT;type:bigint;column:kpi_id"`             // 自增ID
	GitlabUserId     int64  `gorm:"type:bigint;column:gitlab_user_id"`                                // 关联gitlab_users
	ProjectId        int64  `gorm:"type:bigint;column:project_id"`                                    // 项目ID
	ProjectName      string `gorm:"type:varchar(64);column:project_name"`                             // 项目名称
	GitlabCommitRows int64  `gorm:"type:bigint;column:project_commits_rows"`                          // 提交行数
	GitlabCommitTmz   int64 `gorm:"type:bigint;column:project_commits_tmz;default:0"`                 // 提交次数
	CreatedTm        string `gorm:"type:timestamp(6);column:created_tm;default:CURRENT_TIMESTAMP(6)"` // 创建时间
	UpdatedTm        string `gorm:"type:timestamp(6);column:updated_tm;default:CURRENT_TIMESTAMP(6)"` // 更新时间
}

func (*GitlabEventKpi) TableName() string {
	return "gitlab.gitlab_event_kpi"
}
