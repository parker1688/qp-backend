package gameControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const (
	GAMEHistorySlots = "HistoryGAMESlots:%v" //历史游戏缓存
	GAMECollectSlots = "CollectGAMESlots:%v" //收藏游戏缓存
)

// GetGameSlots
//
//	@Description: 获取游戏场馆
//	@param c
func GetGameSlots(c *gin.Context) {
	gameType := c.DefaultQuery("game_type", "")
	venueCode := c.DefaultQuery("venue_code", "")
	venueType := c.DefaultQuery("venue_type", "")
	hot := tool.Atoi(c.DefaultQuery("hot", "0"))
	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = c.DefaultQuery("language", "")
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	var data []*dos.FcVenueGame
	query := global.G_DB.WithContext(ctx).Model(&dos.FcVenueGame{})
	query = query.Where("status in ?", []int{1, 3})
	if len(gameType) > 0 {
		query = query.Where("game_type = ?", gameType)
	}

	venueCode = strings.Trim(venueCode, "[")
	venueCode = strings.Trim(venueCode, "]")
	if len(venueCode) > 0 {
		query = query.Where("venue_code in ?", strings.Split(venueCode, ","))
	}
	if len(venueType) > 0 {
		query = query.Where("venue_type = ?", venueType)
	}

	if hot > 0 {
		query = query.Where("hot = ?", hot)
	}
	//if len(language) > 0 {
	//	query = query.Where("language = ?", language)
	//}
	query.Order("sort").Find(&data)
	newData := make([]*vo.GameSlotsResp, 0, len(data))
	tool.JsonMapper(data, &newData)

	response.SuccessJSON(c, newData)
}

// QueryGameSlots
//
//	@Description: 游戏场馆搜索
//	@param c
func QueryGameSlots(c *gin.Context) {
	gameType := c.DefaultQuery("game_type", "")
	gameName := c.DefaultQuery("game_name", "")
	venueCode := c.DefaultQuery("venue_code", "")
	venueType := c.DefaultQuery("venue_type", "")
	hot := tool.Atoi(c.DefaultQuery("hot", "0"))
	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = c.DefaultQuery("language", "")
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()
	var data []*dos.FcVenueGame
	query := global.G_DB.WithContext(ctx).Model(&dos.FcVenueGame{})
	query = query.Where("status in ?", []int{1, 3})
	if len(gameType) > 0 {
		query = query.Where("game_type = ?", gameType)
	}
	if len(venueCode) > 0 {
		query = query.Where("venue_code = ?", venueCode)
	}
	if len(venueType) > 0 {
		query = query.Where("venue_type = ?", venueType)
	}
	if len(gameName) > 0 {
		query = query.Where("game_name like ?", "%"+gameName+"%")
	}

	if hot > 0 {
		query = query.Where("hot = ?", hot)
	}
	//if len(language) > 0 {
	//	query = query.Where("language = ?", language)
	//}
	query.Order("hot desc, sort").Find(&data)
	newData := make([]*vo.GameSlotsResp, 0, len(data))
	tool.JsonMapper(data, &newData)

	response.SuccessJSON(c, newData)
}

// HistoryGame
//
//	@Description: 用户历史游戏
//	@param c
func HistoryGame(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	key := fmt.Sprintf(GAMEHistorySlots, userInfo.UserId)
	data, _ := global.G_REDIS.ZRevRange(context.Background(), key, 0, -1).Result()
	newData := make([]*vo.HistoryGameSlotsReq, 0, len(data))
	for _, v := range data {
		var m *vo.HistoryGameSlotsReq
		err := tool.JsonUnmarshalFromString(v, &m)
		if err == nil {
			newData = append(newData, m)
		}
	}
	response.SuccessJSON(c, newData)
}

// HistoryGameAdd
//
//	@Description: 用户历史游戏添加
//	@param c
func HistoryGameAdd(c *gin.Context) {
	var jsonp vo.HistoryGameSlotsReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	key := fmt.Sprintf(GAMEHistorySlots, userInfo.UserId)
	count := global.G_REDIS.ZCard(context.Background(), key).Val()
	if count > 15 {
		global.G_REDIS.ZRemRangeByRank(context.Background(), key, 0, 0)
	}
	global.G_REDIS.ZAdd(context.Background(), key, redis.Z{Score: float64(time.Now().Unix()), Member: tool.String(jsonp)})
	global.G_REDIS.Expire(context.Background(), key, 31*24*time.Hour)
	response.SuccessJSON(c, true)
}

// DelHistoryGame
//
//	@Description: 清空历史记录
//	@param c
func DelHistoryGame(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	key := fmt.Sprintf(GAMEHistorySlots, userInfo.UserId)
	count := global.G_REDIS.ZCard(context.Background(), key).Val()
	global.G_REDIS.ZRemRangeByRank(context.Background(), key, 0, count)
	response.SuccessJSON(c, true)
}

// CollectGameSlots
//
//	@Description: 收藏
//	@param c
func CollectGameSlots(c *gin.Context) {
	var jsonp vo.HistoryGameSlotsReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	key := fmt.Sprintf(GAMECollectSlots, userInfo.UserId)
	count := global.G_REDIS.ZCard(context.Background(), key).Val()
	if count >= 20 {
		global.G_REDIS.ZRemRangeByRank(context.Background(), key, 0, 0)
	}
	global.G_REDIS.ZAdd(context.Background(), key, redis.Z{Score: float64(time.Now().Unix()), Member: tool.String(jsonp)})
	global.G_REDIS.Expire(context.Background(), key, 31*24*time.Hour)
	response.SuccessJSON(c, true)
}

// GetCollectGameSlots
//
//	@Description: 获取收藏
//	@param c
func GetCollectGameSlots(c *gin.Context) {
	venueType := c.DefaultQuery("venue_type", "")
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	key := fmt.Sprintf(GAMECollectSlots, userInfo.UserId)
	data, _ := global.G_REDIS.ZRevRange(context.Background(), key, 0, -1).Result()
	newData := make([]*vo.HistoryGameSlotsReq, 0)
	for _, v := range data {
		var m *vo.HistoryGameSlotsReq
		err := tool.JsonUnmarshalFromString(v, &m)
		if err != nil {
			global.G_LOG.Errorf("GetCollectGameSlots username=%s JsonUnmarshal err: %v", userInfo.UserName, err)
			continue
		}

		// 场馆类型
		if venueType == "" {
			newData = append(newData, m)
		} else if m.VenueType == venueType {
			newData = append(newData, m)
		}
	}

	response.SuccessJSON(c, newData)
}

// DelCollectGameSlots
//
//	@Description: 删除收藏游戏
//	@param c
func DelCollectGameSlots(c *gin.Context) {
	var jsonp vo.HistoryGameSlotsReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if len(jsonp.Id) == 0 {
		response.SuccessJSON(c, true)
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	key := fmt.Sprintf(GAMECollectSlots, userInfo.UserId)
	data, _ := global.G_REDIS.ZRevRange(context.Background(), key, 0, -1).Result()
	for _, v := range data {
		var m vo.HistoryGameSlotsReq
		_ = tool.JsonUnmarshalFromString(v, &m)
		if m.Id == jsonp.Id {
			global.G_REDIS.ZRem(context.Background(), key, v)
			break
		}
	}
	response.SuccessJSON(c, true)
}
