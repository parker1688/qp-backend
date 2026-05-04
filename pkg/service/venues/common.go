package venues

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/venues/venueDetail"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

var venuesCache = cache.New(120*time.Second, 120*time.Second)

func checkAndFlushVenuesCache() {
	if global.G_REDIS.Get(context.Background(), "VenuesSyncFlag").Val() == "1" {
		venuesCache.Flush()
		global.G_REDIS.Set(context.Background(), "VenuesSyncFlag", "0", -1)
	}
}

const (
	venuesCacheLine           = "venuesCache_Line:%s"             //场馆在线
	venuesCodeUniqueCacheLine = "venuesCodeUniqueCacheLine:%s"    //场馆在线
	venuesCacheLineGameType   = "venuesCache_Line_GameType:%s:%s" //场馆在线
	venuesCacheStatus         = "venuesCache_Status:%s:%s"        //场馆维护缓存
	venuesAllCacheStatus      = "venuesCache_ALL_Status:%s"       //获取全部场馆缓存
	venuesCacheType           = "venuesCache_Type:%s"             //获取场馆类别
)

// GetVenuesLine
//
//	@Description: 获取商户场馆线路
//	@param merchantCode
//	@return []*dos.FcMerchantVenue
func GetVenuesLine(merchantCode string) []*dos.FcMerchantVenue {
	checkAndFlushVenuesCache()
	ojbData, ok := venuesCache.Get(fmt.Sprintf(venuesCacheLine, merchantCode))
	var data []*dos.FcMerchantVenue
	if !ok {
		//全部回收
		query := global.G_DB.Model(&dos.FcMerchantVenue{})
		query = query.Where("merchant_code = ? and status in ?", merchantCode, []int{1, 3})
		query.Find(&data)
		venuesCache.SetDefault(fmt.Sprintf(venuesCacheLine, merchantCode), data)
	} else {
		data = ojbData.([]*dos.FcMerchantVenue)
	}
	return data
}

// GetVenueCodeUniqueVenueLine
//
//	@Description: 获取商户场馆所有场馆，venue_code 唯一
//	@param merchantCode
//	@return []*dos.FcMerchantVenue
func GetVenueCodeUniqueVenueLine(merchantCode string) []*dos.FcMerchantVenue {
	checkAndFlushVenuesCache()
	ojbData, ok := venuesCache.Get(fmt.Sprintf(venuesCodeUniqueCacheLine, merchantCode))
	var data []*dos.FcMerchantVenue
	if !ok {
		tmpVenueArr := []*dos.FcMerchantVenue{}
		tmpVenueMap := make(map[string]bool)

		//全部回收
		query := global.G_DB.Model(&dos.FcMerchantVenue{})
		query = query.Where("merchant_code = ? and status in ?", merchantCode, []int{1, 3})
		query.Find(&tmpVenueArr)
		for i := range tmpVenueArr {
			tmpData := tmpVenueArr[i]
			_, ok := tmpVenueMap[tmpData.VenueCode]
			if !ok {
				tmpVenueMap[tmpData.VenueCode] = true
				data = append(data, tmpData)
			}
		}

		venuesCache.SetDefault(fmt.Sprintf(venuesCodeUniqueCacheLine, merchantCode), data)
	} else {
		data = ojbData.([]*dos.FcMerchantVenue)
	}
	return data
}

// GetVenuesLine
//
//	@Description: 获取商户场馆线路
//	@param merchantCode
//	@return []*dos.FcMerchantVenue
func GetVenuesLineWithGameType(merchantCode, gameType string) []*dos.FcMerchantVenue {
	checkAndFlushVenuesCache()

	ojbData, ok := venuesCache.Get(fmt.Sprintf(venuesCacheLineGameType, merchantCode, gameType))
	var data []*dos.FcMerchantVenue
	if !ok {
		//全部回收，仅返回在 fc_venue（gameStadium）中已配置的场馆
		query := global.G_DB.Model(&dos.FcMerchantVenue{}).
			Where("venue_code IN (SELECT venue_code FROM fc_venue)")
		if gameType != "" {
			query = query.Where("merchant_code = ? and status in ? AND game_type like ?", merchantCode, []int{1, 3}, "%"+gameType+"%")

		} else {
			query = query.Where("merchant_code = ? and status in ? ", merchantCode, []int{1, 3})
		}
		query.Find(&data)
		venuesCache.SetDefault(fmt.Sprintf(venuesCacheLineGameType, merchantCode, gameType), data)
	} else {
		data = ojbData.([]*dos.FcMerchantVenue)
	}
	return data
}

