package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/vo"
	"fmt"
	"strconv"
)

// GetAnalysisRetentionData - 获取留存统计数据
// @param {string} merchantCode 商户码
// @param {int} inviteCode 推广码
// @param {string} stime 开始时间
// @param {string} etime 结束时间
// @returns []vo.AnalysisRetentionResp
func GetAnalysisRetentionData(merchantCode string, inviteCode int,
	stime, etime string) []vo.AnalysisRetentionResp {
	data := []vo.AnalysisRetentionResp{}

	// 获取时间范围列表
	dateLis := tool.GetRangeDateList(stime, etime, [3]int{23, 59, 59})
	retentionRanges := map[string]int{
		"day2":  1,
		"day3":  2,
		"day7":  6,
		"day14": 13,
		"day30": 29,
		"day60": 59,
		"day90": 89,
	}

	extWhereSql := ""
	if len(merchantCode) > 0 {
		extWhereSql += " AND merchant_code = '" + merchantCode + "'"
	}

	if inviteCode > 0 {
		extWhereSql += fmt.Sprintf(" AND agent_invite_code = %d", inviteCode)
	}

	for _, sDate := range dateLis {
		temp := map[string]string{}

		// 获取注册人数
		var regNum int64
		global.G_DB.Raw(`
				SELECT
					COUNT(DISTINCT user_id) AS regNum
				FROM
				  fc_user_material
				WHERE DATE(create_time) = DATE(?)`+extWhereSql+`
			`, sDate).Scan(&regNum)

		temp["regnum"] = strconv.FormatInt(regNum, 10)

		for k, days := range retentionRanges {
			// 获取留存计算开始时间
			eDate := tool.GetDistanceDay(sDate, days, [3]int{0, 0, 0})

			var result struct {
				RetentionRate float64 `json:"retention_rate"`
			}

			// 获取登录人数
			var logNum int64
			global.G_DB.Raw(`
				SELECT 
					COUNT(DISTINCT d2.user_id) AS logNum
				FROM 
					(SELECT user_id, create_time
					FROM fc_user_material 
					WHERE DATE(create_time) = DATE(?)`+extWhereSql+`) as d1
				LEFT JOIN 
					fc_login_log d2 ON d1.user_id = d2.user_id 
					AND DATE(d2.create_time) = DATE(?)
			`, sDate, eDate).Scan(&logNum)

			// 获取留存
			global.G_DB.Raw(`
				SELECT 
					COUNT(DISTINCT d2.user_id) * 100.0 / COUNT(DISTINCT d1.user_id) AS retention_rate
				FROM 
					(SELECT user_id, create_time
					FROM fc_user_material 
					WHERE DATE(create_time) = DATE(?)`+extWhereSql+`) as d1
				LEFT JOIN 
					fc_login_log d2 ON d1.user_id = d2.user_id 
					AND DATE(d2.create_time) = DATE(?)
			`, sDate, eDate).Scan(&result)

			temp[k] = fmt.Sprintf("%.2f%% (%d)", result.RetentionRate, logNum)
		}

		tempRes := vo.AnalysisRetentionResp{}
		tool.JsonMapper(temp, &tempRes)
		tempRes.Date = sDate.Format(tool.TimeDateLayoutCN)

		data = append(data, tempRes)
	}

	return data
}
