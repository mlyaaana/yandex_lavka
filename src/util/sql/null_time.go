package sql

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

type NullTime time.Time

func (t NullTime) Value() (driver.Value, error) {
	x := time.Time(t)
	if x.IsZero() {
		return nil, nil
	}
	nt := sql.NullTime{
		Time:  x,
		Valid: true,
	}
	return nt.Value()
}

func (t *NullTime) Scan(in interface{}) error {
	var nt sql.NullTime
	if err := nt.Scan(in); err != nil {
		return err
	}
	var x time.Time
	if nt.Valid {
		x = nt.Time
	}
	*t = NullTime(x)
	return nil
}
