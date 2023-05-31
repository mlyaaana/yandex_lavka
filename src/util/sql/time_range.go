package sql

import (
	"database/sql/driver"

	"yandex-team.ru/bstask/util/time"
)

type TimeRange time.Range

func (t *TimeRange) Scan(src any) error {
	str := src.(string)
	timeRange, err := time.ParseRange(str)
	if err != nil {
		return err
	}
	*t = TimeRange(timeRange)
	return nil
}

func (t TimeRange) Value() (driver.Value, error) {
	return time.Range(t).String(), nil
}
