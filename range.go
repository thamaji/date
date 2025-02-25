package date

import (
	"iter"
)

// Range は、指定された開始日(start)から終了日(end)までの連続した日付を生成するイテレータを返します。
// イテレータには開始日(start)、終了日(end)を含みます。
// `start` が `end` より前の場合は昇順に、後の場合は降順に日付を生成します。
func Range(start, end Date) iter.Seq[Date] {
	return func(yield func(Date) bool) {
		for d := range RangeWithIndex(start, end) {
			if !yield(d) {
				return
			}
		}
	}
}

// RangeWithIndex は、指定された開始日(start)から終了日(end)までの日付と、そのインデックスを含むイテレータを返します。
// イテレータには開始日(start)、終了日(end)を含みます。
// `start` が `end` より前の場合は昇順に、後の場合は降順に日付を生成します。
func RangeWithIndex(start, end Date) iter.Seq2[Date, int] {
	return func(yield func(Date, int) bool) {
		if start.Equal(end) {
			yield(start, 0)
			return
		}
		if start.Before(end) {
			i := 0
			for d := start; !d.After(end); d = d.Add(0, 0, 1) {
				if !yield(d, i) {
					return
				}
				i++
			}
		} else {
			i := 0
			for d := start; !d.Before(end); d = d.Add(0, 0, -1) {
				if !yield(d, i) {
					return
				}
				i++
			}
		}
	}
}
