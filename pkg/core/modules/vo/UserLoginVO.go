package vo

type UserLoginVO struct {
	UserName  string `validate:"required" json:"userName" form:"userName"`
	PassWord  string `validate:"required" json:"password" form:"password"`
	AutoLogin bool   `json:"autoLogin" form:"autoLogin"`
	Key       string `form:"key" json:"key"`
}

type CaptchaDataVO struct {
	Dots []int  `json:"dots"`
	Key  string `json:"key"`
}
