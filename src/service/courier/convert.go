package courier

import (
	"github.com/samber/lo"

	"yandex-team.ru/bstask/api/courier"
	"yandex-team.ru/bstask/domain"
)

func packCreateCourier(in *courier.CreateCourier) *domain.Courier {
	return &domain.Courier{
		Type:         in.Type,
		Regions:      in.Regions,
		WorkingHours: in.WorkingHours,
	}
}

func packCreateCouriers(in []*courier.CreateCourier) []*domain.Courier {
	return lo.Map(in, func(o *courier.CreateCourier, _ int) *domain.Courier {
		return packCreateCourier(o)
	})
}
