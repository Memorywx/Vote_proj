package router

import (
	"WxVote/app/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func New() {
	r := gin.Default()
	r.LoadHTMLGlob("app/view/*")

	// 放相关的路径在这里
	index := r.Group("")

	index.Use(checkUser)
	index.GET("/index", logic.Index)
	index.GET("/vote", logic.GetVoteInfo)
	index.POST("/vote", logic.DoVOte)

	r.GET("/", logic.Index)
	r.GET("/login", logic.GetLogin)
	r.POST("/login", logic.PostLogin)
	r.GET("logout", logic.Logout)
	if err := r.Run(":8080"); err != nil {
		panic("gin 启动失败!")
	}
}

func checkUser(c *gin.Context) {
	name, err := c.Cookie("name")
	if err != nil || name == "" {
		c.Redirect(http.StatusFound, "/login")
	}
	c.Next()
}
