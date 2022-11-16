package wxquestion

import (
	"ferry/apis/wxquestion"
	// "ferry/middleware"
	"ferry/pkg/authjwt"
	jwt "ferry/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

/*
  @Author : helight
*/

func RegisterRankRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	// taskRouter := v1.Group("/wx").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	taskRouter := v1.Group("/wx").Use(authjwt.Auth())
	{
		taskRouter.GET("/ranklist", wxquestion.RankList)
		taskRouter.POST("/upscore", wxquestion.UpdateRank)
	}
}
