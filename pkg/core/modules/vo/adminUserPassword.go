package vo

type AdminUserPasswordVO struct {
	OldPwd     string `gorm:"-" json:"old_password" form:"old_password" uri:"old_password" `             // 旧密码
	NewPwd     string `gorm:"-" json:"new_password" form:"new_password" uri:"new_password" `             // 新密码
	ConfirmPwd string `gorm:"-" json:"confirm_password" form:"confirm_password" uri:"confirm_password" ` //确认新密码
}

type AdminUserSecurityReq struct {
	Mobile        string `json:"mobile" form:"mobile" uri:"mobile"`
	ConfirmMobile string `json:"confirm_mobile" form:"confirm_mobile" uri:"confirm_mobile"`
	NewPwd        string `json:"new_password" form:"new_password" uri:"new_password"`
	ConfirmPwd    string `json:"confirm_password" form:"confirm_password" uri:"confirm_password"`
}
