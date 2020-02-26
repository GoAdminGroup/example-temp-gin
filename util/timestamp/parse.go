package timestamp

import "time"

// time string to timestamp
//	format -> 2006-01-02 15:04:05.999
//	locationName -> UTC Asia/Shanghai
//	timeStr -> for parse time string
// when error will return 0 and err
func ParseTimestamp(format, locationName, timeStr string) (int64, error) {
	location, err := time.LoadLocation(locationName)
	if err != nil {
		return 0, err
	}
	parseInLocation, err := time.ParseInLocation(format, timeStr, location)
	if err != nil {
		return 0, err
	}
	return parseInLocation.Unix(), nil
}

// time string must as 2006-01-02 15:04:05 UTC
func ParseTimestampSecond(timeStr string) (int64, error) {
	return ParseTimestamp("2006-01-02 15:04:05", "UTC", timeStr)
}

// time string must as 2006-01-02 15:04:05.999999 UTC
func ParesTimestampMicro(timeStr string) (int64, error) {
	return ParseTimestamp("2006-01-02 15:04:05.999999", "UTC", timeStr)
}

// Parse location from location to next
//	format -> 2006-01-02 15:04:05.999999
//	timeStr -> for parse time string
//	fromLocation -> parse from like UTC
//	toLocation   -> parse to   like Asia/Shanghai
func ParseLocation(format, timeStr, fromLocation, toLocation string) (string, error) {
	fromLoc, err := time.LoadLocation(fromLocation)
	if err != nil {
		return "0", err
	}
	toLoc, err := time.LoadLocation(toLocation)
	if err != nil {
		return "0", err
	}
	parse, err := time.ParseInLocation(format, timeStr, fromLoc)
	if err != nil {
		return "", err
	}
	return parse.In(toLoc).Format(format), nil
}

// use format 2006-01-02 15:04:05 as ParseLocation()
func ParseLocationSecond(timeStr, fromLocation, toLocation string) (string, error) {
	return ParseLocation("2006-01-02 15:04:05", timeStr, fromLocation, toLocation)
}

// use format 2006-01-02 15:04:05.999999 as ParseLocation()
func ParseLocationMicro(timeStr, fromLocation, toLocation string) (string, error) {
	return ParseLocation("2006-01-02 15:04:05.999999", timeStr, fromLocation, toLocation)
}
