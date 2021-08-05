package validator

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)
		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		return
	}
	return
}

func TransError(err validator.ValidationErrors) string {
	var errStr []string
	translate := err.Translate(Trans)
	for _, v := range translate {
		errStr = append(errStr, v)
	}
	return strings.Join(errStr, ",")
}

// 注册参数校验
type RegisterValidator struct {
	Username   string `json:"username" binding:"required"`                     // 用户名
	Password   string `json:"password" binding:"required"`                     // 密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password"` // 重复密码
	Email      string `json:"email" binding:"required,email"`                  // 邮箱
	Gender     uint8  `json:"gender" binding:"gte=0,lte=1"`                    // 性别
}

// 登录参数校验
type LoginValidator struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

// 创建帖子参数校验
type CreatePostValidator struct {
	Title       string `json:"title" binding:"required"`        // 名称
	Content     string `json:"content" binding:"required"`      // 内容
	CommunityId int64  `json:"community_id" binding:"required"` // 所属社区ID
}

type PostDetail struct {
	PostID        int64     `json:"post_id"`        // 帖子 ID
	Title         string    `json:"title"`          // 帖子名称
	Content       string    `json:"content"`        // 帖子内容
	UserId        int64     `json:"user_id"`        // 用户 ID
	UserName      string    `json:"user_name"`      // 用户名称
	CommunityId   int64     `json:"community_id"`   // 社区 ID
	CommunityName string    `json:"community_name"` // 社区名称
	CreateTime    time.Time `json:"create_time"`    // 帖子创建时间
}
