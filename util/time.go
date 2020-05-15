package util

import (
	"fmt"
	"time"
)

var timeFormate = "2006-01-02 15:04:05"
var localTimeZone, _ = time.LoadLocation("Local") //服务器设置的时区

// GetDatetime 获取当前时间
func GetNowStr() (dateTime string) {
	dateTime = time.Now().In(localTimeZone).Format(timeFormate)
	return
}

type JsonTime time.Time

//MarshalJSON 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Local().Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
