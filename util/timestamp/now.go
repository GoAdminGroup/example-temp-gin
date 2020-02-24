package timestamp

import "time"

func LocalTimeSecond() string {
	return time.Now().Local().Format("2006-01-02 15:04:05")
}
func LocalTimeMicro() string {
	return time.Now().Local().Format("2006-01-02 15:04:05.999999")
}
