package tool

import (
	"bootpkg/common/expands/automaticType"
	"strconv"
	"strings"
	"time"
)

const (
	TimeLayout              = "2006-01-02 15:04:05"
	TimeZeroLayout          = "2006-01-02 00:00:00"
	TimeDateLayout          = "2006-01-02"
	TimeDateYearMonLayout   = "2006-01"
	TimeDateLayoutCN        = "2006年1月02日"
	TimeYMDHMSNoSpaceLayout = "20060102150405"
)

func StrToTimeZero(s string, args ...string) time.Time {
	var format string
	if len(s) > 10 {
		format = "2006-01-02 15:04:05"
	} else {
		format = "2006-01-02"
	}

	if len(args) > 0 {
		format = strings.Trim(args[0], " ")
	}
	if len(s) != len(format) {
		var zeroTime time.Time
		return zeroTime
	}
	ti, _ := time.ParseInLocation(format, s, time.Local)
	return ti
}

// 时间转时间戳 timeLayout := "2006-01-02 15:04:05" 转化所需模板
func DateToTimeStamp(timeLayout, toBeCharge string) int64 {
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	s := theTime.Unix()
	return s
}

func DateToTime(timeLayout, toBeCharge string) *time.Time {

	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	return &theTime
}

// 时间戳转时间
func TimeStampToDate(timeLayout string, timeStamp int64) string {
	//时间戳转日期
	return time.Unix(timeStamp, 0).Format(timeLayout)
}

func TodayStartEndDate() (start, end string) {
	timeStr := time.Now().Format("2006-01-02")
	return timeStr + " 00:00:00", timeStr + " 23:59:59"
	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	//t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	//t2, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	//todayStartTime := t.Unix() + 1
	//todayEndTime := t2.AddDate(0, 0, 1).Unix()
	//return TimeStampToDate(TimeLayout, todayStartTime), TimeStampToDate(TimeLayout, todayEndTime)
}

func DayStartEndDate(day time.Time) (start, end string) {
	timeStr := day.Format("2006-01-02")
	return timeStr + " 00:00:00", timeStr + " 23:59:59"
	//使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	//t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	//t2, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	//todayStartTime := t.Unix() + 1
	//todayEndTime := t2.AddDate(0, 0, 1).Unix()
	//return TimeStampToDate(TimeLayout, todayStartTime), TimeStampToDate(TimeLayout, todayEndTime)
}

// TimeTomorrowTime
//
//	@Description: 获取现在与明日相差时间
func TimeTomorrowTime() time.Duration {
	now := time.Now()
	timeStr := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	tomorrow, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	return tomorrow.Sub(now)
}

// TimeRangeCurrentTime
//
//	@Description: 计算time1是否在开始和结束时间之间
//	@param time1 需要计算时间
//	@param start 开始时间
//	@param end 结束时间
//	@return bool true 在  false不在
func TimeRangeCurrentTime(time1 time.Time, start time.Time, end time.Time) bool {
	if start.Year() < 2008 || end.Year() < 2008 {
		return false
	}
	if time1.After(start) && time1.Before(end) {
		return true
	}
	return false
}

func TimeNowString() string {
	return time.Now().Format(TimeLayout)
}

func TimeNowTimestampString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func TimeNowYMDHMSNoSpaceString() string {
	return time.Now().Format(TimeYMDHMSNoSpaceLayout)
}

// 获取毫米级别时间字符串
func TimeNowYMDHMSMNoSpaceString() string {
	nowTime := time.Now()
	// 获得毫秒级别秒数
	nowS := nowTime.Unix() * 1000
	nowMs := nowTime.UnixNano() / 1e6
	subMs := nowMs - nowS

	nowMsStr := nowTime.Format(TimeYMDHMSNoSpaceLayout) + strconv.FormatInt(subMs, 10)
	return nowMsStr
}

// 获取本月多少周
func GetDateWeek(dateStr string) (int, int, int) {
	//dateStr := "2019-12-31"
	date, _ := time.Parse("2006-01-02", dateStr)
	// 获取当前时间数据当前第几周
	_, week := date.ISOWeek()

	weekday := date.Weekday()
	if weekday == 0 {
		weekday = 7
	}
	// 获取日期所在周的周四日期
	thursday := date.AddDate(0, 0, int(4-weekday))
	// 所属月
	month := int(thursday.Month())
	// 所属年
	year := thursday.Year()
	// 获取日期所在周的周四所在的月是第几周
	_, week1 := time.Date(thursday.Year(), thursday.Month(), 4, 0, 0, 0, 0, time.Local).ISOWeek()
	// 日期所在周数减去日期所在那个周的周四的月份的4号所在的周数再加一即为本月第几周
	week = week - week1 + 1
	return year, month, week
	//fmt.Println(year, month, week)
}

