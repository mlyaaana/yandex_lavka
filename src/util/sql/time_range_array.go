package sql

import (
	"database/sql/driver"

	"github.com/lib/pq"
	"github.com/samber/lo"

	"yandex-team.ru/bstask/util/time"
)

type TimeRangeArray []time.Range

func (t *TimeRangeArray) Scan(src any) error {
	ranges := pq.StringArray{}
	if err := ranges.Scan(src); err != nil {
		return err
	}

	arr := make(TimeRangeArray, 0, len(ranges))
	for _, s := range ranges {
		timeRange, err := time.ParseRange(s)
		if err != nil {
			return err
		}
		arr = append(arr, timeRange)
	}
	*t = arr
	return nil
}

func (t TimeRangeArray) Value() (driver.Value, error) {
	return lo.Map(t, func(item time.Range, _ int) string {
		return item.String()
	}), nil
}
