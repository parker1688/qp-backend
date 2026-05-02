// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveGuide(vo *dos.Guide) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageGuide(page, pageSize int, vo *dos.Guide, c *gin.Context) (ret []*dos.Guide, total int64) {
	query := global.G_DB.Model(&dos.Guide{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Key) > 0 {
		query = query.Where("`key` like ?", "%"+vo.Key+"%")
	}

	if len(vo.Name) > 0 {
		query = query.Where("name like ?", "%"+vo.Name+"%")
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.Guide
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyGuide(vo *dos.Guide, c *gin.Context) []*dos.Guide {
	var data []*dos.Guide
	query := global.G_DB.Model(&dos.Guide{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Key) > 0 {
		query = query.Where("`key` like ?", "%"+vo.Key+"%")
	}

	if len(vo.Name) > 0 {
		query = query.Where("name like ?", "%"+vo.Name+"%")
	}

	query.Find(&data)
	return data
}

func FindByKeyGuideFirst(vo *dos.Guide) *dos.Guide {
	var data *dos.Guide
	query := global.G_DB.Model(&dos.Guide{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Key) > 0 {
		query = query.Where("`key` like ?", "%"+vo.Key+"%")
	}

	if len(vo.Name) > 0 {
		query = query.Where("name like ?", "%"+vo.Name+"%")
	}

	query.Take(&data)
	return data
}

func UpdateGuide(vo *dos.Guide) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"key":         vo.Key,
		"name":        vo.Name,
		"data":        vo.Data,
		"update_time": automaticType.Time(time.Now()),
		"update_by":   vo.UpdateBy,
	}).Error == nil
}

func DeleteGuide(vo *dos.Guide) bool {
	return global.G_DB.Model(&dos.Guide{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// GetGuideInfo - 获取文本导航信息
// @param {string} guideId
// @returns []dos.GuideInfoResp
func GetGuideInfo(guideId string) []dos.GuideInfoResp {
	data := []dos.GuideInfoResp{}
	query := global.G_DB.Model(&dos.Guide{}).Select("id", "key", "name", "data")
	if len(guideId) > 0 {
		query = query.Where("id = ?", guideId)
	}
	err := query.Find(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetGuideInfo] Find guide data failed: %v", err.Error())
		return data
	}

	return data
}

// GetGuideDetails - 获取文本导航内容
func GetGuideDetails(guideId string) string {
	data := ""
	err := global.G_DB.Model(&dos.Guide{}).Select("data").Where("id = ?", guideId).Scan(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetGuideDetails] Find guide details failed: id=%s, err=%v", guideId, err.Error())
		return ""
	}

	if len(data) > 0 {
		var navData struct {
			//NavList interface{} `json:"navList"`
			Content string `json:"content"`
		}
		err = tool.JsonUnmarshalFromString(data, &navData)
		if err != nil {
			global.G_LOG.Errorf("[GetGuideDetails] Json unmarshal from string failed: id=%s, data=%s, err=%v",
				guideId, data, err.Error())
			return ""
		}

		return navData.Content
	} else {
		global.G_LOG.Errorf("[GetGuideDetails] Data is empty: id=%s", guideId)
	}

	return ""
}
