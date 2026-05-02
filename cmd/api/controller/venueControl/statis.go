package venueControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/venues"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var venueList = map[string]string{
	"TYQP":  "elecgame,chess,fish",
	"LYQP":  "elecgame,chess,fish",
	"FGDZ":  "elecgame,chess,fish",
	"MTQP":  "elecgame,chess,fish",
	"WUGDZ": "elecgame",
	"FBTY":  "sport",
	"HGTY":  "sport",
	"KYQP":  "elecgame,chess,fish",
	"MGDZ":  "elecgame,chess,fish",
	"AGZR":  "live",
	"BGZR":  "live",
	"CQ9":   "elecgame",
	"JDB":   "elecgame,chess,fish",
	"VRCP":  "lottery",
	"PGDZ":  "elecgame",
	"BBIN":  "live,sport,elecgame,lottery",
	"WALI":  "live,sport,chess",
	//"VGQP":  "",
	"DBDJ": "esport",
	"JLDZ": "elecgame",
	"PPDZ": "elecgame",
	"PTDZ": "elecgame",
	"KXDZ": "chess",
	"SBTY": "sport",
}

func Statis(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		vo.BetRecordReq
	}{}

	jsonp.PageQuery.PageNo = tool.Atoi(c.DefaultQuery("current", "1"))
	jsonp.PageQuery.PageSize = tool.Atoi(c.DefaultQuery("pageSize", "10"))
	jsonp.VenueCode = c.DefaultQuery("venue_code", "")
	jsonp.DateEnd = c.DefaultQuery("date_end", "")
	jsonp.DateStart = c.DefaultQuery("date_start", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", "0"))
	timeType := tool.Atoi(c.DefaultQuery("time_type", "-1"))

	if jsonp.PageNo == 0 {
		jsonp.PageNo = 1
	}
	if jsonp.PageSize == 0 {
		jsonp.PageSize = 10
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	k1 := 0
	st := ""
	et := ""
	if timeType == -1 {
		if jsonp.DateStart != "" {
			if len(jsonp.DateStart) < 11 {
				jsonp.DateStart += " 00:00:00"
			}
		}
		if jsonp.DateEnd != "" {
			if len(jsonp.DateEnd) < 11 {
				jsonp.DateEnd += " 23:59:59"
			}
		}
	} else {
		k1 = 1
		sTime1, eTime1 := tool.GetDayRange(time.Now(), timeType)
		st = sTime1.Format(tool.TimeLayout)
		et = eTime1.Format(tool.TimeLayout)
	}
	keyData := fmt.Sprintf("venue_record_statis:%s:%s:%v:%s:%s:%v:%v:%v", userInfo.UserId, jsonp.VenueCode, timeType, jsonp.DateStart, jsonp.DateEnd, jsonp.PageNo, jsonp.PageSize, jsonp.Status)
	if k1 == 1 {
		keyData = fmt.Sprintf("venue_record_statis:%s:%s:%v:%s:%s:%v:%v:%v", userInfo.UserId, jsonp.VenueCode, timeType, st, et, jsonp.PageNo, jsonp.PageSize, jsonp.Status)
	}
	resData, ok := venuesCache.Get(keyData)
	if !ok {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 8*time.Second)
		defer cancel()
		list := []vo.BetRecordStatisData{}
		for venueCode, typeList := range venueList {
			typeList2 := strings.Split(typeList, ",")
			venueName := venues.VenueNameList[venueCode]
			for _, gameType := range typeList2 {
				win := 0.0
				betAmount := 0.0
				betcount := 0
				betAmount2 := 0.0
				betcount2 := 0

				query1 := global.G_DB.WithContext(ctx).Model(dos.FcBetRecord{})
				query1.Where("user_id = ? and venue_code = ? and game_type = ?", userInfo.UserId, venueCode, gameType)
				if timeType == -1 {
					if jsonp.DateStart != "" {
						query1.Where("bet_time >=?", jsonp.DateStart)
					}
					if jsonp.DateEnd != "" {
						query1.Where("bet_time <=?", jsonp.DateEnd)
					}
				} else {
					sTime, eTime := tool.GetDayRange(time.Now(), timeType)
					query1.Where("bet_time BETWEEN ? AND ?", sTime, eTime)
				}
				query1.Select("count(1) as betcount").Scan(&betcount)
				query1.Select("sum(net_amount) as win").Scan(&win)
				query1.Select("sum(valid_betamount) as betAmount").Scan(&betAmount)

				rec := vo.BetRecordStatisData{VenueCode: venueCode, VenueName: venueName, GameType: gameType, Win: win, BetAmount: betAmount, BetCount: betcount, IsSettled: 0}
				list = append(list, rec)

				if gameType == "esport" || gameType == "sport" || gameType == "lottery" {
					query2 := global.G_DB.WithContext(ctx).Model(dos.FcBetRecordUnsettled{})
					query2.Where("user_id = ? and venue_code = ? and game_type = ?", userInfo.UserId, venueCode, gameType)
					if timeType == -1 {
						if jsonp.DateStart != "" {
							query2.Where("bet_time >=?", jsonp.DateStart)
						}
						if jsonp.DateEnd != "" {
							query2.Where("bet_time <=?", jsonp.DateEnd)
						}
					} else {
						sTime, eTime := tool.GetDayRange(time.Now(), timeType)
						query2.Where("bet_time BETWEEN ? AND ?", sTime, eTime)
					}
					query2.Select("count(1) as betcount").Scan(&betcount2)
					query2.Select("sum(valid_betamount) as betAmount").Scan(&betAmount2)
					rec2 := vo.BetRecordStatisData{VenueCode: venueCode, VenueName: venueName, GameType: gameType, Win: 0, BetAmount: betAmount2, BetCount: betcount2, IsSettled: 1}
					list = append(list, rec2)
				}

			}
		}
		venuesCache.Set(keyData, list, 1*time.Minute)
		data := vo.BetRecordStatisResp{Data: list}
		response.SuccessJSON(c, data)
	} else {
		newData := resData.([]vo.BetRecordStatisData)
		data := vo.BetRecordStatisResp{Data: newData}
		response.SuccessJSON(c, data)
	}
}

