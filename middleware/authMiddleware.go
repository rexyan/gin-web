package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
	"web_app/pkg/jwt"
	"web_app/pkg/response"
)

const CtxUserID = "user_id"
const Bearer = "Bearer"
const AuthorizationKey = "Authorization"

func validateAccessToken(accessToken string) (userID int64, err error) {
	if claims, err := jwt.ParseToken(accessToken); err != nil {
		return 0, err
	} else {
		return claims.UserID, nil
	}
}

func checkAccessTokenFormat(accessToken string) (token string, qualified bool) {
	splitStringSlice := strings.Split(accessToken, " ")
	if len(splitStringSlice) == 2 && splitStringSlice[0] == Bearer {
		return splitStringSlice[1], true
	}
	return "", false
}

func JwtAuthMiddleware(c *gin.Context) {
	var effectiveAccessToken bool
	var effectiveAccessTokenFormat bool
	var userID int64

	// 从请求头中获取 AccessToken
	headerAuthorization := c.Request.Header.Get(AuthorizationKey)
	// 从 URL 中获取 AccessToken
	urlAuthorization := c.DefaultQuery(AuthorizationKey, "")

	if headerAuthorization != "" {
		if token, qualified := checkAccessTokenFormat(headerAuthorization); qualified {
			if id, _ := validateAccessToken(token); id != 0 {
				userID = id
				effectiveAccessToken = true
			}
			effectiveAccessTokenFormat = true
		}
	}

	// 没有在请求头中认证通过就进行在 URL 进行认证
	if urlAuthorization != "" && !effectiveAccessToken {
		if token, qualified := checkAccessTokenFormat(urlAuthorization); qualified {
			if id, _ := validateAccessToken(token); id != 0 {
				userID = id
				effectiveAccessToken = true
			}
			effectiveAccessTokenFormat = true
		}
	}

	if !effectiveAccessTokenFormat {
		zap.L().Error("Invalid AccessToken Format")
		response.BuildResponse(c, nil, response.AccessTokenFormatError, 400)
		c.Abort()
		return
	}

	if !effectiveAccessToken {
		zap.L().Error("Invalid AccessToken")
		response.BuildResponse(c, nil, response.AccessTokenInvalid, 400)
		c.Abort()
		return
	}
	c.Set(CtxUserID, userID)
	c.Next()
}
