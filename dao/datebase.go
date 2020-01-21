package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gitlab_tapd/common"
	"gitlab_tapd/logs"
)

var Ggorm *gorm.DB

func InitDB() *gorm.DB {
	if gg, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%d", common.POST_HOST, common.POST_USER, common.POST_DBNAME, common.POST_PASSWORD, common.POST_PORT)); err != nil {
		logs.MyInfo("Db Init Failed", err)
	} else {
		gg.DB().SetMaxIdleConns(2)
		gg.DB().SetMaxOpenConns(5)
		gg.LogMode(true)
		if err := gg.DB().Ping(); err != nil {
			logs.MyError("Db Init Failed By Test", err)
		} else {
			logs.MyInfo("Db Init Success")
			return gg
		}
	}
	return nil
}
