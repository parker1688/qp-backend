package tool

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ipipdotnet/ipdb-go"
)

var ipDb *ipdb.City

func InitIPDB(path string) {
	var err error
	ipDb, err = ipdb.NewCity(path)
	if err != nil {
		fmt.Printf("init city ipdb filePath=%s fail err: %v\n", path, err)
	}
}

type CityInfo struct {
	CountryName string `json:"country_name"` // 国家
	RegionName  string `json:"region_name"`  // 省会
	CityName    string `json:"city_name"`    // 城市
}

func ClientIP(c *gin.Context) string {
	ip := c.GetHeader("cf-connecting-ip")
	if len(ip) == 0 {
		ip = c.ClientIP()
	}
	return ip
}

// IP数据库地址地域信息
func IPCityInfo(ip string) (*CityInfo, error) {
	cityInfo, err := ipDb.FindInfo(ip, "CN")
	if err != nil {
		return nil, err
	}

	city := CityInfo{}
	city.CountryName = cityInfo.CountryName
	city.RegionName = cityInfo.RegionName
	city.CityName = cityInfo.CityName

	return &city, nil
}
