package vo

type AnalysisRetentionResp struct {
	Date   string `json:"date"`
	RegNum string `json:"regnum"`
	Day2   string `json:"day2"`
	Day3   string `json:"day3"`
	Day7   string `json:"day7"`
	Day14  string `json:"day14"`
	Day30  string `json:"day30"`
	Day60  string `json:"day60"`
	Day90  string `json:"day90"`
}
