package tsTime

import (
	"time"
	"tsEngine/tsString"
)

func MakeCurrTime() time.Time {
	return time.Now()
}

func MakeTimeSe(se int) time.Time {
	return time.Unix(int64(se), 0)
}

// 毫秒时间
func CurrMs() uint64 {
	curr_ms := time.Now().UnixNano()
	curr_ms = curr_ms / 1000000
	return uint64(curr_ms)
}

func CurrMsToString() string {
	return tsString.FromInt64(int64(CurrMs()))
}

// 妙计时间
func CurrSe() uint64 {
	curr := time.Now().Unix()
	return uint64(curr)
}

func CurrSeToString() string {
	return tsString.FromInt64(int64(CurrSe()))
}

// 20060102-150405; 20060102 15:04:05
func CurrSeFormat(format string) string {
	return time.Now().Format(format)
}

// 20060102-150405; 20060102 15:04:05
func CurrSeUtcFormat(format string) string {
	return time.Now().UTC().Format(format)
}

// 20060102-150405.000; 20060102 15:04:05.000
func CurrMsFormat(format string) string {
	return time.Now().Format(format)
}

func CurrDayBeginSe(hour, min, sec int) uint64 {
	curr := time.Now()
	day := time.Date(curr.Year(), curr.Month(), curr.Day(), hour, min, sec, 0, time.Local)
	return uint64(day.Unix())
}

func DaySe(year, month, day, hour, min, sec int) uint64 {
	day1 := time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local)
	return uint64(day1.Unix())
}

func CurrTimeInfo() (year, month, day, hour, min, sec int) {
	curr := time.Now()
	return curr.Year(), int(curr.Month()), curr.Day(), curr.Hour(), curr.Minute(), curr.Second()
}

func GetTimeInfo(se int64) (year, month, day, hour, min, sec int) {
	curr := time.Unix(se, 0)
	return curr.Year(), int(curr.Month()), curr.Day(), curr.Hour(), curr.Minute(), curr.Second()
}

//当前时间距离24点的时间差
func GetDeffTime() int64 {

	begintime := time.Now().Unix()

	endtime := time.Unix(begintime+86400, 0).Format("2006-01-02") + " 00:00:00"

	loc, _ := time.LoadLocation("Local") //重要：获取时区

	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", endtime, loc) //使用模板在对应时区转化为time.time类型

	return theTime.Unix() - begintime
}

//获取本地location
//toBeCharge 待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
func StringToSe(toBeCharge string, types int) uint64 {
	timeLayout := "2006-01-02 15:04:05" //转化所需模板
	if types == 2 {
		timeLayout = "2006-01-02 15:04" //转化所需模板
	}
	if types == 3 {
		timeLayout = "2006-01-02 15" //转化所需模板
	}
	if types == 4 {
		timeLayout = "2006-01-02" //转化所需模板
	}
	if types == 5 {
		timeLayout = "2006.01.02" //转化所需模板
	}
	loc, _ := time.LoadLocation("Local")                              //重要：获取时区
	theTime, err := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return 0
	}

	sr := theTime.Unix() //转化为时间戳 类型是int64

	return uint64(sr)
}

func DaySeParse(f string, t string) uint64 {
	loc, _ := time.LoadLocation("Local")            //重要：获取时区
	theTime, err := time.ParseInLocation(f, t, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return 0
	}

	sr := theTime.Unix() //转化为时间戳 类型是int64

	return uint64(sr)
}
