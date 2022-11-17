package authjwt

import (
	"errors"
	"time"

	// "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// 一些常量
var (
    TokenExpired     error  = errors.New("Token is expired")
    TokenNotValidYet error  = errors.New("Token not active yet")
    TokenMalformed   error  = errors.New("That's not even a token")
    TokenInvalid     error  = errors.New("Couldn't handle this token:")
)

var jwtSigningKey = []byte("91440300MA5H0J0J9W")

func jwtKeyFunc(token *jwt.Token) (i interface{}, err error) {
	return jwtSigningKey, nil
 }

type JWTClaim struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
	jwt.StandardClaims 
}

func GenerateJWT(key1, key2 string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Key1: key1,
		Key2: key2,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()},
	}
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, err = token.SignedString(jwtSigningKey)
	return CreateToken(*claims)
}
// CreateToken 生成一个token
func CreateToken(claims JWTClaim) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSigningKey)
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(signedToken, &JWTClaim{}, jwtKeyFunc)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claim")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
// 更新token
func RefreshToken(tokenString string) (string, error) {
    jwt.TimeFunc = func() time.Time {
        return time.Unix(0, 0)
    }
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, jwtKeyFunc)
    if err != nil {
        return "", err
    }
    if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
        jwt.TimeFunc = time.Now
        claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
        return CreateToken(*claims)
    }
    return "", TokenInvalid
}


// RefreshToken2 使用refreshtoken 刷新 AccessToken
func RefreshToken2(aToken, rToken string) (newAToken, newRToken string, err error) {
   // refresh token无效直接返回
   if _, err = jwt.Parse(rToken, jwtKeyFunc); err != nil {
      return
   }

   // 从旧access token中解析出claims数据
   var claims JWTClaim
   _, err = jwt.ParseWithClaims(aToken, &claims, jwtKeyFunc)
   v, _ := err.(*jwt.ValidationError)

   // 当access token是过期错误 并且 refresh token没有过期时就创建一个新的access token
   if v.Errors == jwt.ValidationErrorExpired {
      token, _ := CreateToken(claims)
      return token, "", nil
   }
   return
}


// 简单认证
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("AuthToken")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}