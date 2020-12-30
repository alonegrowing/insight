package macro

type Error struct {
	Code int64
	Msg  string
}

const (
	Success           = 0
	ErrorParamIllegal = 10000
	ErrorAuthFailed   = 1000000
)

var ErrMsg = map[int64]string{
	Success:           "成功",
	ErrorParamIllegal: "参数非法",
	ErrorAuthFailed:   "登陆校验失败",
}
