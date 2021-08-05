package response

const (
	SUCCESS                     string = "90000"
	ParamError                  string = "30011"
	UnregisteredOrPasswordError string = "30012"
	UserExist                   string = "30013"
	UserRegisterError           string = "30014"
	UserLoginError              string = "30015"
	RefreshTokenError           string = "30016"
	AccessTokenFormatError      string = "30017"
	AccessTokenInvalid          string = "30018"
	CommunityListError          string = "30019"
	CommunityIDError            string = "30020"
	CommunityDetailError        string = "30021"
	CreatePostError             string = "30022"
	PostDetailError             string = "30023"
	PostListError               string = "30024"
)

var Message = map[string]string{
	SUCCESS:                     "成功",
	ParamError:                  "参数错误",
	UnregisteredOrPasswordError: "未注册或者密码错误",
	UserExist:                   "用户已存在",
	UserRegisterError:           "用户注册错误",
	UserLoginError:              "用户登录失败",
	RefreshTokenError:           "获取 RefreshToken 失败",
	AccessTokenFormatError:      "AccessToken 格式错误",
	AccessTokenInvalid:          "无效的 AccessToken",
	CommunityListError:          "获取社区列表错误",
	CommunityIDError:            "获取社区 ID 错误",
	CommunityDetailError:        "获取社区详情错误",
	CreatePostError:             "创建帖子错误",
	PostDetailError:             "获取帖子详情错误",
	PostListError:               "获取帖子列表错误",
}
