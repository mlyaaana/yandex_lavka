package domain

import (
	"github.com/shopspring/decimal"

	"yandex-team.ru/bstask/util/time"
)

type OrderGroup struct {
	CourierId CourierId
	GroupId   OrderGroupId
	Orders    []*Order
	Regions   map[Region]struct{}
	Range     time.Range
	Weight    decimal.Decimal
}
