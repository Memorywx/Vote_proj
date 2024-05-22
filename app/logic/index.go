package logic

import "C"
import (
	"WxVote/app/model"
	"WxVote/app/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	ret := model.GetVotes()
	c.HTML(http.StatusOK, "index.html", gin.H{"vote": ret})
}

func GetVoteInfo(c *gin.Context) {
	var id int64
	idStr := c.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	ret := model.GetVote(id)
	c.HTML(200, "vote.html", gin.H{"vote": ret})
}

func DoVOte(c *gin.Context) {
	userIDStr, _ := c.Cookie("Id")
	voteIDStr, _ := c.GetPostForm("vote_id")
	optStr, _ := c.GetPostFormArray("opt[]")

	userID, _ := strconv.ParseInt(userIDStr, 10, 64)
	voteID, _ := strconv.ParseInt(voteIDStr, 10, 64)
	opt := make([]int64, 0)
	for _, v := range optStr {
		optId, _ := strconv.ParseInt(v, 10, 64)
		opt = append(opt, optId)
	}

	model.DoVote(userID, voteID, opt)
	c.JSON(200, tools.Ecode{
		Message: "投票成功!",
	})
}
