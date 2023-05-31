package rest

import (
	"github.com/samber/lo"

	"yandex-team.ru/bstask/api/courier"
	"yandex-team.ru/bstask/domain"
	"yandex-team.ru/bstask/gen"
	"yandex-team.ru/bstask/util/slice"
	"yandex-team.ru/bstask/util/time"
)

func packRegions(in []int32) []domain.Region {
	return lo.Map(in, func(r int32, _ int) domain.Region {
		return domain.Region(r)
	})
}

func unpackRegions(in []domain.Region) []int32 {
	return lo.Map(in, func(r domain.Region, _ int) int32 {
		return int32(r)
	})
}

func packWorkingHours(in []string) ([]time.Range, error) {
	return slice.Map(in, func(t string) (time.Range, error) {
		return time.ParseRange(t)
	})
}

func unpackWorkingHours(in []time.Range) []string {
	return lo.Map(in, func(t time.Range, _ int) string {
		return t.String()
	})
}

func unpackCourier(in *domain.Courier) *gen.CourierDto {
	return &gen.CourierDto{
		CourierId:    int64(in.Id),
		CourierType:  string(in.Type),
		Regions:      unpackRegions(in.Regions),
		WorkingHours: unpackWorkingHours(in.WorkingHours),
	}
}

func unpackCouriers(in []*domain.Courier) []*gen.CourierDto {
	return lo.Map(in, func(c *domain.Courier, _ int) *gen.CourierDto {
		return unpackCourier(c)
	})
}

func packCreateCouriers(
	in *gen.CreateCourierRequest,
) (*courier.CreateCouriersIn, error) {
	couriers := make([]*courier.CreateCourier, 0, len(in.Couriers))
	for _, c := range in.Couriers {
		workingHours, err := packWorkingHours(c.WorkingHours)
		if err != nil {
			return nil, err
		}

		couriers = append(couriers, &courier.CreateCourier{
			Type:         domain.CourierType(c.CourierType),
			Regions:      packRegions(c.Regions),
			WorkingHours: workingHours,
		})
	}
	return &courier.CreateCouriersIn{
		Couriers: couriers,
	}, nil
}

func unpackCourierMetaInfo(
	in *courier.GetCourierMetaInfoOut,
) *gen.GetCourierMetaInfoResponse {
	return &gen.GetCourierMetaInfoResponse{
		CourierId:    int64(in.CourierId),
		CourierType:  string(in.CourierType),
		Regions:      unpackRegions(in.Regions),
		WorkingHours: unpackWorkingHours(in.WorkingHours),
		Rating:       int32(in.Rating),
		Earnings:     int32(in.Earnings),
	}
}
