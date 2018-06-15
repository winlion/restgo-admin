package restgo

import (
	"time"
)


type JsonDateTime time.Time
type JsonDate time.Time
type JsonTime time.Time

const (
	dateTimeFormart = "2006-01-02 15:04:05"
	dateFormart = "2006-01-02"
	timeFormart = "15:04:05"
)

func (p *JsonDateTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+dateTimeFormart+`"`, string(data), time.Local)
	*p = JsonDateTime(now)
	return
}
func (p *JsonDate) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+dateFormart+`"`, string(data), time.Local)
	*p = JsonDate(now)
	return
}
func (p *JsonTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*p = JsonTime(now)
	return
}
func (c JsonDateTime) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0)
	data = append(data, '"')
	data = time.Time(c).AppendFormat(data, dateTimeFormart)
	data = append(data, '"')
	return data, nil
}
func (c JsonDate) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0)
	data = append(data, '"')
	data = time.Time(c).AppendFormat(data, dateFormart)
	data = append(data, '"')
	return data, nil
}
func (c JsonTime) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0)
	data = append(data, '"')
	data = time.Time(c).AppendFormat(data, timeFormart)
	data = append(data, '"')
	return data, nil
}
func (c JsonDateTime) String() string {
	return time.Time(c).Format(dateTimeFormart)
}

func (c JsonTime) String() string {
	return time.Time(c).Format(timeFormart)
}