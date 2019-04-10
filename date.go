package date

import (
	"errors"
	"time"
)

const DefaultLayout = "2006-01-02"

func Parse(layout, value string) (Date, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return Date{}, err
	}
	return FromTime(t), nil
}

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

func New(year int, month time.Month, day int) Date {
	return Date{time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC), loc: time.Local}
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

func (d Date) ISOWeek() (int, int) {
	return d.time.ISOWeek()
}

func (d Date) MarshalJSON() ([]byte, error) {
	if y := d.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DefaultLayout)+2)
	b = append(b, '"')
	b = d.time.AppendFormat(b, DefaultLayout)
	b = append(b, '"')
	return b, nil
}

func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	t, err := time.Parse(`"`+DefaultLayout+`"`, string(data))
	if err != nil {
		return err
	}

	year, month, day := t.Date()
	d.time = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	d.loc = t.Location()
	return nil
}

func (d Date) MarshalText() ([]byte, error) {
	if y := d.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalText: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DefaultLayout))
	return d.time.AppendFormat(b, DefaultLayout), nil
}

func (d *Date) UnmarshalText(data []byte) error {
	t, err := time.Parse(DefaultLayout, string(data))
	if err != nil {
		return err
	}

	year, month, day := t.Date()
	d.time = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	d.loc = t.Location()
	return nil
}

func (d Date) String() string {
	return d.Time().Format(DefaultLayout)
}

func (d Date) Format(layout string) string {
	return d.Time().Format(layout)
}
