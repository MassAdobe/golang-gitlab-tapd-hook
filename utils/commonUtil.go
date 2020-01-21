package utils

import "time"

const (
	TimeFormatMS    = "2006-01-02 15:04:05"
	TimeFormatMonth = "2006-01-02"
)

func GetTime() (timsStr string) {
	timsStr = time.Now().Format(TimeFormatMS)
	return
}

func GetDate() (timsStr string) {
	timsStr = time.Now().Format(TimeFormatMonth)
	return
}

func GetTimeFromUnix(timeInt64 int64) string {
	tm := time.Unix(timeInt64, 0)
	return tm.Format(TimeFormatMS)
}

// 比较两个时间大小 第一个早于第二 return true
func CompareTm(firstTm, secondTm string) bool {
	date1, _ := time.Parse(TimeFormatMS, firstTm)
	date2, _ := time.Parse(TimeFormatMS, secondTm)
	return date1.Before(date2)
}

// 比较日期是否和当前日期前一天相等
func CompareTmSame(tm string) bool {
	tim, _ := time.Parse(TimeFormatMonth, tm)
	tm = tim.Format(TimeFormatMonth)
	if GetDate() == tm {
		return true
	}
	return false
}
