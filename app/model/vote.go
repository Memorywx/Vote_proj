package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

func GetVotes() []Vote {
	ret := make([]Vote, 0)
	if err := Conn.Table("vote").Find(&ret).Error; err != nil {
		fmt.Println("err:", err.Error())
	}
	return ret
}

func GetVote(id int64) VoteWithOpt {
	var ret Vote
	if err := Conn.Table("vote").Where("id = ?", id).First(&ret).Error; err != nil {
		fmt.Printf("err : %s", err.Error())
	}
	opt := make([]VoteOpt, 0)
	if err := Conn.Table("vote_opt").Where("vote_id = ?", id).Find(&opt).Error; err != nil {
		fmt.Printf("err : %s", err.Error())
	}
	return VoteWithOpt{
		Vote: ret,
		Opt:  opt,
	}
}

func DoVote(userId, voteId int64, optIds []int64) bool {
	// 投票选项加1
	var ret Vote
	if err := Conn.Table("vote").Where("id = ?", voteId).First(&ret).Error; err != nil {
		fmt.Printf("err : %s", err.Error())
	}
	for _, value := range optIds {
		if err := Conn.Table("vote_opt").Where("id = ?", value).Update("count", gorm.Expr("count + ?", 1)).Error; err != nil {
			fmt.Printf("err : %s", err.Error())
		}
		// 记录用户投票记录
		user := VoteOptUser{
			VoteId:      voteId,
			UserId:      userId,
			VoteOptId:   value,
			CreatedTime: time.Now(),
		}
		_ = Conn.Create(&user).Error
	}

	return true
}
