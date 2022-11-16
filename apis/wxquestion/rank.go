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

	var data wxquestion.Rank;

	result, count, err := data.GetList()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	app.PageOK(c, result, count, 0, 100, "")
}

// 更新个人数据
func UpdateRank(c *gin.Context) {
	var (
		err          error
		rankscore wxquestion.Rank
		dbscore wxquestion.Rank
		rankCount int
	)

	err = c.ShouldBind(&rankscore)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	logger.Info(rankscore)

	err = orm.Eloquent.Model(&dbscore).
		Where("name = ?", rankscore.Name).
		Count(&rankCount).Error
	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("流程信息查询失败，%v", err.Error()))
		return
	}
	if rankCount > 0 {
		err = orm.Eloquent.Model(&rankscore).
		Where("name = ?", rankscore.Name).
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