func GameRecordUnsettled(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		vo.BetRecordUnsetteledReq
	}{}

	jsonp.PageQuery.PageNo = tool.Atoi(c.DefaultQuery("current", "1"))
	jsonp.PageQuery.PageSize = tool.Atoi(c.DefaultQuery("pageSize", "10"))
	jsonp.VenueCode = c.DefaultQuery("venue_code", "")
	jsonp.DateEnd = c.DefaultQuery("date_end", "")
	jsonp.DateStart = c.DefaultQuery("date_start", "")
	jsonp.GameType = c.DefaultQuery("game_type", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", "0"))
	timeType := tool.Atoi(c.DefaultQuery("time_type", "-1"))

	if jsonp.PageNo == 0 {
		jsonp.PageNo = 1
	}
	if jsonp.PageSize == 0 {
		jsonp.PageSize = 10
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	k1 := 0
	st := ""
	et := ""
	if timeType == -1 {
		if jsonp.DateStart != "" {
			if len(jsonp.DateStart) < 11 {
				jsonp.DateStart += " 00:00:00"
			}
		}
		if jsonp.DateEnd != "" {
			if len(jsonp.DateEnd) < 11 {
				jsonp.DateEnd += " 23:59:59"
			}
		}
	} else {
		k1 = 1
		sTime1, eTime1 := tool.GetDayRange(time.Now(), timeType)
		st = sTime1.Format(tool.TimeLayout)
		et = eTime1.Format(tool.TimeLayout)
	}

	keyData := fmt.Sprintf("venue_record_unsettled:%s:%s:%s:%v:%s:%s:%v:%v:%v", userInfo.UserId, jsonp.VenueCode, jsonp.GameType, timeType, jsonp.DateStart, jsonp.DateEnd, jsonp.PageNo, jsonp.PageSize, jsonp.Status)
	if k1 == 1 {
		keyData = fmt.Sprintf("venue_record_unsettled:%s:%s:%s:%v:%s:%s:%v:%v:%v", userInfo.UserId, jsonp.VenueCode, jsonp.GameType, timeType, st, et, jsonp.PageNo, jsonp.PageSize, jsonp.Status)
	}
	keyTotal := keyData + ":total"
	keyBetAmount := keyData + ":betamount"
	keyNetAmount := keyData + ":netamount"
	resData, ok := venuesCache.Get(keyData)

	//global.G_LOG.Infof("venue betrecord------------------------1:%v, %v, %v", keyData, ok, resData)
	//global.G_LOG.Infof("venue betrecord------------------------2:%v, %v", jsonp.PageSize, jsonp.PageNo)
	if !ok {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()
		clientType := c.GetHeader(enmus.CLIENT_TYPE_HEADER)
		if len(clientType) == 0 {
			clientType = enmus.H5
		}
		language := c.GetHeader(enmus.LANGUAGE_HEADER)
		if len(language) == 0 {
			language = "zh-CN"
		}

		err := validator.New().Struct(jsonp)
		if err != nil {
			response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
			return
		}

		var data []*dos.FcBetRecordUnsettled
		query := global.G_DB.WithContext(ctx).Model(&dos.FcBetRecordUnsettled{}).
			Where("user_id = ? ", userInfo.UserId).
			Order("bet_time desc")
		query1 := global.G_DB.WithContext(ctx).Model(&dos.FcBetRecordUnsettled{}).
			Where("user_id = ? ", userInfo.UserId)

		if timeType == -1 {
			if jsonp.DateStart != "" {
				query.Where("bet_time >=?", jsonp.DateStart)
				query1.Where("bet_time >=?", jsonp.DateStart)
			}

			if jsonp.DateEnd != "" {
				query.Where("bet_time <=?", jsonp.DateEnd)
				query1.Where("bet_time <=?", jsonp.DateEnd)
			}
		} else {
			//sTime, eTime := tool.GetDayRange(time.Now(), timeType)
			query.Where("bet_time BETWEEN ? AND ?", st, et)
			query1.Where("bet_time BETWEEN ? AND ?", st, et)
		}

		if jsonp.VenueCode != "" {
			query.Where("venue_code =?", jsonp.VenueCode)
			query1.Where("venue_code =?", jsonp.VenueCode)
		}

		if jsonp.GameType != "" {
			query.Where("game_type =?", jsonp.GameType)
			query1.Where("game_type =?", jsonp.GameType)
		}

		var totalBetAmount, totalNetAmount float64
		var total int64
		query.Count(&total)

		query1.Select("sum(bet_amount) as totalBetAmount").Scan(&totalBetAmount)
		query1.Select("sum(net_amount) as totalNetAmount").Scan(&totalNetAmount)
		query.Offset((jsonp.PageNo - 1) * jsonp.PageSize).Limit(jsonp.PageSize).
			Scan(&data)

		newData := make([]*vo.BetRecordUnsettledResp, 0, len(data))
		tool.JsonMapper(data, &newData)
		//global.G_LOG.Infof("venue betrecord------------------------5:%v", newData)

		res := map[string]interface{}{
			"list":           newData,
			"totalBetAmount": totalBetAmount,
			"totalNetAmount": totalNetAmount,
		}
		if newData == nil {
			res["list"] = []string{}
		}
		venuesCache.Set(keyData, newData, 1*time.Minute)
		venuesCache.Set(keyTotal, total, 1*time.Minute)
		venuesCache.Set(keyBetAmount, totalBetAmount, 1*time.Minute)
		venuesCache.Set(keyNetAmount, totalNetAmount, 1*time.Minute)
		response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, res)
	} else {
		total, ok1 := venuesCache.Get(keyTotal)
		betAmount, ok2 := venuesCache.Get(keyBetAmount)
		netAmount, ok3 := venuesCache.Get(keyNetAmount)
		//global.G_LOG.Infof("venue betrecord------------------------3:%v, %v, %v", total, betAmount, netAmount)
		//global.G_LOG.Infof("venue betrecord------------------------4:%v, %v, %v", ok1, ok2, ok3)
		//global.G_LOG.Infof("venue betrecord------------------------5:%v, %v, %v", resData, jsonp.PageNo, jsonp.PageSize)
		if !ok1 || !ok2 || !ok3 {
			total = 0
			betAmount = 0
			netAmount = 0
		}
		total1 := total.(int64)
		netAmount = netAmount.(float64)
		betAmount = betAmount.(float64)
		newData := resData.([]*vo.BetRecordUnsettledResp)
		res := map[string]interface{}{
			"list":           newData,
			"totalBetAmount": betAmount,
			"totalNetAmount": netAmount,
		}
		response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total1, res)
	}
}