// 获取截止时间
func GetDayRange(dt time.Time, days int) (sTime time.Time, eTime time.Time) {
	if days > 0 {
		sTime = time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
		eTime = sTime.AddDate(0, 0, days).Add(-time.Second)
	} else if days == 0 {
		sTime = time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
		eTime = sTime.AddDate(0, 0, days).Add(-time.Second)
		sTime = time.Date(eTime.Year(), eTime.Month(), eTime.Day(), 0, 0, 0, 0, dt.Location())
	} else {
		eTime = time.Date(dt.Year(), dt.Month(), dt.Day(), 23, 59, 59, 0, dt.Location())
		sTime = eTime.AddDate(0, 0, days).Add(+time.Second)
	}
	return sTime, eTime
}

// 获取小时数时间戳
func HourTimeStamp() int64 {
	// 获取当前时间
	now := time.Now()
	// 获取整点时间(去除分钟、秒和纳秒)
	hourly := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	// 转换为时间戳
	timestamp := hourly.Unix()
	return timestamp
}

// GetDistanceDay - 获取距离时间
// @param {time.Time} dt
// @param {int} days
// @param {[3]int} timeFmt
// @returns time.Time
func GetDistanceDay(dt time.Time, days int, timeFmt [3]int) time.Time {
	sTime := time.Date(dt.Year(), dt.Month(), dt.Day(), timeFmt[0], timeFmt[1], timeFmt[2], 0, dt.Location())
	return sTime.AddDate(0, 0, days)
}

// GetTimeFromString - 根据字符串获取时间类型
// @param {string} t
// @returns time.Time
func GetTimeFromString(t string) time.Time {
	/*tm, _ := time.Parse(TimeLayout, t)
	return tm*/
	tm, _ := time.ParseInLocation(TimeLayout, t,
		time.FixedZone("CST", 8*60*60))
	return tm
}

// GetRangeDateList - 获取范围时间列表
// @param {string} stime 开始时间
// @param {string} etime 结束时间
// @param {[3]int} timeFmt 时间格式(时分秒)
// @returns []time.Time
func GetRangeDateList(stime, etime string, timeFmt [3]int) []time.Time {
	var dates []time.Time

	start := GetTimeFromString(stime)
	end := GetTimeFromString(etime)

	now := time.Now()
	start = time.Date(start.Year(), start.Month(), start.Day(), timeFmt[0], timeFmt[1], timeFmt[2], 0, now.Location())
	end = time.Date(end.Year(), end.Month(), end.Day(), timeFmt[0], timeFmt[1], timeFmt[2], 0, now.Location())

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d)
	}

	return dates
}

// CovertAutomaticTypeTimeFromDateString - 转换时间字符串为自动类型时间
// @param {string} date 时间字符串
// @returns *automaticType.Time
func CovertAutomaticTypeTimeFromDateString(date string) *automaticType.Time {
	if len(date) == 0 {
		return nil
	}

	d := new(automaticType.Time)
	if err := d.UnmarshalJSON([]byte(date)); err != nil {
		return nil
	}

	return d
}

// CovertDateStringFromAutomaticTypeTime - 转换自动类型时间为时间字符串
// @param {*automaticType.Time} t
// @returns string
func CovertDateStringFromAutomaticTypeTime(t *automaticType.Time) string {
	if t == nil {
		return ""
	}

	return t.String()
}

// CovertTimestampFromAutomaticTypeTime - 转换自动类型时间为时间戳
// @param {*automaticType.Time} t
// @returns int64
func CovertTimestampFromAutomaticTypeTime(t *automaticType.Time) int64 {
	if t == nil {
		return 0
	}

	return t.Timer().Unix()
}

// IsDifferentWeekCustom - 判断是否是隔周
func IsDifferentWeekCustom(t1, t2 time.Time, startDay time.Weekday) bool {
	getWeekStart := func(t time.Time, startDay time.Weekday) time.Time {
		// 计算到本周开始日的时间差
		daysDiff := (int(t.Weekday()) - int(startDay) + 7) % 7
		return t.AddDate(0, 0, -daysDiff).Truncate(24 * time.Hour)
	}

	// 计算两个时间所在周的起始时间
	weekStart1 := getWeekStart(t1, startDay)
	weekStart2 := getWeekStart(t2, startDay)

	// 比较周的起始时间是否相同
	return !weekStart1.Equal(weekStart2)
}
