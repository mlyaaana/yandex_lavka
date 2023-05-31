package courier

import (
	"context"
	"time"

	"yandex-team.ru/bstask/domain"
	timeutil "yandex-team.ru/bstask/util/time"
)

type Service interface {
	CreateCouriers(
		ctx context.Context,
		in *CreateCouriersIn,
	) (*CreateCouriersOut, error)
	GetCourier(
		ctx context.Context,
		in *GetCourierIn,
	) (*GetCourierOut, error)
	ListCouriers(
		ctx context.Context,
		in *ListCouriersIn,
	) (*ListCouriersOut, error)
	GetCourierMetaInfo(
		ctx context.Context,
		in *GetCourierMetaInfoIn,
	) (*GetCourierMetaInfoOut, error)
}

type CreateCourier struct {
	Type         domain.CourierType
	Regions      []domain.Region
	WorkingHours []timeutil.Range
}

type CreateCouriersIn struct {
	Couriers []*CreateCourier
}

type CreateCouriersOut struct {
	Couriers []*domain.Courier
}

type GetCourierIn struct {
	CourierId domain.CourierId
}

type GetCourierOut struct {
	Courier *domain.Courier
}

type ListCouriersIn struct {
	Limit  int
	Offset int
}

type ListCouriersOut struct {
	Couriers []*domain.Courier
}

type GetCourierMetaInfoIn struct {
	CourierId domain.CourierId
	StartDate time.Time
	EndDate   time.Time
}

type GetCourierMetaInfoOut struct {
	CourierId    domain.CourierId
	CourierType  domain.CourierType
	Regions      []domain.Region
	WorkingHours []timeutil.Range
	Rating       int
	Earnings     int
}
