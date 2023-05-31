package order

import (
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"yandex-team.ru/bstask/api/order"
	"yandex-team.ru/bstask/domain"
	timeutil "yandex-team.ru/bstask/util/time"
)

func packCreateOrder(in *order.CreateOrder) *domain.Order {
	return &domain.Order{
		Weight:        decimal.NewFromFloat32(in.Weight),
		Region:        in.Region,
		DeliveryHours: in.DeliveryHours,
		Cost:          in.Cost,
		CreatedAt:     timeutil.NowDate(),
	}
}

func packCreateOrders(in []*order.CreateOrder) []*domain.Order {
	return lo.Map(in, func(o *order.CreateOrder, _ int) *domain.Order {
		return packCreateOrder(o)
	})
}

func unpackAssignOrders(in map[domain.CourierId][]*domain.OrderGroup) []*order.AssignOrder {
	return lo.MapToSlice(in, func(id domain.CourierId, groups []*domain.OrderGroup) *order.AssignOrder {
		return &order.AssignOrder{
			CourierId: id,
			Groups: lo.Map(groups, func(group *domain.OrderGroup, _ int) *order.Group {
				return &order.Group{
					Id:     group.GroupId,
					Orders: group.Orders,
				}
			}),
		}
	})
}
