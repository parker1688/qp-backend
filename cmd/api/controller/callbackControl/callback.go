package callbackControl

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/venues/venuevo"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PGAuthSuccessResp struct {
	Data struct {
		PlayerName string `json:"player_name"`
		Nickname   string `json:"nickname"`
		Currency   string `json:"currency"`
	} `json:"data"`
	Error interface{} `json:"error"`
}

type PGAuthFailResp struct {
	Error struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	} `json:"error"`
	Data interface{} `json:"data"`
}

type WUGAuthResp struct {
	StatusCode int    `json:"status_code"` // 0 为成功， 1 失败
	MemberId   string `json:"member_id"`
	Balance    int    `json:"balance"`
}

// PGVerifySession
//
//	@Description: PG回调验证
//	@param c
func PGVerifySession(c *gin.Context) {

	token := c.PostForm("operator_player_session")
	operator_token := c.PostForm("operator_token")
	secret_key := c.PostForm("secret_key")
	global.G_LOG.Infof("PG回调 operator_player_session：%s   operator_token:%s  secret_key:%s", token, operator_token, secret_key)
	//config := global.CONFIG.Venue.PGDZ
	//if operator_token != config.OperatorToken {
	//	pgAuthFailResp := PGAuthFailResp{}
	//	pgAuthFailResp.Error.Code = "1034"
	//	pgAuthFailResp.Error.Message = "AuthFail"
	//	c.JSON(200, pgAuthFailResp)
	//	return
	//}
	//if secret_key != config.SecretKey {
	//	pgAuthFailResp := PGAuthFailResp{}
	//	pgAuthFailResp.Error.Code = "1034"
	//	pgAuthFailResp.Error.Message = "AuthFail"
	//	c.JSON(200, pgAuthFailResp)
	//	return
	//}

	tokenKey := fmt.Sprintf(enmus.PG_CALLBACK_TOKEN_KEY, token)
	userVal := global.G_REDIS.Get(context.Background(), tokenKey).Val()
	global.G_LOG.Infof("PG回调 tokenKey:%s", tokenKey)
	global.G_LOG.Infof("PG回调 userVal:%s", userVal)
	var user *venuevo.TokenData
	err := tool.JsonUnmarshalFromString(userVal, &user)
	global.G_LOG.Infof("PG回调 JsonUnmarshalFromString user:%v  err:%v", user, err)
	if err != nil {
		pgAuthFailResp := PGAuthFailResp{}
		pgAuthFailResp.Error.Code = "1034"
		pgAuthFailResp.Error.Message = "AuthFail"
		c.JSON(200, pgAuthFailResp)
		return
	}

	global.G_LOG.Infof("PGVerifySession token:%s data:%v", token, user)
	if user.UserName != "" {
		pgAuthSuccessResp := PGAuthSuccessResp{}
		pgAuthSuccessResp.Data.Nickname = user.UserName
		pgAuthSuccessResp.Data.PlayerName = user.NickName
		pgAuthSuccessResp.Data.Currency = global.CONFIG.General.DefaultCurrency

		c.JSON(200, pgAuthSuccessResp)
		return
	}

	pgAuthFailResp := PGAuthFailResp{}
	pgAuthFailResp.Error.Code = "1034"
	pgAuthFailResp.Error.Message = "AuthFail"
	c.JSON(200, pgAuthFailResp)
}

// WUGVerifySession
//
//	@Description: WUG回调验证
//	@param c
func WUGVerifySession(c *gin.Context) {
	token := c.Query("access_token")
	resp := WUGAuthResp{}
	if token == "" {
		global.G_LOG.Infof("WUGDZ-CALLBACK access_token is empty")
		resp.StatusCode = 1
		c.JSON(200, resp)
		return
	}

	tokenKey := fmt.Sprintf(enmus.WUG_CALLBACK_TOKEN_KEY, token)
	userVal := global.G_REDIS.Get(context.Background(), tokenKey).Val()
	//global.G_LOG.Infof("WUGDZ-CALLBACK tokenKey: %s, userVal: %s", tokenKey, userVal)
	if userVal == "" {
		global.G_LOG.Infof("WUGDZ-CALLBACK access_token value is empty")
		resp.StatusCode = 1
		c.JSON(200, resp)
		return
	}

	var user *venuevo.TokenData
	err := tool.JsonUnmarshalFromString(userVal, &user)
	//global.G_LOG.Infof("WUGDZ-CALLBACK JsonUnmarshalFromString user:%v  err:%v", user, err)
	if err != nil {
		resp.StatusCode = 1
		global.G_LOG.Infof("WUGDZ-CALLBACK json-unmarshal")
		c.JSON(200, resp)
		return
	}

	//global.G_LOG.Infof("WUGDZ-CALLBACK verifySession token:%s data:%v", token, user)
	if user.UserName != "" {
		resp.StatusCode = 0
		resp.MemberId = user.UserName
		c.JSON(200, resp)
		return
	}
	//global.G_LOG.Infof("WUGDZ-CALLBACK userName:%s data:%v", token, user.UserName)
	resp.StatusCode = 0
	c.JSON(200, resp)
}
