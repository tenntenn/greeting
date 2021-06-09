package greeting

import (
	"time"
)

// Do関数は挨拶文を返します.
// 以下のように指定した時刻によって返す文を変えます.
// 04:00-09:59: おはよう
// 10:00-16:59: こんにちは
// 17:00-03:59: こんばんは
func Do(t time.Time) string {
	switch h := t.Hour(); {
	case h >= 4 && h <= 9:
		return "おはよう"
	case h >= 10 && h <= 16:
		return "こんにちは"
	default:
		return "こんばんは"
	}
}


// DoNow関数は現在時刻に合わせた挨拶文を返します.
// 以下のように指定した時刻によって返す文を変えます.
// 04:00-09:59: おはよう
// 10:00-16:59: こんにちは
// 17:00-03:59: こんばんは
ffunc DoNow() string {
	t := time.Now()
	switch h := t.Hour(); {
	case h >= 4 && h <= 9:
		return "おはよう"
	case h >= 10 && h <= 16:
		return "こんにちは"
	default:
		return "こんばんは"
	}
}
