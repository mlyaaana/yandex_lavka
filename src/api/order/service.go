package order

import (
	"context"
	"time"

	"yandex-team.ru/bstask/domain"
	timeutil "yandex-team.ru/bstask/util/time"
)

type Service interface {
	CreateOrders(
		ctx context.Context,
		in *CreateOrdersIn,
	) (*CreateOrdersOut, error)
	GetOrder(
		ctx context.Context,
		in *GetOrderIn,
	) (*GetOrderOut, error)
	ListOrders(
		ctx context.Context,
		in *ListOrdersIn,
	) (*ListOrdersOut, error)
	CompleteOrders(
		ctx context.Context,
		in *CompleteOrdersIn,
	) (*CompleteOrdersOut, error)
	AssignOrders(
		ctx context.Context,
		in *AssignOrdersIn,
	) (*AssignOrdersOut, error)
	GetAssignedOrders(
		ctx context.Context,
		in *GetAssignedOrdersIn,
	) (*GetAssignedOrdersOut, error)
}

type CreateOrder struct {
	Weight        float32
	Region        domain.Region
	DeliveryHours timeutil.Range
	Cost          int
}

type CreateOrdersIn struct {
	Orders []*CreateOrder
}

type CreateOrdersOut struct {
	Orders []*domain.Order
}

type GetOrderIn struct {
	OrderId domain.OrderId
}

type GetOrderOut struct {
	Order *domain.Order
}

type ListOrdersIn struct {
	Limit  int
	Offset int
}

type ListOrdersOut struct {
	Orders []*domain.Order
}

type CompleteOrder struct {
	CourierId    domain.CourierId
	OrderId      domain.OrderId
	CompleteTime time.Time
}

type CompleteOrdersIn struct {
	CompleteOrders []*CompleteOrder
}

type CompleteOrdersOut struct {
	Orders []*domain.Order
}

type Group struct {
	Id     domain.OrderGroupId
	Orders []*domain.Order
}

type AssignOrder struct {
	CourierId domain.CourierId
	Groups    []*Group
}

type AssignOrdersIn struct {
	Date time.Time
}

type AssignOrdersOut struct {
	Date         time.Time
	AssignOrders []*AssignOrder
}

type GetAssignedOrdersIn struct {
	Date      time.Time
	CourierId domain.CourierId
}

type GetAssignedOrdersOut struct {
	Date         time.Time
	AssignOrders []*AssignOrder
}
