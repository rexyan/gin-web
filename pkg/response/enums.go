package response

const (
	SUCCESS                     string = "90000"
	ParamError                  string = "30011"
	UnregisteredOrPasswordError string = "30012"
	UserExist                   string = "30013"
	UserRegisterError           string = "30014"
	UserLoginError              string = "30015"
)

var Message = map[string]string{
	SUCCESS:                     "成功",
	ParamError:                  "参数错误",
	UnregisteredOrPasswordError: "未注册或者密码错误",
	UserExist:                   "用户已存在",
	UserRegisterError:           "用户注册错误",
	UserLoginError:              "用户登录失败",
}
