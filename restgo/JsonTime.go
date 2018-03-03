package restgo

import (
	"time"
)

type JsonTime time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = JsonTime(now)
	return
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t JsonTime) String() string {
	return time.Time(t).Format(timeFormart)
}