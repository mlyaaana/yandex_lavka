package order

import (
	"context"
	"fmt"
	"time"

	"github.com/samber/lo"

	"yandex-team.ru/bstask/api/order"
	"yandex-team.ru/bstask/domain"
	courierrepo "yandex-team.ru/bstask/repository/courier"
	orderrepo "yandex-team.ru/bstask/repository/order"
	"yandex-team.ru/bstask/service"
	timeutil "yandex-team.ru/bstask/util/time"
)

type Service struct {
	orders   orderrepo.Repository
	couriers courierrepo.Repository
	assigner domain.OrderAssigner
}

var _ order.Service = new(Service)

type ServiceParams struct {
	OrderRepository   orderrepo.Repository
	CourierRepository courierrepo.Repository
}

func NewService(p *ServiceParams) *Service {
	return &Service{
		orders:   p.OrderRepository,
		couriers: p.CourierRepository,
		assigner: &domain.DefaultOrderAssigner{},
	}
}

func (s *Service) CreateOrders(
	ctx context.Context,
	in *order.CreateOrdersIn,
) (*order.CreateOrdersOut, error) {
	createOrders := packCreateOrders(in.Orders)
	if err := validateOrders(createOrders); err != nil {
		return nil, &service.Error{
			Code:    service.ErrorCodeInvalidArgument,
			Message: fmt.Sprintf("failed to create orders: %v", err),
		}
	}

	orders, err := s.orders.CreateOrders(ctx, createOrders)
	if err != nil {
		return nil, err
	}
	return &order.CreateOrdersOut{Orders: orders}, nil
}

func (s *Service) GetOrder(
	ctx context.Context,
	in *order.GetOrderIn,
) (*order.GetOrderOut, error) {
	o, err := s.orders.GetOrder(ctx, in.OrderId)
	switch err {
	case nil:
	case orderrepo.ErrorNotFound:
		return nil, &service.Error{
			Code:    service.ErrorCodeNotFound,
			Message: fmt.Sprintf("order %d not found", in.OrderId),
		}
	default:
		return nil, err
	}
	return &order.GetOrderOut{Order: o}, nil
}

func (s *Service) ListOrders(
	ctx context.Context,
	in *order.ListOrdersIn,
) (*order.ListOrdersOut, error) {
	orders, err := s.orders.ListOrders(ctx, in.Offset, in.Limit)
	if err != nil {
		return nil, err
	}
	return &order.ListOrdersOut{Orders: orders}, nil
}

func (s *Service) CompleteOrders(
	ctx context.Context,
	in *order.CompleteOrdersIn,
) (*order.CompleteOrdersOut, error) {
	ids := make([]domain.OrderId, 0, len(in.CompleteOrders))
	completeOrders := make(map[domain.OrderId]*order.CompleteOrder)
	for _, c := range in.CompleteOrders {
		if completeOrders[c.OrderId] != nil {
			continue
		}
		ids = append(ids, c.OrderId)
		completeOrders[c.OrderId] = c
	}

	var orders []*domain.Order
	if err := s.orders.Atomic(ctx, func(repo orderrepo.Repository) (e error) {
		orders, e = repo.GetOrders(ctx, ids)
		if e != nil {
			return e
		}

		getOrders := lo.SliceToMap(orders, func(o *domain.Order) (domain.OrderId, *domain.Order) {
			return o.Id, o
		})
		for id, completeOrder := range completeOrders {
			getOrder, ok := getOrders[id]
			if !ok || getOrder == nil {
				return &service.Error{
					Code:    service.ErrorCodeInvalidArgument,
					Message: fmt.Sprintf("order %d not found", id),
				}
			}
			if getOrder.CourierId == domain.NilCourierId {
				return &service.Error{
					Code:    service.ErrorCodeInvalidArgument,
					Message: fmt.Sprintf("order %d has no assigned courier", id),
				}
			}
			if getOrder.CourierId != completeOrder.CourierId {
				return &service.Error{
					Code:    service.ErrorCodeInvalidArgument,
					Message: fmt.Sprintf("order %d has another courier assigned", id),
				}
			}
			if !getOrder.CompletedAt.IsZero() {
				continue
			}

			getOrder.CompletedAt = completeOrder.CompleteTime
		}

		return repo.UpdateOrders(ctx, orders)
	}); err != nil {
		return nil, err
	}

	return &order.CompleteOrdersOut{Orders: orders}, nil
}

func (s *Service) AssignOrders(
	ctx context.Context,
	in *order.AssignOrdersIn,
) (*order.AssignOrdersOut, error) {
	date := time.Now()
	if !in.Date.IsZero() {
		date = in.Date
	}
	date = date.UTC().Truncate(timeutil.Day)
	out := &order.AssignOrdersOut{Date: date}

	couriers, err := s.couriers.GetAvailableCouriers(ctx)
	if err != nil {
		return nil, err
	}
	if len(couriers) == 0 {
		return out, nil
	}

	if err = s.orders.Atomic(ctx, func(repo orderrepo.Repository) (e error) {
		orders, err := s.orders.GetUnassignedOrders(ctx, date)
		if err != nil {
			return err
		}
		if len(orders) == 0 {
			return nil
		}
		out.AssignOrders = unpackAssignOrders(s.assigner.Assign(orders, couriers, date))

		return repo.UpdateOrders(ctx, orders)
	}); err != nil {
		return nil, err
	}

	return out, nil
}

func (s *Service) GetAssignedOrders(
	ctx context.Context,
	in *order.GetAssignedOrdersIn,
) (*order.GetAssignedOrdersOut, error) {
	date := time.Now()
	if !in.Date.IsZero() {
		date = in.Date
	}
	date = date.UTC().Truncate(timeutil.Day)

	orders, err := s.orders.GetAssignedOrders(ctx, in.CourierId, in.Date)
	if err != nil {
		return nil, err
	}

	groupsMap := make(map[domain.CourierId]map[domain.OrderGroupId][]*domain.Order)
	for _, o := range orders {
		if _, ok := groupsMap[o.CourierId]; !ok {
			groupsMap[o.CourierId] = make(map[domain.OrderGroupId][]*domain.Order)
		}
		if _, ok := groupsMap[o.CourierId][o.GroupId]; !ok {
			groupsMap[o.CourierId][o.GroupId] = make([]*domain.Order, 0)
		}
		groupsMap[o.CourierId][o.GroupId] = append(groupsMap[o.CourierId][o.GroupId], o)
	}

	groups := make(map[domain.CourierId][]*domain.OrderGroup)
	for courierId, m := range groupsMap {
		if _, ok := groups[courierId]; !ok {
			groups[courierId] = make([]*domain.OrderGroup, 0)
		}
		for groupId, groupOrders := range m {
			groups[courierId] = append(groups[courierId], &domain.OrderGroup{
				CourierId: courierId,
				GroupId:   groupId,
				Orders:    groupOrders,
			})
		}
	}

	return &order.GetAssignedOrdersOut{
		Date:         date,
		AssignOrders: unpackAssignOrders(groups),
	}, nil
}
