package timestamp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTimestamp(t *testing.T) {
	timestamp, err := ParseTimestamp("2006-01-02 15:04:05", "Asia/Shanghai", "2018-07-11 15:07:51")
	if err != nil {
		t.Errorf("ParseTimestamp err %v", err)
	}
	t.Logf("timestamp %v", timestamp)
	assert.NotZero(t, timestamp)
}

func TestParseTimestampSecond(t *testing.T) {
	timestampSecond, err := ParseTimestampSecond("2018-07-11 15:07:51")
	if err != nil {
		t.Errorf("ParseTimestampSecond err %v", err)
	}
	t.Logf("timestampSecond %v", timestampSecond)
	assert.NotZero(t, timestampSecond)

	errorTimeStr := "2018:07:11 15:07:51"
	str2TimestampSecondErr, err := ParseTimestampSecond(errorTimeStr)
	if err == nil {
		t.Errorf("ParseTimestampSecond not found error at %v", errorTimeStr)
	}
	assert.Zero(t, str2TimestampSecondErr)
}

func TestParesTimestampMicro(t *testing.T) {
	timestampMicro, err := ParesTimestampMicro("2018-07-11 15:07:51.456123")
	if err != nil {
		t.Errorf("ParesTimestampMicro err %v", err)
	}
	t.Logf("timestampMicro %v", timestampMicro)
	assert.NotZero(t, timestampMicro)

	errorTimeStr := "2018:07:11 15:07:51.456123"
	str2TimestampMicroErr, err := ParesTimestampMicro(errorTimeStr)
	if err == nil {
		t.Errorf("ParesTimestampMicro not found error at %v", errorTimeStr)
	}
	assert.Zero(t, str2TimestampMicroErr)
}

func TestParseLocationSecond(t *testing.T) {
	parseLocationSecond, err := ParseLocationSecond("2020-02-26 10:08:58", "UTC", "Asia/Shanghai")
	if err != nil {
		t.Errorf("ParseLocationSecond err: %v", parseLocationSecond)
	}
	t.Logf("parseLocationSecond %v", parseLocationSecond)
	assert.NotEmpty(t, parseLocationSecond)

	errorTimeString := "2020:02#26 10:08:58"
	locationSecondErr, err := ParseLocationSecond(errorTimeString, "UTC", "Asia/Shanghai")
	if err == nil {
		t.Errorf("ParseLocationSecond not return error")
	}
	assert.Empty(t, locationSecondErr)
}

func TestParseLocationMicro(t *testing.T) {
	parseLocationMicro, err := ParseLocationMicro("2020-02-26 10:08:58", "UTC", "Asia/Shanghai")
	if err != nil {
		t.Errorf("ParseLocationMicro err: %v", parseLocationMicro)
	}
	t.Logf("parseLocationMicro %v", parseLocationMicro)
	assert.NotEmpty(t, parseLocationMicro)

	errorTimeString := "2020:02#26 10:08:58"
	locationMicroErr, err := ParseLocationMicro(errorTimeString, "UTC", "Asia/Shanghai")
	if err == nil {
		t.Errorf("ParseLocationMicro not return error")
	}
	assert.Empty(t, locationMicroErr)
}
