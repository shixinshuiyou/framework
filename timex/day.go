package timex

import "time"

type Day struct {
	time.Time
}

const (
	dayFormat = "2006-01-02"
)

func (t *Day) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+dayFormat+`"`, string(data), time.Local)
	*t = Day{now}
	return
}
