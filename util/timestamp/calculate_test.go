package timestamp

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCalcDurationSecondUTC(t *testing.T) {
	calcDurationSecondUTC := CalcDurationSecondUTC(time.Duration(-30) * time.Minute)
	t.Logf("calcDurationSecondUTC: %v", calcDurationSecondUTC)
	assert.NotEmpty(t, calcDurationSecondUTC)
}

func TestCalcDurationMicroUTC(t *testing.T) {
	calcDurationMicroUTC := CalcDurationMicroUTC(time.Duration(30) * time.Minute)
	t.Logf("calcDurationMicroUTC: %v", calcDurationMicroUTC)
	assert.NotEmpty(t, calcDurationMicroUTC)
}

func TestCalcDayUTC(t *testing.T) {
	calcDaySecondUTC := CalcDaySecondUTC(1)
	t.Logf("calcDaySecondUTC: %v", calcDaySecondUTC)
	assert.NotEmpty(t, calcDaySecondUTC)
	calcDayMicroUTC := CalcDayMicroUTC(2)
	t.Logf("calcDayMicroUTC: %v", calcDayMicroUTC)
	assert.NotEmpty(t, calcDayMicroUTC)
}

func TestCalcDayUTCZero(t *testing.T) {
	daySecondUTCZero := CalcDaySecondUTCZero(1)
	t.Logf("daySecondUTCZero: %v", daySecondUTCZero)
	assert.NotEmpty(t, daySecondUTCZero)

	calcDayMicroUTCZero := CalcDayMicroUTCZero(2)
	t.Logf("calcDayMicroUTCZero: %v", calcDayMicroUTCZero)
	assert.NotEmpty(t, calcDayMicroUTCZero)
}

func TestCalcDayLocation(t *testing.T) {
	calcDayLocation := CalcDayLocation(1, "2006-01-02 15:04:05", "Asia/Shanghai")
	t.Logf("calcDayLocation: %v", calcDayLocation)
	assert.NotEmpty(t, calcDayLocation)
}

func TestCalcDayLocationZero(t *testing.T) {
	calcDayLocationZero := CalcDayLocationZero(2, "2006-01-02 15:04:05", "Asia/Shanghai")
	t.Logf("calcDayLocationZero: %v", calcDayLocationZero)
	assert.NotEmpty(t, calcDayLocationZero)
}
