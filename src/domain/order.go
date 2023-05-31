package domain

import (
	"time"

	"github.com/shopspring/decimal"

	timeutil "yandex-team.ru/bstask/util/time"
)

type OrderId int64
type OrderGroupId int64

type Order struct {
	Id            OrderId
	CourierId     CourierId
	GroupId       OrderGroupId
	Weight        decimal.Decimal
	Region        Region
	DeliveryHours timeutil.Range
	Cost          int
	CreatedAt     time.Time
	CompletedAt   time.Time
	AssignedAt    time.Time
}
