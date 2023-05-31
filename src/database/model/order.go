package model

import (
	"time"

	"github.com/shopspring/decimal"

	"yandex-team.ru/bstask/util/sql"
)

type Order struct {
	Id            int64           `gorm:"id;type:bigserial;primary_key"`
	CourierId     int64           `gorm:"courier_id;type:bigint;index:idx_courier_id"`
	GroupId       int64           `gorm:"group_id;type:bigint"`
	Weight        decimal.Decimal `gorm:"weight;type:numeric;not null"`
	Region        int             `gorm:"region;type:int;not null"`
	DeliveryHours sql.TimeRange   `gorm:"delivery_hours;type:varchar;not null"`
	Cost          int             `gorm:"cost;type:int;not null"`
	CreatedAt     time.Time       `gorm:"created_at;type:timestamp"`
	AssignedAt    sql.NullTime    `gorm:"assigned_at;type:timestamp"`
	CompletedAt   sql.NullTime    `gorm:"completed_at;type:timestamp"`
}

func (Order) TableName() string {
	return "orders"
}
