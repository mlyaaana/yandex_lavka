package courier

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"yandex-team.ru/bstask/api/courier"
	"yandex-team.ru/bstask/domain"
	"yandex-team.ru/bstask/repository/courier/mock"
	timeutil "yandex-team.ru/bstask/util/time"
)

func TestCourierMetaInfo(t *testing.T) {
	controller := gomock.NewController(t)
	repository := mock.NewMockRepository(controller)
	service := &Service{
		couriers: repository,
	}

	testCases := []struct {
		orders   int
		income   int
		days     int
		typ      domain.CourierType
		earnings int
		rating   int
	}{
		{
			orders:   15,
			income:   250,
			days:     1,
			typ:      domain.CourierTypeFoot,
			earnings: 500,
			rating:   2,
		},
		{
			orders:   7,
			income:   100,
			days:     1,
			typ:      domain.CourierTypeFoot,
			earnings: 200,
			rating:   1,
		},
		{
			orders:   4,
			income:   100,
			days:     1,
			typ:      domain.CourierTypeFoot,
			earnings: 200,
			rating:   1,
		},
		{
			orders:   3,
			income:   100,
			days:     1,
			typ:      domain.CourierTypeFoot,
			earnings: 200,
			rating:   0,
		},
		{
			orders:   8,
			income:   100,
			days:     1,
			typ:      domain.CourierTypeFoot,
			earnings: 200,
			rating:   1,
		},
		{
			orders:   8,
			income:   100,
			days:     2,
			typ:      domain.CourierTypeFoot,
			earnings: 200,
			rating:   1,
		},
		{
			orders:   15,
			income:   250,
			days:     1,
			typ:      domain.CourierTypeBike,
			earnings: 750,
			rating:   1,
		},
		{
			orders:   18,
			income:   250,
			days:     1,
			typ:      domain.CourierTypeBike,
			earnings: 750,
			rating:   2,
		},
		{
			orders:   18,
			income:   250,
			days:     1,
			typ:      domain.CourierTypeBike,
			earnings: 750,
			rating:   2,
		},
		{
			orders:   18,
			income:   250,
			days:     1,
			typ:      domain.CourierTypeAuto,
			earnings: 1000,
			rating:   1,
		},
		{
			orders:   12,
			income:   250,
			days:     1,
			typ:      domain.CourierTypeAuto,
			earnings: 1000,
			rating:   1,
		},
		{
			orders:   58,
			income:   250,
			days:     4,
			typ:      domain.CourierTypeAuto,
			earnings: 1000,
			rating:   1,
		},
	}
	for _, test := range testCases {
		repository.EXPECT().GetCourierStats(gomock.Any(), gomock.Any(), gomock.Any()).Return(&domain.CourierStats{
			Courier: &domain.Courier{
				Id:   1,
				Type: test.typ,
			},
			Orders: test.orders,
			Income: test.income,
		}, nil)
		meta, err := service.GetCourierMetaInfo(context.Background(), &courier.GetCourierMetaInfoIn{
			StartDate: timeutil.NowDate(),
			EndDate:   timeutil.NowDate().Add(time.Duration(test.days) * timeutil.Day),
		})
		require.NoError(t, err)
		require.Equal(t, domain.CourierId(1), meta.CourierId)
		require.Equal(t, test.typ, meta.CourierType)
		require.Equal(t, test.earnings, meta.Earnings)
		require.Equal(t, test.rating, meta.Rating)
	}
}
