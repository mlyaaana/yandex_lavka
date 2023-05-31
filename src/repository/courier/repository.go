package courier

import (
	"context"
	"errors"

	"yandex-team.ru/bstask/domain"
	timeutil "yandex-team.ru/bstask/util/time"
)

var ErrorNotFound = errors.New("courier not found")

type Repository interface {
	CreateCouriers(
		ctx context.Context,
		couriers []*domain.Courier,
	) ([]*domain.Courier, error)
	GetCourier(
		ctx context.Context,
		id domain.CourierId,
	) (*domain.Courier, error)
	GetAvailableCouriers(
		ctx context.Context,
	) ([]*domain.Courier, error)
	GetCourierStats(
		ctx context.Context,
		id domain.CourierId,
		dates timeutil.Range,
	) (*domain.CourierStats, error)
	ListCouriers(
		ctx context.Context,
		limit, offset int,
	) ([]*domain.Courier, error)
}
