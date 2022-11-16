package wxquestion

import (
  "crypto/sha256"
  "fmt"
  "log"
  "net/http"

  "ferry/pkg/authjwt"
  "ferry/pkg/service"

  "github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/base/request"
  "github.com/gin-gonic/gin"
)

func APISNSSession(c *gin.Context) {
  // {"openid":"otFFw5A6eLBUxQxmiUskp2lXVpNY","session_key":"RVLHimKLIauQ7jj4H2SRfA==","unionid":""}
  type AuthInfo struct {
		OpenID string `json:"openid"`
		SessionKey string `json:"session_key"`
    UnionID string `json:"unionid"`
    Token string `json:"token"`
	}

  code, exist := c.GetQuery("code")
  if !exist {
    panic("parameter code expected")
  }

  rs, err := service.MiniProgramApp.Auth.Session(code)

  if err != nil {
    panic(err)
  }
  tokenString, err := authjwt.GenerateJWT(rs.OpenID, rs.SessionKey)
	if err != nil {
		// context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		panic(err)
		return
	}

  var authoinfo AuthInfo
  authoinfo.OpenID = rs.OpenID
  authoinfo.SessionKey = rs.SessionKey
  authoinfo.UnionID = rs.UnionID
  authoinfo.Token = tokenString

  c.JSON(http.StatusOK, authoinfo)

}

func APICheckEncryptedData(c *gin.Context) {
  encryptedData := c.DefaultQuery("encryptedData", "sTWzm26PrbsXlSA8AoW+GpiyNLJP0H5p2UT4dXKwLSvXv8aU4wIiJcZUcM/IzNXnoFtERY3BDRbZh6bwd0ZGENVhucqDPXmchTqseryIZnJiKsiNMHCpAkCA2Yl00q4UpOZYtGMuTX5BTuo1yB3bOOuIfDu6neHV3D158CofGB9m7TxFQ8A/JcauWzhvmEAPygfFaqCgDTEmluLu7S8wMA==")
  hashByte := sha256.Sum256([]byte(encryptedData))
  hash := hashByte[:]
  rs, err := service.MiniProgramApp.Base.CheckEncryptedData(fmt.Sprintf("%x", hash))

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)

}

func APIGetPaidUnionID(c *gin.Context) {
  openid := c.DefaultQuery("openid", "")
  log.Printf("openid: %s\n", openid)
  rs, err := service.MiniProgramApp.Base.GetPaidUnionID(&request.RequestGetPaidUnionID{
    OpenID: openid,
    // TransactionID: "",
    // MchID:         "",
    // OutTradeNo:    "",
  })

  if err != nil {
    panic(err)
  }

  c.JSON(http.StatusOK, rs)

}
