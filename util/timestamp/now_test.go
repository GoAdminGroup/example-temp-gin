package timestamp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUTCTimeSecond(t *testing.T) {
	utcTimeSecond := UTCTimeSecond()
	t.Logf("utcTimeSecond %v", utcTimeSecond)
	assert.NotEmpty(t, utcTimeSecond)
}

func TestUTCTimeMicro(t *testing.T) {
	utcTimeMicro := UTCTimeMicro()
	t.Logf("utcTimeMicro %v", utcTimeMicro)
	assert.NotEmpty(t, utcTimeMicro)
}

func TestLocationSecond(t *testing.T) {
	locationSecond := LocationSecond("Asia/Shanghai")
	t.Logf("locationSecond %v", locationSecond)
	assert.NotEmpty(t, locationSecond)

	locationError := LocationSecond("Asia")
	assert.Empty(t, locationError)
}

func TestLocationMicro(t *testing.T) {
	locationMicro := LocationMicro("Asia/Shanghai")

	t.Logf("locationSecond %v", locationMicro)
	assert.NotEmpty(t, locationMicro)

	locationError := LocationMicro("Asia")
	assert.Empty(t, locationError)
}

func TestLocalTimeSecond(t *testing.T) {
	localTimeSecond := LocalTimeSecond()
	t.Logf("localTimeSecond %v", localTimeSecond)
	assert.NotEmpty(t, localTimeSecond)
}

func TestLocalTimeMicro(t *testing.T) {
	localTimeMicro := LocalTimeMicro()
	t.Logf("localTimeMicro %v", localTimeMicro)
	assert.NotEmpty(t, localTimeMicro)
}
