package getAllUsers

import (
	"encoding/json"
	"fmt"
	"gitlab_tapd/logs"
	"gitlab_tapd/utils"
)

var (
	Users []*User
	//Wg    sync.WaitGroup
)

// 获取所有用户信息
func PerUsers(page int) bool {
	res := utils.Get(fmt.Sprintf("https://git.guangl.cn/api/v4/users?page=%d", page))
	subUsers := make([]*User, 0)
	if err := json.Unmarshal(res, &subUsers); err != nil {
		logs.MyError("GET All Users From Gitlab Failed", err.Error)
		panic(err.Error())
	}
	if len(subUsers) == 0 {
		return true
	}
	Users = append(Users, subUsers...)
	return false
	//defer func() { Wg.Done() }()
}

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	State     string `json:"state"`
	AvatarUrl string `json:"avatar_url"`
	Email     string `json:"email"`
}
