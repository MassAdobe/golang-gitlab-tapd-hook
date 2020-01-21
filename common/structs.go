package common

var (
	FUser map[int64]*FinalUser
	TUser map[int64]*TapdUsers
)

type FinalUser struct {
	Id           int64                  `json:"id"`
	Name         string                 `json:"name"`
	Username     string                 `json:"username"`
	State        string                 `json:"state"`
	AvatarUrl    string                 `json:"avatar_url"`
	Email        string                 `json:"email"`
	ProjectInfos map[int64]*ProjectInfo `json:"project_infos"`
}

type ProjectInfo struct {
	ProjectId    int64  `json:"project_id"`
	ProjectName  string `json:"project_name"`
	CommitColums int64  `json:"commit_colums"`
	CommitTms    int64  `json:"commit_tms"`
}

const (
	BEGIN_TIME     = "T00:00:00"
	END_TIME       = "T23:59:59"
	FORMAT_TIME    = "2006-01-02T15-04-05"
	SAVE_PATH_FILE = "/Users/zhangzhen/excel/KPI-%s.xlsx"
	//LOG_FILE_NAME  = "/Users/zhangzhen/logs/golang/gitlab/git_lab.log"
	LOG_FILE_NAME = "/data/gitlab-tapd-log/gitlab_tapd.log"
)

const (
	HTTP_HEADER_ENCODE_TITLE        = "Content-Type"
	HTTP_HEADER_ENCODE_INNER        = "application/json;charset=UTF-8"
	HTTP_HEADER_PRIVATE_TOKEN_TITLE = "private-token"
	HTTP_HEADER_PRIVATE_TOKEN_INNER = ""
	HTTP_HEADER_AUTHORIZATION_TITLE = "Authorization"
	HTTP_HEADER_AUTHORIZATION_INNER = "Basic "
)

const (
	POST_USER     = "post_master"
	POST_HOST     = "192.168.2.85"
	POST_PORT     = 5432
	POST_DBNAME   = "gitlab_tapd"
	POST_PASSWORD = "post_weiyuanzhang"
)

type TapdUsers struct {
	Id          string `json:"id"`
	UserId      string `json:"user_id"`
	UserName    string `json:"user"`
	Status      string `json:"status"`
	CreatedTm   string `json:"created"`
	ModifiedTm  string `json:"modified"`
	Position    string `json:"position"`
	WorkspaceId string `json:"workspace_id"`
}

const (
	TAPD_WORKSPACE_ID = 52477016
)

type TapdProject struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PrettyName  string `json:"pretty_name"`
	Status      string `json:"status"`
	Security    string `json:"security"`
	CreatedTm   string `json:"created"`
	CreatorId   string `json:"creator_id"`
	MemberCount int64  `json:"member_count"`
	CreatorName string `json:"creator"`
}

type TapdBug struct {
	Id             string `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Priority       string `json:"priority"`
	Severity       string `json:"severity"`
	Status         string `json:"status"`
	CurrentOwner   string `json:"current_owner"`
	Reporter       string `json:"reporter"`
	Te             string `json:"te"`
	De             string `json:"de"`
	Auditer        string `json:"auditer"`
	Confimer       string `json:"confimer"`
	Closer         string `json:"closer"`
	LastModify     string `json:"lastmodify"`
	Created        string `json:"created"`
	InProgressTime string `json:"in_progress_time"`
	Resolved       string `json:"resolved"`
	VerifyTime     string `json:"verify_time"`
	Closed         string `json:"closed"`
	RejectTime     string `json:"reject_time"`
	Modified       string `json:"modified"`
	Begin          string `json:"begin"`
	Due            string `json:"due"`
	Deadline       string `json:"deadline"`
}

const (
	SPLIT_SEMICOLON_MARK = ";"
)
