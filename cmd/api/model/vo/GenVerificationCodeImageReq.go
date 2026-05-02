package vo

type GenVerificationCodeImageReq struct {
	Type string `json:"type"`
}

type GenVerificationCodeImageResp struct {
	VeryCodeRandom string `json:"veryCodeRandom"`
	Image          string `json:"image"`
}
