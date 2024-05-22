package app

import (
	"WxVote/app/model"
	"WxVote/app/router"
)

// 启动器方法

func Start() {
	model.MysqlConnect()
	defer func() {
		model.Close()
	}()

	router.New()
}
