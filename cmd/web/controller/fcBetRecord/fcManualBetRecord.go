package fcBetRecord

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/go-resty/resty/v2"
)

func ManualPullRecord(c *gin.Context) {
	var jsonp dos.FcManualBetRecordReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err1 := validator.New().Struct(jsonp)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}
	startDate := jsonp.StartAt
	venue := jsonp.VenueCode
	startAt := tool.DateToTimeStamp("2006-01-02 15:04:05", startDate)
	//这是某个会员的注单
	param := fmt.Sprintf("venue=%s&startAt=%v", venue, startAt)
	//global.G_LOG.Infof("manual pull record -------------------------------0:%v, %v", startAt, param)
	client := resty.New()
	url := global.CONFIG.General.BetRecordUrl + "/betrecord/manualPullBetRecord?" + param
	resp, err2 := client.R().Get(url)

	//global.G_LOG.Infof("manual pull record -------------------------------1:%v", url)
	if err2 != nil {
		global.G_LOG.Errorf("manual pull betrecord err, param: %s apiUrl: %s err: %v", param, url, err2)
	}
	respData := struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{}
	//global.G_LOG.Infof("manual pull record -------------------------------2:%v", resp.String())
	err = tool.JsonUnmarshal(resp.Body(), &respData)
	if err != nil {
		global.G_LOG.Errorf("manual pull betrecord jsonunmarshal err, err: %v", err)
	}

	rsdata := dos.FcManualBetRecordResp{Code: respData.Code, Msg: respData.Msg}
	//global.G_LOG.Infof("manual pull record -------------------------------3:%v", rsdata)
	response.SuccessJSON(c, rsdata)
}
