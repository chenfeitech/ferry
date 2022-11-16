package wxquestion

import (
	"ferry/models/base"
	"ferry/global/orm"
)

/*
  @Author : helight
*/

// 任务
type Rank struct {
	base.Model
	Uid		string `gorm:"column:uid; type: varchar(256)" json:"uid" form:"uid"`  // 用户id
	Name	string `gorm:"column:name; type: varchar(256)" json:"name" form:"name"`  // 用户名
	Score	int    `gorm:"column:score; type: int(11)" json:"score" form:"score"`      // 分数
	Openid	string `gorm:"column:openid; type: varchar(128)" json:"openid" form:"openid"`  // openid
	Info	string `gorm:"column:info; type: longtext" json:"info" form:"info"`        // 备注
}

func (Rank) TableName() string {
	return "wx_rank"
}

func (rank *Rank) GetList() ([]Rank, int, error) {
	var (
		rlist   []Rank
		count int
	)

	table := orm.Eloquent.Table("wx_rank")

	if err := table.Order("score desc").Find(&rlist).Error; err != nil {
		return nil, 0, err
	}
	table.Count(&count)

	return rlist, count, nil
}