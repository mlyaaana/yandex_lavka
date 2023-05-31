package domain

import (
	"time"

	"github.com/shopspring/decimal"

	timeutil "yandex-team.ru/bstask/util/time"
)

type CourierId int64

var NilCourierId CourierId = 0

type CourierType string

const (
	CourierTypeFoot CourierType = "FOOT"
	CourierTypeBike CourierType = "BIKE"
	CourierTypeAuto CourierType = "AUTO"
)

var CourierTypes = map[CourierType]bool{
	CourierTypeFoot: true,
	CourierTypeBike: true,
	CourierTypeAuto: true,
}

type Courier struct {
	Id           CourierId
	Type         CourierType
	Regions      []Region
	WorkingHours []timeutil.Range
}

type CourierStats struct {
	Courier *Courier
	Orders  int
	Income  int
}

var CourierRegionsCount = map[CourierType]int{
	CourierTypeFoot: 1,
	CourierTypeBike: 2,
	CourierTypeAuto: 3,
}

var CourierEarningsCoefficients = map[CourierType]int{
	CourierTypeFoot: 2,
	CourierTypeBike: 3,
	CourierTypeAuto: 4,
}

var CourierRatingCoefficients = map[CourierType]int{
	CourierTypeFoot: 3,
	CourierTypeBike: 2,
	CourierTypeAuto: 1,
}

var CourierGroupSize = map[CourierType]int{
	CourierTypeFoot: 2,
	CourierTypeBike: 4,
	CourierTypeAuto: 7,
}

var CourierGroupWeight = map[CourierType]decimal.Decimal{
	CourierTypeFoot: decimal.NewFromInt(10),
	CourierTypeBike: decimal.NewFromInt(20),
	CourierTypeAuto: decimal.NewFromInt(40),
}

var CourierFirstOrderTravelTime = map[CourierType]time.Duration{
	CourierTypeFoot: time.Minute * time.Duration(25),
	CourierTypeBike: time.Minute * time.Duration(12),
	CourierTypeAuto: time.Minute * time.Duration(8),
}

var CourierGroupOrdersTravelTime = map[CourierType]time.Duration{
	CourierTypeFoot: time.Minute * time.Duration(10),
	CourierTypeBike: time.Minute * time.Duration(8),
	CourierTypeAuto: time.Minute * time.Duration(4),
}
