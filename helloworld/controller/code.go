package controller

type ResCode int

const (
	CodeSuccess ResCode = 1000 + iota*10
	CodeInvalidParams
	CodeUserExisted
	CodeUserNotExisted
	CodeUnknownError
)

var codeMsg = map[ResCode]string{
	CodeSuccess:        "成功",
	CodeInvalidParams:  "参数错误",
	CodeUserExisted:    "用户已存在",
	CodeUserNotExisted: "用户名或密码错误",
	CodeUnknownError:   "未知错误",
}

func (c ResCode) Msg() string {
	return codeMsg[c]
}
