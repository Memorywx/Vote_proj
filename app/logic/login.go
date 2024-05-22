package logic

import (
	"WxVote/app/model"
	"WxVote/app/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

func GetLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func PostLogin(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, tools.Ecode{
			Message: err.Error(),
		})
	}

	ret := model.GetUser(user.Name)
	fmt.Println("Retrieved password:", ret.Password)
	fmt.Println("Submitted password:", user.Password)
	if ret.Password != user.Password {
		c.JSON(200, tools.Ecode{
			Message: "账号密码错误!",
		})
		return
	}

	//走到这里说明密码正确，一切正常，设置cookie, 有效时间一小时
	c.SetCookie("name", user.Name, 3600, "/", "", true, false)
	c.SetCookie("Id", fmt.Sprint(ret.Id), 3600, "/", "", true, false)
	c.JSON(200, tools.Ecode{
		Message: "登录成功",
	})
	return
}

func Logout(c *gin.Context) {
	c.SetCookie("name", "", 3600, "/", "", true, false)
	c.SetCookie("Id", "", 3600, "/", "", true, false)
	c.Redirect(http.StatusFound, "/login")
}
