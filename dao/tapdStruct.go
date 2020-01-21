package dao

type TapdUsers struct {
	TapdUserId          int64  `gorm:"primary_key;type:bigint;column:tapd_user_id"`                                                           // 系统主键ID
	TapdId              int64  `gorm:"type:bigint;column:tapd_id"`                                                                            // 目前不知道什么作用
	TapdUserName        string `gorm:"type:varchar(128);column:tapd_user_name"`                                                               // tapd用户名
	TapdUserStatus      int64  `gorm:"type:smallint;default:-1;column:tapd_user_status"`                                                      // tapd用户状态
	TapdUserPosition    string `gorm:"type:varchar(128);column:tapd_user_position"`                                                           // tapd用户职位
	TapdUserWorkspaceId int64  `gorm:"type:bigint;column:tapd_user_workspace_id"`                                                             // tapd用户所属单位ID
	SourceCreatedTm     string `gorm:"type:timestamp(6);default:make_timestamp(1970,1,1,0,0,(0)::double precision);column:source_created_tm"` // 原系统中的创建时间
	SourceUpdatedTm     string `gorm:"type:timestamp(6);default:make_timestamp(1970,1,1,0,0,(0)::double precision);column:source_updated_tm"` // 原系统中的修改时间
	CreatedTm           string `gorm:"type:timestamp(6);column:created_tm;default:CURRENT_TIMESTAMP(6)"`                                      // 创建时间
	UpdatedTm           string `gorm:"type:timestamp(6);column:updated_tm;default:CURRENT_TIMESTAMP(6)"`                                      // 更新时间
}

func (*TapdUsers) TableName() string {
	return "tapd.tapd_users"
}

type TapdProjects struct {
	ProjectId          int64  `gorm:"primary_key;type:bigint;column:project_id"`                                                              // 项目ID
	ProjectName        string `gorm:"type:varchar(256);column:project_name"`                                                                  // 项目名称
	ProjectStatus      string `gorm:"type:varchar(64);column:project_status"`                                                                 // 项目状态
	ProjectCreatedTm   string `gorm:"type:timestamp(6);default:make_timestamp(1970,1,1,0,0,(0)::double precision);column:project_created_tm"` // 项目创建时间
	ProjectCreatorId   int64  `gorm:"type:smallint;column:project_creator_id"`                                                                // 创建项目的人员ID
	ProjectCreator     string `gorm:"type:varchar(256);column:project_creator"`                                                               // 创建项目的人员名称和Email地址
	ProjectMemberCount int64  `gorm:"type:bigint;default:0;column:project_member_count"`                                                      // 参与项目的人数
	CreatedTm          string `gorm:"type:timestamp(6);column:created_tm;default:CURRENT_TIMESTAMP(6)"`                                       // 创建时间
	UpdatedTm          string `gorm:"type:timestamp(6);column:updated_tm;default:CURRENT_TIMESTAMP(6)"`                                       // 更新时间
}

func (*TapdProjects) TableName() string {
	return "tapd.tapd_projects"
}

