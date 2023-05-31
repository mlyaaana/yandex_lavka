package model

import (
	"github.com/lib/pq"

	"yandex-team.ru/bstask/util/sql"
)

type Courier struct {
	Id           int64              `gorm:"id;type:bigserial;primary_key"`
	Type         string             `gorm:"type;type:varchar;not null"`
	Regions      pq.Int32Array      `gorm:"regions;type:int[];not null"`
	WorkingHours sql.TimeRangeArray `gorm:"working_hours;type:varchar[];not null"`
}

func (Courier) TableName() string {
	return "couriers"
}
