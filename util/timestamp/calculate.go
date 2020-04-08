package timestamp

import "time"

// get location time by duration
//	duration like time.Duration(yourTime) * time.Second
//	layout 2006-01-02 15:04:05
//	locationName UTC Asia/Shanghai of others, warning if LoadLocation not support will return ""
func CalcDurationLocation(duration time.Duration, layout string, locationName string) string {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return ""
	}
	return time.Now().
		Add(duration).
		In(loc).
		Format(layout)
}

// get time duration, location UTC
//	duration like time.Duration(yourTime) * time.Second
//	layout 2006-01-02 15:04:05
func CalcDurationUTC(duration time.Duration, layout string) string {
	if locationUTC == nil {
		locationUTC, _ = time.LoadLocation("UTC")
	}
	return time.Now().
		Add(duration).
		In(locationUTC).
		Format(layout)
}

// get time duration, location UTC format as "2006-01-02 15:04:05"
//	duration like time.Duration(yourTime) * time.Second
func CalcDurationSecondUTC(duration time.Duration) string {
	return CalcDurationUTC(duration, "2006-01-02 15:04:05")
}

// get time duration, location UTC format as "2006-01-02 15:04:05.999999"
//	duration like time.Duration(yourTime) * time.Second
func CalcDurationMicroUTC(duration time.Duration) string {
	return CalcDurationUTC(duration, "2006-01-02 15:04:05.999999")
}

// get location time day
//	day 1 tomorrow -1 yesterday
//	layout 2006-01-02 15:04:05
//	locationName UTC Asia/Shanghai of others, warning if LoadLocation not support will return ""
func CalcDayLocation(days int, layout string, locationName string) string {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return ""
	}
	return time.Now().
		AddDate(0, 0, days).
		In(loc).
		Format(layout)
}

// get location time day, location UTC
//	day 1 tomorrow -1 yesterday
//	layout 2006-01-02 15:04:05
func CalcDayUTC(days int, layout string) string {
	if locationUTC == nil {
		locationUTC, _ = time.LoadLocation("UTC")
	}
	return time.Now().
		AddDate(0, 0, days).
		In(locationUTC).
		Format(layout)
}

// get location time day, location UTC "2006-01-02 15:04:05"
//	day 1 tomorrow -1 yesterday
func CalcDaySecondUTC(days int) string {
	return CalcDayUTC(days, "2006-01-02 15:04:05")
}

// get location time day, location UTC format "2006-01-02 15:04:05.999999"
//	day 1 tomorrow -1 yesterday
func CalcDayMicroUTC(days int) string {
	return CalcDayUTC(days, "2006-01-02 15:04:05.999999")
}

// get location time day zero.
//	day 1 tomorrow -1 yesterday
//	layout 2006-01-02 15:04:05
//	locationName UTC Asia/Shanghai of others, warning if LoadLocation not support will return ""
func CalcDayLocationZero(days int, layout string, locationName string) string {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return ""
	}
	timeNow := time.Now().In(locationUTC).Format("2006-01-02")
	timeZero, _ := time.Parse("2006-01-02", timeNow)
	return timeZero.
		AddDate(0, 0, days).
		In(loc).
		Format(layout)
}

// get location time day zero, location UTC
//	day 1 tomorrow -1 yesterday
//	layout 2006-01-02 15:04:05
func CalcDayUTCZero(days int, layout string) string {
	if locationUTC == nil {
		locationUTC, _ = time.LoadLocation("UTC")
	}
	timeNow := time.Now().In(locationUTC).Format("2006-01-02")
	timeZero, _ := time.Parse("2006-01-02", timeNow)
	return timeZero.
		AddDate(0, 0, days).
		In(locationUTC).
		Format(layout)
}

func CalcDaySecondUTCZero(days int) string {
	return CalcDayUTCZero(days, "2006-01-02 15:04:05")
}

func CalcDayMicroUTCZero(days int) string {
	return CalcDayUTCZero(days, "2006-01-02 15:04:05.999999")
}
