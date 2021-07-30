package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	"web_app/settings"
)

const TokenExpireDuration = time.Hour * 2
const RefreshExpireDuration = TokenExpireDuration * 10

type CustomClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

func generateJwtToken(userID int64, userName string, timeDuration time.Duration) (string, error) {
	var Secret = []byte(settings.Config.ServerConfig.Secret)
	// 创建一个我们自己的声明
	c := CustomClaims{
		userID,
		userName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(timeDuration).Unix(), // 过期时间
			Issuer:    settings.Config.ServerConfig.Name,   // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Secret)
}

func GenerateJwtRefreshToken(userID int64, userName string) (string, error) {
	return generateJwtToken(userID, userName, RefreshExpireDuration)
}

func GenerateJwtAccessToken(userID int64, userName string) (string, error) {
	return generateJwtToken(userID, userName, TokenExpireDuration)
}

func ParseToken(tokenString string) (*CustomClaims, error) {
	var Secret = []byte(settings.Config.ServerConfig.Secret)
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
