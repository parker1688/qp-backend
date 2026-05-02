package fcPayment

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

func SyncFcPayment(c *gin.Context) {
	channels := modules.FindByKeyFcPayChannel(&dos.FcPayChannel{}, nil)
	for _, v := range channels {

		pay := modules.FindByKeyFcPayment(&dos.FcPayment{
			ChannelCode: v.ChannelCode,
			Status:      1,
		}, nil)
		if len(pay) > 0 {
			rangeMap := make(map[float64]struct{})

			var minAmount, maxAmount float64

			for _, vv := range pay {
				if vv.MinAmount > 0 && minAmount < vv.MinAmount {
					minAmount = vv.MinAmount
				}

				if vv.MaxAmount > 0 && maxAmount < vv.MaxAmount {
					maxAmount = vv.MaxAmount
				}
				if len(vv.AmountRange) > 0 {
					ar := tool.ToFloat64Zero(vv.AmountRange)
					for _, arv := range ar {
						rangeMap[arv] = struct{}{}
					}
				}
			}

			//排序
			rangeArray := make([]float64, 0, len(rangeMap))
			for k, _ := range rangeMap {
				rangeArray = append(rangeArray, k)
			}
			sort.Float64s(rangeArray)

			//转为string
			strNums := make([]string, len(rangeArray))
			for i, vNum := range rangeArray {
				strNums[i] = tool.String(int(vNum))
			}
			//区间值
			//最大值
			//最小值
			v.MinAmount = minAmount
			v.MaxAmount = maxAmount
			v.AmountRange = strings.Join(strNums, ",")
			var inputAmountDisplay int
			if v.MinAmount > 0 {
				inputAmountDisplay = 1
			}
			global.G_DB.Model(&dos.FcPayChannel{}).Where("id = ?", v.Id).Updates(map[string]interface{}{
				"min_amount":           v.MinAmount,
				"max_amount":           v.MaxAmount,
				"amount_range":         v.AmountRange,
				"input_amount_display": inputAmountDisplay,
				"status":               1, //上线
			})
		} else {
			global.G_DB.Model(&dos.FcPayChannel{}).Where("id = ?", v.Id).Updates(map[string]interface{}{
				"status": 2, //下线
			})
		}
	}
	response.SuccessJSON(c, nil)
}

func FcPaymentBatchStatus(c *gin.Context) {
	jsonp := struct {
		Ids    []string `json:"ids"`    //ID集合
		Status int      `json:"status"` //状态 1 上线 2 下线
	}{}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	global.G_DB.Model(&dos.FcPayment{}).Where("id in ?", jsonp.Ids).Update("status", jsonp.Status)
	response.SuccessJSON(c, true)
}