type TapdBugKpi struct {
	BugId               int64  `gorm:"primary_key;type:bigint;column:bug_id"`                                                                     // 系统缺陷id
	BugTitle            string `gorm:"type:varchar(256);column:bug_title"`                                                                        // 缺陷标题
	BugPriority         string `gorm:"type:varchar(64);column:bug_priority"`                                                                      // 缺陷优先级
	BugSeverity         string `gorm:"type:varchar(64);column:bug_severity"`                                                                      // 缺陷严重程度
	BugStatus           string `gorm:"type:varchar(64);column:bug_status"`                                                                        // 缺陷状态
	BugLastModifiedId   int64  `gorm:"type:bigint;default:-1;column:bug_last_modified_id"`                                                        // 最后修改人ID
	BugLastModifiedName string `gorm:"type:varchar(64);column:bug_last_modified_name"`                                                            // 最后修改人username
	SourceCreatedTm     string `gorm:"type:timestamp(6);default:make_timestamp(1970,1,1,0,0,(0)::double precision);column:source_created_tm"`     // 原系统创建时间
	SourceInProgressTm  string `gorm:"type:timestamp(6);default:make_timestamp(1970,1,1,0,0,(0)::double precision);column:source_in_progress_tm"` // 原系统接收处理时间
	SourceResolvedTm    string `gorm:"type:timestamp(6);default:make_timestamp(1970,1,1,0,0,(0)::double precision);column:source_resolved_tm"`    // 原系统解决时间
	SourceVerifyTm      string `gorm:"type:timestamp(6);default:make_timestamp(1970,1,1,0,0,(0)::double precision);column:source_verify_tm"`      // 原系统验证时间
	SourceClosedTm      string `gorm:"type:timestamp(6);default:make_timestamp(1970,1,1,0,0,(0)::double precision);column:source_closed_tm"`      // 原系统关闭时间
	SourceRejectTm      string `gorm:"type:timestamp(6);default:make_timestamp(1970,1,1,0,0,(0)::double precision);column:source_reject_tm"`      // 原系统拒绝时间
	SourceModifiedTm    string `gorm:"type:timestamp(6);default:make_timestamp(1970,1,1,0,0,(0)::double precision);column:source_modified_tm"`    // 原系统最后修改时间
	SourceBeginTm       string `gorm:"type:timestamp(6);default:make_timestamp(1970,1,1,0,0,(0)::double precision);column:source_begin_tm"`       // 原系统预计开始时间
	SourceDueTm         string `gorm:"type:timestamp(6);default:make_timestamp(1970,1,1,0,0,(0)::double precision);column:source_due_tm"`         // 原系统预计结束时间
	SourceDeadlineTm    string `gorm:"type:timestamp(6);default:make_timestamp(1970,1,1,0,0,(0)::double precision);column:source_deadline_tm"`    // 原系统解决期限时间
	ProjectId           int64  `gorm:"type:bigint;column:project_id"`                                                                             // 关联项目ID
	BugDescription      string `gorm:"type:text;column:bug_description"`                                                                          // 缺陷详细描述
	CreatedTm           string `gorm:"type:timestamp(6);column:created_tm;default:CURRENT_TIMESTAMP(6)"`                                          // 创建时间
	UpdatedTm           string `gorm:"type:timestamp(6);column:updated_tm;default:CURRENT_TIMESTAMP(6)"`                                          // 更新时间
}

func (*TapdBugKpi) TableName() string {
	return "tapd.tapd_bug_kpi"
}

type TapdBugRela struct {
	BugRelaId   int64  `gorm:"primary_key;AUTO_INCREMENT;type:bigint;column:bug_rela_id"`        // rela主键(自增)
	BugId       int64  `gorm:"type:bigint;column:bug_id"`                                        // 关联表tapd_bug_kpi的主键
	ProjectId   int64  `gorm:"type:bigint;column:project_id"`                                    // 关联表tapd_projects的主键
	RelaType    int64  `gorm:"type:integer;column:rela_type"`                                    // 关联类型：1.当前处理人；2.创建人；3.测试人员；4.开发人员；5.审核人员；6.验证人员；7.关闭人员；
	TapdUsersId int64  `gorm:"type:bigint;column:tapd_users_id"`                                 // 关联表tapd_users的主键
	IsDeleted   int64  `gorm:"type:smallint;column:is_deleted"`                                  // 是否删除：1.未删除；2.已删除
	CreatedTm   string `gorm:"type:timestamp(6);column:created_tm;default:CURRENT_TIMESTAMP(6)"` // 创建时间
	UpdatedTm   string `gorm:"type:timestamp(6);column:updated_tm;default:CURRENT_TIMESTAMP(6)"` // 更新时间
}

func (*TapdBugRela) TableName() string {
	return "tapd.tapd_bug_rela"
}
