package timestamp

import "time"

func CompareUTCSecondLT(timestampFrom, timestampTo string) (bool, error) {
	return CompareUTCLT(timestampFrom, timestampTo, "2006-01-02 15:04:05")
}

func CompareUTCLT(timestampFrom, timestampTo string, layout string) (bool, error) {
	if locationUTC == nil {
		locationUTC, _ = time.LoadLocation("UTC")
	}
	from, err := time.ParseInLocation(layout, timestampFrom, locationUTC)
	if err != nil {
		return false, err
	}
	to, err := time.ParseInLocation(layout, timestampTo, locationUTC)
	if err != nil {
		return false, err
	}
	return from.Before(to), nil
}

func CompareLocationLT(timestampFrom, timestampTo string, layout string, locationName string) (bool, error) {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return false, err
	}
	from, err := time.ParseInLocation(layout, timestampFrom, loc)
	if err != nil {
		return false, err
	}
	to, err := time.ParseInLocation(layout, timestampTo, loc)
	if err != nil {
		return false, err
	}
	return from.Before(to), nil
}

func CompareUTCSecondGT(timestampFrom, timestampTo string) (bool, error) {
	return CompareUTCGT(timestampFrom, timestampTo, "2006-01-02 15:04:05")
}

func CompareUTCGT(timestampFrom, timestampTo string, layout string) (bool, error) {
	if locationUTC == nil {
		locationUTC, _ = time.LoadLocation("UTC")
	}
	from, err := time.ParseInLocation(layout, timestampFrom, locationUTC)
	if err != nil {
		return false, err
	}
	to, err := time.ParseInLocation(layout, timestampTo, locationUTC)
	if err != nil {
		return false, err
	}
	return from.After(to), nil
}

func CompareLocationGT(timestampFrom, timestampTo string, layout string, locationName string) (bool, error) {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return false, err
	}
	from, err := time.ParseInLocation(layout, timestampFrom, loc)
	if err != nil {
		return false, err
	}
	to, err := time.ParseInLocation(layout, timestampTo, loc)
	if err != nil {
		return false, err
	}
	return from.After(to), nil
}

func CompareUTCEQ(timestampFrom, timestampTo string, layout string) (bool, error) {
	if locationUTC == nil {
		locationUTC, _ = time.LoadLocation("UTC")
	}
	from, err := time.ParseInLocation(layout, timestampFrom, locationUTC)
	if err != nil {
		return false, err
	}
	to, err := time.ParseInLocation(layout, timestampTo, locationUTC)
	if err != nil {
		return false, err
	}
	return from.Equal(to), nil
}

func CompareUTCSecondEQ(timestampFrom, timestampTo string) (bool, error) {
	return CompareUTCEQ(timestampFrom, timestampTo, "2006-01-02 15:04:05")
}

func CompareLocationEQ(timestampFrom, timestampTo string, layout string, locationName string) (bool, error) {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return false, err
	}
	from, err := time.ParseInLocation(layout, timestampFrom, loc)
	if err != nil {
		return false, err
	}
	to, err := time.ParseInLocation(layout, timestampTo, loc)
	return from.Equal(to), nil
}
