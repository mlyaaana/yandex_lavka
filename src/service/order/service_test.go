package order

import (
	"context"
	"sort"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"yandex-team.ru/bstask/api/order"
	"yandex-team.ru/bstask/domain"
	"yandex-team.ru/bstask/repository/order/mock"
)

func TestAssignedOrders(t *testing.T) {
	controller := gomock.NewController(t)
	repository := mock.NewMockRepository(controller)
	service := &Service{
		orders: repository,
	}

	orders := []*domain.Order{
		{Id: 1, CourierId: 1, GroupId: 1},
		{Id: 2, CourierId: 1, GroupId: 1},
		{Id: 3, CourierId: 2, GroupId: 1},
		{Id: 4, CourierId: 2, GroupId: 2},
		{Id: 5, CourierId: 2, GroupId: 1},
		{Id: 6, CourierId: 1, GroupId: 2},
		{Id: 7, CourierId: 4, GroupId: 2},
		{Id: 8, CourierId: 4, GroupId: 2},
		{Id: 9, CourierId: 4, GroupId: 1},
		{Id: 10, CourierId: 5, GroupId: 1},
	}
	assignOrders := []*order.AssignOrder{
		{
			CourierId: 1,
			Groups: []*order.Group{
				{
					Id:     1,
					Orders: []*domain.Order{orders[0], orders[1]},
				},
				{
					Id:     2,
					Orders: []*domain.Order{orders[5]},
				},
			},
		},
		{
			CourierId: 2,
			Groups: []*order.Group{
				{
					Id:     1,
					Orders: []*domain.Order{orders[2], orders[4]},
				},
				{
					Id:     2,
					Orders: []*domain.Order{orders[3]},
				},
			},
		},
		{
			CourierId: 4,
			Groups: []*order.Group{
				{
					Id:     1,
					Orders: []*domain.Order{orders[8]},
				},
				{
					Id:     2,
					Orders: []*domain.Order{orders[6], orders[7]},
				},
			},
		},
		{
			CourierId: 5,
			Groups: []*order.Group{
				{
					Id:     1,
					Orders: []*domain.Order{orders[9]},
				},
			},
		},
	}
	repository.EXPECT().GetAssignedOrders(gomock.Any(), gomock.Any(), gomock.Any()).Return(orders, nil)
	out, err := service.GetAssignedOrders(context.Background(), &order.GetAssignedOrdersIn{})
	require.NoError(t, err)
	require.Len(t, out.AssignOrders, len(assignOrders))
	sort.Slice(out.AssignOrders, func(i, j int) bool {
		return out.AssignOrders[i].CourierId < out.AssignOrders[j].CourierId
	})

	for i := range assignOrders {
		expected, actual := assignOrders[i], out.AssignOrders[i]
		require.Equal(t, expected.CourierId, actual.CourierId)
		sort.Slice(actual.Groups, func(l, r int) bool {
			return actual.Groups[l].Id < actual.Groups[r].Id
		})
		require.Equal(t, expected.Groups, actual.Groups)
	}
}
