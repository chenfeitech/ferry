package wxquestion

import (
  "ferry/apis/wxquestion"

  jwt "ferry/pkg/jwtauth"
  "github.com/gin-gonic/gin"
)

func InitMiniProgramAPIRoutes(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
  routerMiniProgram := v1.Group("/miniprogram")
  {
    // Handle the auth route
    routerMiniProgram.GET("/auth", wxquestion.APISNSSession)
    routerMiniProgram.POST("/auth/checkEncryptedData", wxquestion.APICheckEncryptedData)
    routerMiniProgram.GET("/auth/getPaidUnionID", wxquestion.APIGetPaidUnionID)
  }
}
