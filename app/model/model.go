package model

import (
	"time"
)

type VoteOptUser struct {
	Id          int64     `gorm:"column:id ;primary_key;AUTO_INCREMENT; NOT NULL"`
	UserId      int64     `gorm:"column:user_id"`
	VoteId      int64     `gorm:"column:vote_id"`
	VoteOptId   int64     `gorm:"column:vote_opt_id"`
	CreatedTime time.Time `gorm:"column:created_time;default:CURRENT_TIMESTAMP"`
}

func (v *VoteOptUser) TableName() string {
	return "vote_opt_user"
}

type VoteOpt struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT; NOT NULL"`
	Name        string    `gorm:"column:name"`
	VoteId      int64     `gorm:"column:vote_id"`
	Count       int32     `gorm:"column:count"`
	CreatedTime time.Time `gorm:"column:created_time;default:CURRENT_TIMESTAMP"`
	UpdatedTime time.Time `gorm:"column:updated_time;default:CURRENT_TIMESTAMP"`
}

func (v *VoteOpt) TableName() string {
	return "vote_opt"
}

type Vote struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Title       string    `gorm:"column:title"`
	Type        int32     `gorm:"column:type;comment:'0单选1多选'"`
	Status      int32     `gorm:"column:status;comment:'0正常1超时'"`
	Time        int64     `gorm:"column:time;comment:'有效时长'"`
	UserId      int64     `gorm:"column:user_id;comment:'创建人'"`
	CreatedTime time.Time `gorm:"column:created_time;default:CURRENT_TIMESTAMP"`
	UpdatedTime time.Time `gorm:"column:updated_time;default:CURRENT_TIMESTAMP"`
}

func (v *Vote) TableName() string {
	return "vote"
}

type User struct {
	Id          int64     `gorm:"column:id; primary_key; AUTO_INCREMENT; NOT NULL"`
	Name        string    `gorm:"column:name; default:NULL"`
	Password    string    `gorm:"column:password; default:NULL"`
	CreatedTime time.Time `gorm:"column:created_time; default:CURRENT_TIMESTAMP"`
	UpdatedTime time.Time `gorm:"column:updated_time; default:CURRENT_TIMESTAMP"`
}

func (User) TableName() string {
	return "user"
}

type VoteWithOpt struct {
	Vote Vote
	Opt  []VoteOpt
}
