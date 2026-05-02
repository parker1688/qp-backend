package automaticType

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Time time.Time

const timeFormatter = "2006-01-02 15:04:05"
const timeFormatter_nosec = "2006-01-02 15:04"
const timeFormatter_date = "2006-01-02"

var nullTime = time.Date(1949, 1, 1, 1, 1, 1, 1, time.Local)

func Now() Time {
	return Time(time.Now())
}

func (t Time) Timer() time.Time {
	return time.Time(t)
}

func (t *Time) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	loc, err := time.LoadLocation("Local")
	if err != nil {
		loc = time.FixedZone("CST", 8*3600)
	}
	sv := string(b)
	if len(sv) == 10 {
		sv += " 00:00:00"
	} else if len(sv) == 16 {
		sv += ":00"
	}
	now, err := time.ParseInLocation(timeFormatter, string(b), loc)
	if err != nil {
		if now, err = time.ParseInLocation(timeFormatter_nosec, string(b), loc); err != nil {
			now, err = time.ParseInLocation(timeFormatter_date, string(b), loc)
		}
	}
	if err != nil {
		err = nil
		*t = Time(nullTime)
	} else {
		*t = Time(now)
	}
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.Timer().Before(nullTime) {
		return []byte("null"), nil
	}
	b := make([]byte, 0, len(timeFormatter)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormatter)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormatter)
}

// Value ...
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan valueof time.Time 注意是指针类型 method
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
