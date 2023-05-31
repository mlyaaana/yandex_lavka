package model

import (
	"time"

	"github.com/samber/lo"

	"yandex-team.ru/bstask/domain"
	"yandex-team.ru/bstask/util/sql"
	timeutil "yandex-team.ru/bstask/util/time"
)

func PackCourier(in *domain.Courier) *Courier {
	regions := lo.Map(in.Regions, func(r domain.Region, _ int) int32 {
		return int32(r)
	})
	return &Courier{
		Type:         string(in.Type),
		Regions:      regions,
		WorkingHours: in.WorkingHours,
	}
}

func PackCouriers(in []*domain.Courier) []*Courier {
	return lo.Map(in, func(c *domain.Courier, _ int) *Courier {
		return PackCourier(c)
	})
}

func UnpackCourier(in *Courier) *domain.Courier {
	regions := lo.Map(in.Regions, func(r int32, _ int) domain.Region {
		return domain.Region(r)
	})
	return &domain.Courier{
		Id:           domain.CourierId(in.Id),
		Type:         domain.CourierType(in.Type),
		Regions:      regions,
		WorkingHours: in.WorkingHours,
	}
}

func UnpackCouriers(in []*Courier) []*domain.Courier {
	return lo.Map(in, func(c *Courier, _ int) *domain.Courier {
		return UnpackCourier(c)
	})
}

func PackOrder(in *domain.Order) *Order {
	return &Order{
		Id:            int64(in.Id),
		CourierId:     int64(in.CourierId),
		GroupId:       int64(in.GroupId),
		Weight:        in.Weight,
		Region:        int(in.Region),
		DeliveryHours: sql.TimeRange(in.DeliveryHours),
		Cost:          in.Cost,
		CreatedAt:     in.CreatedAt,
		CompletedAt:   sql.NullTime(in.CompletedAt),
		AssignedAt:    sql.NullTime(in.AssignedAt),
	}
}

func PackOrders(in []*domain.Order) []*Order {
	return lo.Map(in, func(o *domain.Order, _ int) *Order {
		return PackOrder(o)
	})
}

func UnpackOrder(in *Order) *domain.Order {
	return &domain.Order{
		Id:            domain.OrderId(in.Id),
		CourierId:     domain.CourierId(in.CourierId),
		GroupId:       domain.OrderGroupId(in.GroupId),
		Weight:        in.Weight,
		Region:        domain.Region(in.Region),
		DeliveryHours: timeutil.Range(in.DeliveryHours),
		Cost:          in.Cost,
		CreatedAt:     in.CreatedAt,
		CompletedAt:   time.Time(in.CompletedAt),
		AssignedAt:    time.Time(in.AssignedAt),
	}
}

func UnpackOrders(in []*Order) []*domain.Order {
	return lo.Map(in, func(o *Order, _ int) *domain.Order {
		return UnpackOrder(o)
	})
}
