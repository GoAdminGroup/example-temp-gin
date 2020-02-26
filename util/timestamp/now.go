package timestamp

import "time"

var locationUTC *time.Location

func UTCTimeSecond() string {
	if locationUTC == nil {
		locationUTC, _ = time.LoadLocation("UTC")
	}
	return time.Now().In(locationUTC).Format("2006-01-02 15:04:05")
}
func UTCTimeMicro() string {
	if locationUTC == nil {
		locationUTC, _ = time.LoadLocation("UTC")
	}
	return time.Now().In(locationUTC).Format("2006-01-02 15:04:05.999999")
}

// get location time format as 2006-01-02 15:04:05
//	locationName UTC Asia/Shanghai of others, warning if LoadLocation not support will return ""
func LocationSecond(locationName string) string {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return ""
	}
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

// get location time format as 2006-01-02 15:04:05.999999
//	locationName UTC Asia/Shanghai of others, warning if LoadLocation not support will return ""
func LocationMicro(locationName string) string {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return ""
	}
	return time.Now().In(loc).Format("2006-01-02 15:04:05.999999")
}

func LocalTimeSecond() string {
	return time.Now().Local().Format("2006-01-02 15:04:05")
}

func LocalTimeMicro() string {
	return time.Now().Local().Format("2006-01-02 15:04:05.999999")
}
