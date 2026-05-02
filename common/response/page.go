package response

const (
	DefaultPageNo   = 1
	DefaultPageSize = 10
	MaxPageSize     = 100
)

type PageQuery struct {
	PageSize int `json:"pageSize" binding:"required" form:"pageSize" uri:"pageSize"`
	PageNo   int `json:"current" binding:"required" form:"current" uri:"current"`
}

type PageTimeQuery struct {
	PageSize    int    `json:"pageSize" binding:"required" form:"pageSize" uri:"pageSize"`
	PageNo      int    `json:"current" binding:"required" form:"current" uri:"current"`
	StartAt     string `json:"startAt" binding:"required" form:"startAt" uri:"startAt"`                //创建开始时间
	EndAt       string `json:"endAt" binding:"required" form:"endAt" uri:"endAt"`                      //创建结束时间
	TimeType    int    `json:"time_type" form:"time_type" uri:"time_type"`                             //时间类型（昨天 0｜今天 1｜7天 7｜30天 30）
	LastStartAt string `json:"last_startAt" binding:"required" form:"last_startAt" uri:"last_startAt"` //最后开始时间
	LastEndAt   string `json:"last_endAt" binding:"required" form:"last_endAt" uri:"last_endAt"`       //最后结束时间

	IsFree string `json:"is_free" binding:"required" form:"is_free" uri:"is_free"` //玩家默认状态
}

func NormalizePage(pageNo, pageSize int) (int, int) {
	if pageNo < 1 {
		pageNo = DefaultPageNo
	}
	if pageSize < 1 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	return pageNo, pageSize
}

func NormalizePageQuery(query *PageQuery) {
	query.PageNo, query.PageSize = NormalizePage(query.PageNo, query.PageSize)
}

func NormalizePageTimeQuery(query *PageTimeQuery) {
	query.PageNo, query.PageSize = NormalizePage(query.PageNo, query.PageSize)
}
