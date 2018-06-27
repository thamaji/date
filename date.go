package date

import "time"

type Date struct {
	time time.Time
	loc  *time.Location
}

func Now() Date {
	return FromTime(time.Now())
}

func FromTime(t time.Time) Date {
	year, month, day := t.Date()
	return Date{time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC), loc: t.Location()}
}

func New(year, month, day int) Date {
	return Date{time: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), loc: time.Local}
}

func (d Date) Add(years int, months int, days int) Date {
	return Date{time: d.time.AddDate(years, months, days), loc: d.loc}
}

func (d Date) Sub(date Date) int {
	return int(d.time.Sub(date.time) / (24 * time.Hour))
}

func (d Date) Before(date Date) bool {
	return d.time.Before(date.time)
}

func (d Date) After(date Date) bool {
	return d.time.After(date.time)
}

func (d Date) Equal(date Date) bool {
	return d.time.Equal(date.time)
}

func (d Date) Time() time.Time {
	year, month, day := d.time.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, d.loc)
}

func (d Date) TimeInLocation(loc *time.Location) time.Time {
	year, month, day := d.time.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, loc)
}

func (d Date) Weekday() time.Weekday {
	return d.time.Weekday()
}

func (d Date) Year() int {
	return d.time.Year()
}

func (d Date) Month() time.Month {
	return d.time.Month()
}

func (d Date) Day() int {
	return d.time.Day()
}

func (d Date) YMD() (int, time.Month, int) {
	return d.time.Date()
}

func (d Date) String() string {
	return d.Time().Format("2006-01-02")
}
