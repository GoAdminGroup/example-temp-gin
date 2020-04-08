package timestamp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareLocationEQ(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compareLocationEQ, err := CompareLocationEQ(timeFrom, timeTo, "2006-01-02 15:04:05", "Asia/Shanghai")
	if err != nil {
		t.Errorf("CompareLocationEQ err %v", err)
	}
	t.Logf("compareLocationEQ %v", compareLocationEQ)
	assert.False(t, compareLocationEQ)

	timeFrom2 := "2018-07-11 15:07:51"
	timeTo2 := "2018-07-11 15:07:51.000000"
	locationEQ, err := CompareLocationEQ(timeFrom2, timeTo2, "2006-01-02 15:04:05", "Asia/Shanghai")
	if err != nil {
		t.Errorf("CompareLocationEQ err %v", err)
	}
	assert.True(t, locationEQ)
}

func TestCompareLocationLT(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compareLocationLT, err := CompareLocationLT(timeFrom, timeTo, "2006-01-02 15:04:05", "Asia/Shanghai")
	if err != nil {
		t.Errorf("CompareLocationLT err %v", err)
	}
	t.Logf("compareLocationLT %v", compareLocationLT)
	assert.True(t, compareLocationLT)
}

func TestCompareLocationGT(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compareLocationGT, err := CompareLocationGT(timeFrom, timeTo, "2006-01-02 15:04:05", "Asia/Shanghai")
	if err != nil {
		t.Errorf("CompareLocationGT err %v", err)
	}
	t.Logf("compareLocationGT %v", compareLocationGT)
	assert.False(t, compareLocationGT)
}

func TestCompareUTCSecondGT(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compareUTCSecondGT, err := CompareUTCSecondGT(timeFrom, timeTo)
	if err != nil {
		t.Errorf("CompareUTCSecondGT err %v", err)
	}
	t.Logf("compareUTCSecondGT %v", compareUTCSecondGT)
	assert.False(t, compareUTCSecondGT)
}

func TestCompareUTCSecondLT(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compareUTCSecondLT, err := CompareUTCSecondLT(timeFrom, timeTo)
	if err != nil {
		t.Errorf("CompareUTCSecondLT err %v", err)
	}
	t.Logf("compareUTCSecondLT %v", compareUTCSecondLT)
	assert.True(t, compareUTCSecondLT)
}

func TestCompareUTCSecondEQ(t *testing.T) {
	timeFrom := "2018-07-11 15:07:51"
	timeTo := "2018-07-12 15:07:51"
	compareUTCSecondEQ, err := CompareUTCSecondEQ(timeFrom, timeTo)
	if err != nil {
		t.Errorf("CompareUTCSecondEQ err %v", err)
	}
	t.Logf("compareUTCSecondEQ %v", compareUTCSecondEQ)
	assert.False(t, compareUTCSecondEQ)

	timeFrom2 := "2018-07-11 15:07:51"
	timeTo2 := "2018-07-11 15:07:51.000000"
	utcSecondEQ, err := CompareUTCSecondEQ(timeFrom2, timeTo2)
	if err != nil {
		t.Errorf("CompareUTCSecondEQ err %v", err)
	}
	assert.True(t, utcSecondEQ)
}
