package model

import "time"

const (
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05"
	DateTimeFormat = "2006-01-02 15:04:05"
)

type Time struct {
	time.Time
	pattern string
}

func Now() Time {
	return Time{
		Time:    time.Now(),
		pattern: DateTimeFormat,
	}
}

func (t *Time) SetPattern(pattern string) {
	t.pattern = pattern
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+t.pattern+`"`, string(data), time.Local)
	*t = Time{
		Time:    now,
		pattern: t.pattern,
	}
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(t.pattern)+2)
	b = append(b, '"')
	b = t.Time.AppendFormat(b, t.pattern)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return t.Time.Format(t.pattern)
}
