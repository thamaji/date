package date

import (
	"errors"
	"time"
)

// DefaultLayout は日付のデフォルトフォーマットを定義します。
const DefaultLayout = "2006-01-02"

// Parse は指定されたフォーマットで文字列を解析し、Date を返します。
func Parse(layout, value string) (Date, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return Date{}, err
	}
	return FromTime(t), nil
}

// Date は時刻情報を持たない日付を表します。
type Date struct {
	time time.Time
	loc  *time.Location
}

// Now は現在の日付を返します。
func Now() Date {
	return FromTime(time.Now())
}

// FromTime は time.Time から Date を生成します。
func FromTime(t time.Time) Date {
	year, month, day := t.Date()
	return Date{time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC), loc: t.Location()}
}

// New は指定した年月日で Date を作成します。
func New(year int, month time.Month, day int) Date {
	return Date{time: time.Date(year, month, day, 0, 0, 0, 0, time.UTC), loc: time.Local}
}

// IsZero は Date がゼロ値かどうかを判定します。
func (d Date) IsZero() bool {
	return d.loc == nil && d.time.IsZero()
}

// Add は指定された年、月、日を加算した新しい Date を返します。
func (d Date) Add(years int, months int, days int) Date {
	return Date{time: d.time.AddDate(years, months, days), loc: d.loc}
}

// Sub は2つの日付の差を日数で返します。
func (d Date) Sub(date Date) int {
	return int(d.time.Sub(date.time) / (24 * time.Hour))
}

// Since は指定された日付から現在までの日数を返します。
func Since(date Date) int {
	return Now().Sub(date)
}

// Until は現在から指定された日付までの日数を返します。
func Until(date Date) int {
	return date.Sub(Now())
}

// Before は d が指定した日付より前かどうかを判定します。
func (d Date) Before(date Date) bool {
	return d.time.Before(date.time)
}

// After は d が指定した日付より後かどうかを判定します。
func (d Date) After(date Date) bool {
	return d.time.After(date.time)
}

// Compare は d と指定された日付を比較し、-1, 0, 1 を返します。
func (d Date) Compare(date Date) int {
	switch {
	case d.Before(date):
		return -1
	case d.After(date):
		return +1
	}
	return 0
}

// Equal は d と指定された日付が等しいかを判定します。
func (d Date) Equal(date Date) bool {
	return d.time.Equal(date.time)
}

// Time は Date を time.Time に変換します。
func (d Date) Time() time.Time {
	year, month, day := d.time.Date()
	loc := d.loc
	if loc == nil {
		loc = time.Local
	}
	return time.Date(year, month, day, 0, 0, 0, 0, loc)
}

// TimeInLocation は、指定されたタイムゾーンの 0 時間 0 分 0 秒の time.Time を返します。
func (d Date) TimeInLocation(loc *time.Location) time.Time {
	year, month, day := d.time.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, loc)
}

// Weekday は、日付の曜日を返します。
func (d Date) Weekday() time.Weekday {
	return d.time.Weekday()
}

// Year は、西暦年を返します。
func (d Date) Year() int {
	return d.time.Year()
}

// Month は、月を返します。
func (d Date) Month() time.Month {
	return d.time.Month()
}

// Day は、日を返します。
func (d Date) Day() int {
	return d.time.Day()
}

// YMD は、西暦年、月、日をそれぞれ返します。
func (d Date) YMD() (int, time.Month, int) {
	return d.time.Date()
}

// ISOWeek は、ISO 8601 形式の年と週番号を返します。
func (d Date) ISOWeek() (int, int) {
	return d.time.ISOWeek()
}

// YearDay は、その年の通算日（1月1日を1とする）を返します。
func (d Date) YearDay() int {
	return d.time.YearDay()
}

// SetYear は、西暦年を設定します。
func (d *Date) SetYear(year int) {
	_, month, day := d.time.Date()
	d.time = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// SetMonth は、月を設定します。
func (d *Date) SetMonth(month time.Month) {
	year, _, day := d.time.Date()
	d.time = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// SetDay は、日を設定します。
func (d *Date) SetDay(day int) {
	year, month, _ := d.time.Date()
	d.time = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// MarshalBinary は、日付をバイナリ形式にエンコードします。
func (d Date) MarshalBinary() ([]byte, error) {
	return d.Time().MarshalBinary()
}

// UnmarshalBinary は、バイナリデータから日付をデコードします。
func (d *Date) UnmarshalBinary(data []byte) error {
	t := time.Time{}
	if err := t.UnmarshalBinary(data); err != nil {
		return err
	}
	*d = FromTime(t)
	return nil
}

// GobEncode は、日付をバイナリ形式にエンコードします。
func (d Date) GobEncode() ([]byte, error) {
	return d.MarshalBinary()
}

// GobDecode は、バイナリデータから日付をデコードします。
func (d *Date) GobDecode(data []byte) error {
	return d.UnmarshalBinary(data)
}

// MarshalJSON は、日付を JSON 文字列にエンコードします。
func (d Date) MarshalJSON() ([]byte, error) {
	if y := d.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DefaultLayout)+2)
	b = append(b, '"')
	b = d.AppendFormat(b, DefaultLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON は、JSON 文字列から日付をデコードします。
func (d *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	var err error
	*d, err = Parse(`"`+DefaultLayout+`"`, string(data))
	return err
}

// MarshalText は、日付をテキスト形式にエンコードします。
func (d Date) MarshalText() ([]byte, error) {
	if y := d.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalText: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DefaultLayout))
	return d.AppendFormat(b, DefaultLayout), nil
}

// UnmarshalText は、テキストから日付をデコードします。
func (d *Date) UnmarshalText(data []byte) error {
	var err error
	*d, err = Parse(DefaultLayout, string(data))
	return err
}

// String は、日付をデフォルトのフォーマットで文字列に変換します。
func (d Date) String() string {
	return d.Time().Format(DefaultLayout)
}

// Format は、指定されたレイアウトで日付をフォーマットします。
func (d Date) Format(layout string) string {
	return d.Time().Format(layout)
}

// AppendFormat は、指定されたレイアウトでフォーマットされた日付をバイトスライスに追加します。
func (d Date) AppendFormat(b []byte, layout string) []byte {
	return d.Time().AppendFormat(b, layout)
}
