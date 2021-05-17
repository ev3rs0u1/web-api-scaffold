package model

import (
	"database/sql/driver"
	"time"
)

type BaseTime struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type BaseJsonTime struct {
	CreatedAt DateTime `json:"created_at"`
	UpdatedAt DateTime `json:"updated_at"`
}

type DateTime struct {
	time.Time
}

func (t DateTime) UnixTimestamp() int64 {
	return t.UnixNano() / 1e6
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Format(`"2006-01-02 15:04:05"`)), nil
}

func (t DateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
