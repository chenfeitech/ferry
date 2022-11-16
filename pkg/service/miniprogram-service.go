package service

import (
  "log"
  "os"

  "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
  "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
  "github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram"
)

var MiniProgramApp *miniProgram.MiniProgram

const TIMEZONE = "asia/shanghai"
const DATETIME_FORMAT = "20060102"

func NewMiniMiniProgramService(appid, secret, redisadd string) (*miniProgram.MiniProgram, error) {
  log.Printf("miniprogram app_id: %s", os.Getenv("miniprogram_app_id"))
  var cache kernel.CacheInterface
  if redisadd != "" {
    cache = kernel.NewRedisClient(&kernel.RedisOptions{
      Addr: redisadd,
    })
  }
  app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
    AppID:        appid,  // 小程序、公众号或者企业微信的appid
    Secret:       secret, // 商户号 appID
    ResponseType: response.TYPE_MAP,
    Log: miniProgram.Log{
      Level: "debug",
      File:  "./wechat.log",
    },
    //"sandbox": true,
    Cache:     cache,
    HttpDebug: true,
    Debug:     false,
  })

  return app, err
}
