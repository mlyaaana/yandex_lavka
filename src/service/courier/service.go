package courier

import (
	"context"
	"fmt"
	"math"

	"yandex-team.ru/bstask/api/courier"
	"yandex-team.ru/bstask/domain"
	repo "yandex-team.ru/bstask/repository/courier"
	"yandex-team.ru/bstask/service"
	timeutil "yandex-team.ru/bstask/util/time"
)

type Service struct {
	couriers repo.Repository
}

var _ courier.Service = new(Service)

type ServiceParams struct {
	CourierRepository repo.Repository
}

func NewService(p *ServiceParams) *Service {
	return &Service{couriers: p.CourierRepository}
}

func (s *Service) ListCouriers(
	ctx context.Context,
	in *courier.ListCouriersIn,
) (*courier.ListCouriersOut, error) {
	couriers, err := s.couriers.ListCouriers(ctx, in.Offset, in.Limit)
	if err != nil {
		return nil, err
	}
	return &courier.ListCouriersOut{Couriers: couriers}, nil
}

func (s *Service) CreateCouriers(
	ctx context.Context,
	in *courier.CreateCouriersIn,
) (*courier.CreateCouriersOut, error) {
	createCouriers := packCreateCouriers(in.Couriers)
	if err := validateCouriers(createCouriers); err != nil {
		return nil, &service.Error{
			Code:    service.ErrorCodeInvalidArgument,
			Message: fmt.Sprintf("failed to create couriers: %v", err),
		}
	}

	couriers, err := s.couriers.CreateCouriers(ctx, createCouriers)
	if err != nil {
		return nil, err
	}
	return &courier.CreateCouriersOut{Couriers: couriers}, nil
}

func (s *Service) GetCourier(
	ctx context.Context,
	in *courier.GetCourierIn,
) (*courier.GetCourierOut, error) {
	c, err := s.couriers.GetCourier(ctx, in.CourierId)
	switch err {
	case nil:
	case repo.ErrorNotFound:
		return nil, &service.Error{
			Code:    service.ErrorCodeNotFound,
			Message: fmt.Sprintf("courier %d not found", in.CourierId),
		}
	default:
		return nil, err
	}
	return &courier.GetCourierOut{Courier: c}, nil
}

func (s *Service) GetCourierMetaInfo(
	ctx context.Context,
	in *courier.GetCourierMetaInfoIn,
) (*courier.GetCourierMetaInfoOut, error) {
	c, err := s.couriers.GetCourierStats(ctx, in.CourierId, timeutil.Range{
		Start: in.StartDate,
		End:   in.EndDate,
	})
	switch err {
	case nil:
	case repo.ErrorNotFound:
		return nil, &service.Error{
			Code:    service.ErrorCodeNotFound,
			Message: fmt.Sprintf("courier %d not found", in.CourierId),
		}
	default:
		return nil, err
	}

	hours := in.EndDate.Sub(in.StartDate).Hours()
	ratingCoefficient := domain.CourierRatingCoefficients[c.Courier.Type]
	rating := math.Round(float64(ratingCoefficient) * float64(c.Orders) / hours)
	earnings := c.Income * domain.CourierEarningsCoefficients[c.Courier.Type]

	return &courier.GetCourierMetaInfoOut{
		CourierId:    c.Courier.Id,
		CourierType:  c.Courier.Type,
		Regions:      c.Courier.Regions,
		WorkingHours: c.Courier.WorkingHours,
		Rating:       int(rating),
		Earnings:     earnings,
	}, err
}
