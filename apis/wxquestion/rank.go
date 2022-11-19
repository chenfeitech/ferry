package wxquestion

import (
	"fmt"

	"ferry/tools/app"
	"ferry/pkg/logger"
	"ferry/global/orm"
	"ferry/models/wxquestion"

	"github.com/gin-gonic/gin"
)

/*
 @Author : helight
*/

// 获取排行榜
func RankList(c *gin.Context) {
	type RankListID struct {
		Examid string `json:"examid" form:"examid"`  // 用户id
		Subjectid string `json:"subjectid" form:"subjectid"`  // 用户id
	}

	var rankListID RankListID

	err := c.ShouldBind(&rankListID)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	logger.Info(rankListID)

	var data wxquestion.Rank;

	result, count, err := data.GetList(rankListID.Subjectid, rankListID.Examid)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	app.PageOK(c, result, int(count), 0, 100, "")
}

// 更新个人数据
func UpdateRank(c *gin.Context) {
	var (
		err          error
		rankscore wxquestion.Rank
		dbscore wxquestion.Rank
		rankCount int64
	)

	err = c.ShouldBind(&rankscore)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	logger.Info(rankscore)

	err = orm.Eloquent.Model(&dbscore).
		Where("name = ? and examid = ? and subjectid = ?", rankscore.Name, rankscore.Examid, rankscore.Subjectid).
		Count(&rankCount).Error
	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("流程信息查询失败，%v", err.Error()))
		return
	}
	if rankCount > 0 {
		err = orm.Eloquent.Model(&rankscore).
		Where("name = ? and examid = ? and subjectid = ?", rankscore.Name, rankscore.Examid, rankscore.Subjectid).
		Updates(map[string]interface{}{
			"score":      rankscore.Score,
		}).Error
	} else {
		err = orm.Eloquent.Create(&rankscore).Error
		if err != nil {
			app.Error(c, -1, err, fmt.Sprintf("更新失败，%v", err.Error()))
			return
		}
	}

	app.OK(c, rankscore, "更新成功")
}