// GetVenuesStatus
//
//	@Description: 获取场馆状态
//	@param merchantCode 商户Code
//	@param venueCode 场馆Code
//	@return []*dos.FcMerchantVenue
func GetVenuesStatus(merchantCode string, venueCode string) *dos.FcMerchantVenue {
	checkAndFlushVenuesCache()
	ojbData, ok := venuesCache.Get(fmt.Sprintf(venuesCacheStatus, merchantCode, venueCode))
	var data *dos.FcMerchantVenue
	if !ok {
		data = modules.FindByKeyFcMerchantVenueFirst(&dos.FcMerchantVenue{
			MerchantCode: merchantCode,
			VenueCode:    venueCode,
		})
		venuesCache.SetDefault(fmt.Sprintf(venuesCacheStatus, merchantCode, venueCode), data)
	} else {
		data = ojbData.(*dos.FcMerchantVenue)
	}
	return data
}

// GetVenuesStatus
//
//	@Description: 获取场馆状态
//	@param merchantCode 商户Code
//	@param venueCode 场馆Code
//	@return []*dos.FcVenue
func GetVenuesType(venueCode string) *dos.FcVenue {
	checkAndFlushVenuesCache()
	ojbData, ok := venuesCache.Get(fmt.Sprintf(venuesCacheType, venueCode))
	var data *dos.FcVenue
	if !ok {
		data = modules.FindByKeyFcVenueFirst(&dos.FcVenue{
			VenueCode: venueCode,
		})
		venuesCache.SetDefault(fmt.Sprintf(venuesCacheType, venueCode), data)
	} else {
		data = ojbData.(*dos.FcVenue)
	}
	return data
}

// 獲取皇冠體育 token
func GetAndCreateToken(venueCode string) (string, error) {
	// 從緩存中獲取皇冠體育 token
	tokenKey := fmt.Sprintf(enmus.VENUE_LOGIN_TOKEN_KEY, venueCode)
	token := global.G_REDIS.Get(context.Background(), tokenKey).Val()
	if token != "" {
		return token, nil
	}

	return CreateToken(venueCode)
}

// 獲取皇冠體育 token
func CreateToken(venueCode string) (string, error) {
	token := ""

	venueResp := &vo.VenueLoginResponse{}
	if venueCode == enmus.HGTY {
		venueResp = venueDetail.HGTYLogin(global.CONFIG.Venue.HGTY)
		if venueResp == nil {
			global.G_LOG.Errorf("get hgty token fail")
			return "", fmt.Errorf("hgty return nil")
		}
	}

	if venueResp.Code != vo.Login_SUCCESS_CODE {
		return "", errors.New(venueResp.Msg)
	}
	if venueResp.Data.Token == "" {
		tmpStr := fmt.Sprintf("venueCode: %s getToken is empty %s", venueCode, venueResp.Msg)
		global.G_LOG.Errorf(tmpStr)
		return "", errors.New(tmpStr)
	}

	token = venueResp.Data.Token
	//global.G_LOG.Infof("get token success token: %s", token)

	ttl := 2 * time.Hour
	if venueCode == enmus.HGTY {
		ttl = 35 * time.Hour
	}

	// 設置 token
	tokenKey := fmt.Sprintf(enmus.VENUE_LOGIN_TOKEN_KEY, venueCode)
	global.G_REDIS.Set(context.Background(), tokenKey, token, ttl).Val()

	return token, nil
}

func GetOrderSn(req *vo.VenueGetOrderSnReq) string {
	switch req.VenueCode {
	case enmus.LYQP:
		return venueDetail.GetLYQPOrderSn(global.CONFIG.Venue.LYQP, req)
	}

	return ""
}

// 获取回放录像
func GetVenuePlayback(req *vo.VenuePlaybackReq) (*vo.VenuePlaybackResp, error) {
	resp := &vo.VenuePlaybackResp{}

	switch req.VenueCode {
	case enmus.TYQP:
		if req.TableId == "" {
			return nil, errors.New("TYQP getPlayback tableId is empty")
		}
		resp = venueDetail.GetTYQPPlayback(req)
	default:
		return nil, fmt.Errorf("Not support the venueCode %s playback", req.VenueCode)
	}

	if resp == nil {
		return nil, fmt.Errorf("VenuePlaybackResp is nil")
	}
	if resp.Code != vo.Login_SUCCESS_CODE {
		return nil, errors.New(resp.Msg)
	}

	return resp, nil
}
