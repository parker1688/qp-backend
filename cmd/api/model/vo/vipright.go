package vo

type FcVipRightsReq struct {
	Ids          []interface{} `json:"ids"`
	Level        int           `json:"level"`
	RightsName   []string      `json:"rights_name"`
	MerchantCode string        `json:"merchant_code"`
}
