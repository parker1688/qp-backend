package userControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"image/png"
	"time"
)

const (
	ImgCaptchaKey = "ImgCaptchaKey:%s:%s"
)

type ImgCaptchaStore struct {
}

func (i ImgCaptchaStore) Set(id string, digits []byte) {
	//fmt.Println("set", id, digits)
}

func (i ImgCaptchaStore) Get(id string, clear bool) (digits []byte) {
	return []byte{}
}

func GenVerificationCodeImage(c *gin.Context) {
	var jsonp vo.GenVerificationCodeImageReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	captcha.SetCustomStore(&ImgCaptchaStore{})
	id := captcha.NewLen(4)
	digits := captcha.RandomDigits(4)
	img := captcha.NewImage(id, digits, 160, 80)
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {

	}
	var str string
	for _, v := range digits {
		str += tool.String(v)
	}
	global.G_REDIS.Set(context.Background(), fmt.Sprintf(ImgCaptchaKey, jsonp.Type, id), str, 15*time.Minute)
	data := vo.GenVerificationCodeImageResp{
		VeryCodeRandom: id,
		Image:          "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes()),
	}
	response.SuccessJSON(c, data)
}
