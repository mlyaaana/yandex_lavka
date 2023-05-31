package domain

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	timeutil "yandex-team.ru/bstask/util/time"
)

func TestAssign(t *testing.T) {
	today := timeutil.NowDate()
	assigner := new(DefaultOrderAssigner)

	testCases := []struct {
		couriers []*Courier
		orders   []*Order
		groups   map[CourierId][]*OrderGroup
	}{
		{
			couriers: []*Courier{
				{
					Id:           1,
					Type:         CourierTypeFoot,
					Regions:      []Region{1},
					WorkingHours: []timeutil.Range{timeutil.MustParseRange("12:00-15:00")},
				},
			},
			orders: []*Order{
				{
					Id:            1,
					Weight:        decimal.NewFromInt(5),
					Region:        1,
					DeliveryHours: timeutil.MustParseRange("12:30-13:30"),
					Cost:          5,
				},
			},
			groups: map[CourierId][]*OrderGroup{
				1: {
					{
						CourierId: 1,
						GroupId:   1,
						Orders: []*Order{
							{
								Id:            1,
								GroupId:       1,
								CourierId:     1,
								Weight:        decimal.NewFromInt(5),
								Region:        1,
								DeliveryHours: timeutil.MustParseRange("12:30-13:30"),
								Cost:          5,
								AssignedAt:    today,
							},
						},
						Regions: map[Region]struct{}{1: {}},
						Range:   timeutil.MustParseRange("12:30-12:55"),
						Weight:  decimal.NewFromInt(5),
					},
				},
			},
		},
		{
			couriers: []*Courier{
				{
					Id:           1,
					Type:         CourierTypeFoot,
					Regions:      []Region{1},
					WorkingHours: []timeutil.Range{timeutil.MustParseRange("12:00-15:00")},
				},
			},
			orders: []*Order{
				{
					Id:            1,
					Weight:        decimal.NewFromInt(5),
					Region:        1,
					DeliveryHours: timeutil.MustParseRange("12:30-13:30"),
					Cost:          5,
				},
				{
					Id:            2,
					Weight:        decimal.NewFromInt(3),
					Region:        1,
					DeliveryHours: timeutil.MustParseRange("11:00-12:20"),
					Cost:          3,
				},
			},
			groups: map[CourierId][]*OrderGroup{
				1: {
					{
						CourierId: 1,
						GroupId:   1,
						Orders: []*Order{
							{
								Id:            2,
								GroupId:       1,
								CourierId:     1,
								Weight:        decimal.NewFromInt(3),
								Region:        1,
								DeliveryHours: timeutil.MustParseRange("11:00-12:20"),
								Cost:          3,
								AssignedAt:    today,
							},
							{
								Id:            1,
								GroupId:       1,
								CourierId:     1,
								Weight:        decimal.NewFromInt(5),
								Region:        1,
								DeliveryHours: timeutil.MustParseRange("12:30-13:30"),
								Cost:          5,
								AssignedAt:    today,
							},
						},
						Regions: map[Region]struct{}{1: {}},
						Range:   timeutil.MustParseRange("12:00-12:35"),
						Weight:  decimal.NewFromInt(8),
					},
				},
			},
		},
		{
			couriers: []*Courier{
				{
					Id:           1,
					Type:         CourierTypeFoot,
					Regions:      []Region{1},
					WorkingHours: []timeutil.Range{timeutil.MustParseRange("12:00-15:00")},
				},
			},
			orders: []*Order{
				{
					Id:            1,
					Weight:        decimal.NewFromInt(5),
					Region:        1,
					DeliveryHours: timeutil.MustParseRange("12:30-13:30"),
					Cost:          5,
				},
				{
					Id:            2,
					Weight:        decimal.NewFromInt(5),
					Region:        1,
					DeliveryHours: timeutil.MustParseRange("11:00-12:20"),
					Cost:          3,
				},
				{
					Id:            3,
					Weight:        decimal.NewFromInt(8),
					Region:        1,
					DeliveryHours: timeutil.MustParseRange("13:00-14:20"),
					Cost:          4,
				},
				{
					Id:            4,
					Weight:        decimal.NewFromInt(6),
					Region:        2,
					DeliveryHours: timeutil.MustParseRange("13:00-14:00"),
					Cost:          2,
				},
			},
			groups: map[CourierId][]*OrderGroup{
				1: {
					{
						CourierId: 1,
						GroupId:   1,
						Orders: []*Order{
							{
								Id:            2,
								GroupId:       1,
								CourierId:     1,
								Weight:        decimal.NewFromInt(5),
								Region:        1,
								DeliveryHours: timeutil.MustParseRange("11:00-12:20"),
								Cost:          3,
								AssignedAt:    today,
							},
							{
								Id:            1,
								GroupId:       1,
								CourierId:     1,
								Weight:        decimal.NewFromInt(5),
								Region:        1,
								DeliveryHours: timeutil.MustParseRange("12:30-13:30"),
								Cost:          5,
								AssignedAt:    today,
							},
						},
						Regions: map[Region]struct{}{1: {}},
						Range:   timeutil.MustParseRange("12:00-12:35"),
						Weight:  decimal.NewFromInt(10),
					},
					{
						CourierId: 1,
						GroupId:   2,
						Orders: []*Order{
							{
								Id:            3,
								GroupId:       2,
								CourierId:     1,
								Weight:        decimal.NewFromInt(8),
								Region:        1,
								DeliveryHours: timeutil.MustParseRange("13:00-14:20"),
								Cost:          4,
								AssignedAt:    today,
							},
						},
						Regions: map[Region]struct{}{1: {}},
						Range:   timeutil.MustParseRange("13:00-13:25"),
						Weight:  decimal.NewFromInt(8),
					},
				},
			},
		},
		{
			orders: []*Order{
				{
					Id:            1,
					Weight:        decimal.NewFromInt(5),
					Region:        2,
					DeliveryHours: timeutil.MustParseRange("12:00-13:30"),
					Cost:          5,
				},
				{
					Id:            2,
					Weight:        decimal.NewFromInt(5),
					Region:        1,
					DeliveryHours: timeutil.MustParseRange("11:00-12:20"),
					Cost:          3,
				},
				{
					Id:            3,
					Weight:        decimal.NewFromInt(8),
					Region:        1,
					DeliveryHours: timeutil.MustParseRange("13:00-14:20"),
					Cost:          4,
				},
				{
					Id:            4,
					Weight:        decimal.NewFromInt(6),
					Region:        1,
					DeliveryHours: timeutil.MustParseRange("13:00-14:00"),
					Cost:          2,
				},
				{
					Id:            5,
					Weight:        decimal.NewFromInt(7),
					Region:        2,
					DeliveryHours: timeutil.MustParseRange("17:00-18:00"),
					Cost:          7,
				},
			},
			couriers: []*Courier{
				{
					Id:      1,
					Type:    CourierTypeBike,
					Regions: []Region{1, 2},
					WorkingHours: []timeutil.Range{
						timeutil.MustParseRange("12:00-15:00"),
						timeutil.MustParseRange("16:00-18:00"),
					},
				},
			},
			groups: map[CourierId][]*OrderGroup{
				1: {
					{
						CourierId: 1,
						GroupId:   1,
						Orders: []*Order{
							{
								Id:            2,
								CourierId:     1,
								GroupId:       1,
								Weight:        decimal.NewFromInt(5),
								Region:        1,
								DeliveryHours: timeutil.MustParseRange("11:00-12:20"),
								Cost:          3,
								AssignedAt:    today,
							},
							{
								Id:            1,
								CourierId:     1,
								GroupId:       1,
								Weight:        decimal.NewFromInt(5),
								Region:        2,
								DeliveryHours: timeutil.MustParseRange("12:00-13:30"),
								Cost:          5,
								AssignedAt:    today,
							},
						},
						Regions: map[Region]struct{}{1: {}, 2: {}},
						Range:   timeutil.MustParseRange("12:00-12:24"),
						Weight:  decimal.NewFromInt(10),
					},
					{
						CourierId: 1,
						GroupId:   2,
						Orders: []*Order{
							{
								Id:            4,
								CourierId:     1,
								GroupId:       2,
								Weight:        decimal.NewFromInt(6),
								Region:        1,
								DeliveryHours: timeutil.MustParseRange("13:00-14:00"),
								Cost:          2,
								AssignedAt:    today,
							},
							{
								Id:            3,
								CourierId:     1,
								GroupId:       2,
								Weight:        decimal.NewFromInt(8),
								Region:        1,
								DeliveryHours: timeutil.MustParseRange("13:00-14:20"),
								Cost:          4,
								AssignedAt:    today,
							},
						},
						Regions: map[Region]struct{}{1: {}},
						Range:   timeutil.MustParseRange("13:00-13:20"),
						Weight:  decimal.NewFromInt(14),
					},
					{
						CourierId: 1,
						GroupId:   3,
						Orders: []*Order{
							{
								Id:            5,
								CourierId:     1,
								GroupId:       3,
								Weight:        decimal.NewFromInt(7),
								Region:        2,
								DeliveryHours: timeutil.MustParseRange("17:00-18:00"),
								Cost:          7,
								AssignedAt:    today,
							},
						},
						Regions: map[Region]struct{}{2: {}},
						Range:   timeutil.MustParseRange("17:00-17:12"),
						Weight:  decimal.NewFromInt(7),
					},
				},
			},
		},
	}

	for _, test := range testCases {
		groups := assigner.Assign(test.orders, test.couriers, today)
		require.Equal(t, test.groups, groups)
	}
}
