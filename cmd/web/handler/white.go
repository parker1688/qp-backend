package handler

import (
	"github.com/gin-gonic/gin"
)

//var _cache = cache.New(2*time.Minute, 30*time.Second)

func WhiteIPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.GetHeader("cf-connecting-ip")
		if len(ip) == 0 {
			ip = c.ClientIP()
		}
		if ip == "127.0.0.1" || ip == "localhost" {
			c.Next()
			return
		}
		/*if !isWhiteIP(ip) {
			response.FailErrDataJSON(c, ecode.NotWhiteIp, "非白名单IP", ip)
			c.Abort()
			return
		}*/
		c.Next()
	}
}

/*func isWhiteIP(ip string) bool {
	//whiteSettingValue, exits := _cache.Get("WhiteIP_Status")
	//if !exits {
	whiteSetting := modules.FindByKeyDictsDetailFirst(&dos.DictsDetail{
		DictsTag: "WhiteIP",
	})
	//_cache.SetDefault("WhiteIP_Status", whiteSetting.DictsValue)
	whiteSettingValue := whiteSetting.DictsValue
	//}
	if whiteSettingValue == "1" {
		//ipValue, ipExits := _cache.Get(ip)
		//if !ipExits {
		data := modules.FindByKeyWhiteIpFirst(&dos.WhiteIp{
			IpAddr: ip,
		})
		ipValue := len(data.Id) == 0
		//_cache.SetDefault(ip, ipValue)
		//}
		if ipValue {
			return false
		}
	}
	return true
}*/
