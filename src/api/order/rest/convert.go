package rest

import (
	"errors"

	"github.com/samber/lo"

	"yandex-team.ru/bstask/api/order"
	"yandex-team.ru/bstask/domain"
	"yandex-team.ru/bstask/gen"
	"yandex-team.ru/bstask/util/time"
)

func packDeliveryHours(in string) (time.Range, error) {
	return time.ParseRange(in)
}

func unpackDeliveryHours(in time.Range) string {
	return in.String()
}

func unpackOrder(in *domain.Order) gen.OrderDto {
	completedTime := &in.CompletedAt
	if completedTime.IsZero() {
		completedTime = nil
	}
	return gen.OrderDto{
		OrderId:       int64(in.Id),
		Weight:        float32(in.Weight.InexactFloat64()),
		Regions:       int32(in.Region),
		DeliveryHours: []string{unpackDeliveryHours(in.DeliveryHours)},
		Cost:          int32(in.Cost),
		CompletedTime: completedTime,
	}
}

func unpackOrders(in []*domain.Order) []gen.OrderDto {
	return lo.Map(in, func(o *domain.Order, _ int) gen.OrderDto {
		return unpackOrder(o)
	})
}

func packCreateOrders(
	in *gen.CreateOrderRequest,
) (*order.CreateOrdersIn, error) {
	orders := make([]*order.CreateOrder, 0, len(in.Orders))
	for _, o := range in.Orders {
		if len(o.DeliveryHours) != 1 {
			return nil, errors.New("delivery hours should be a single interval")
		}
		deliveryHours, err := packDeliveryHours(o.DeliveryHours[0])
		if err != nil {
			return nil, err
		}

		orders = append(orders, &order.CreateOrder{
			Weight:        o.Weight,
			Region:        domain.Region(o.Regions),
			DeliveryHours: deliveryHours,
			Cost:          int(o.Cost),
		})
	}
	return &order.CreateOrdersIn{
		Orders: orders,
	}, nil
}

func packCompleteOrders(in []gen.CompleteOrder) []*order.CompleteOrder {
	return lo.Map(in, func(c gen.CompleteOrder, _ int) *order.CompleteOrder {
		return &order.CompleteOrder{
			CourierId:    domain.CourierId(c.CourierId),
			OrderId:      domain.OrderId(c.OrderId),
			CompleteTime: c.CompleteTime,
		}
	})
}

func unpackOrderGroups(in []*order.Group) []gen.GroupOrders {
	return lo.Map(in, func(g *order.Group, _ int) gen.GroupOrders {
		return gen.GroupOrders{
			GroupOrderId: int64(g.Id),
			Orders:       unpackOrders(g.Orders),
		}
	})
}

func unpackAssignOrders(in []*order.AssignOrder) []gen.CouriersGroupOrders {
	return lo.Map(in, func(o *order.AssignOrder, _ int) gen.CouriersGroupOrders {
		return gen.CouriersGroupOrders{
			CourierId: int64(o.CourierId),
			Orders:    unpackOrderGroups(o.Groups),
		}
	})
}

func unpackAssignOrdersResponse(in *order.AssignOrdersOut) []gen.OrderAssignResponse {
	return []gen.OrderAssignResponse{
		{
			Date:     in.Date.Format(time.DateLayout),
			Couriers: unpackAssignOrders(in.AssignOrders),
		},
	}
}

func unpackGetAssignedOrdersResponse(in *order.GetAssignedOrdersOut) []gen.OrderAssignResponse {
	return []gen.OrderAssignResponse{
		{
			Date:     in.Date.Format(time.DateLayout),
			Couriers: unpackAssignOrders(in.AssignOrders),
		},
	}
}
