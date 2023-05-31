package order

import (
	"context"
	"errors"
	"time"

	"yandex-team.ru/bstask/domain"
)

var ErrorNotFound = errors.New("order not found")

type Repository interface {
	Atomic(
		ctx context.Context,
		do func(repo Repository) error,
	) error
	CreateOrders(
		ctx context.Context,
		orders []*domain.Order,
	) ([]*domain.Order, error)
	GetOrder(
		ctx context.Context,
		id domain.OrderId,
	) (*domain.Order, error)
	GetOrders(
		ctx context.Context,
		ids []domain.OrderId,
	) ([]*domain.Order, error)
	ListOrders(
		ctx context.Context,
		offset, limit int,
	) ([]*domain.Order, error)
	UpdateOrders(
		ctx context.Context,
		orders []*domain.Order,
	) error
	GetAssignedOrders(
		ctx context.Context,
		courierId domain.CourierId,
		date time.Time,
	) ([]*domain.Order, error)
	GetUnassignedOrders(
		ctx context.Context,
		date time.Time,
	) ([]*domain.Order, error)
}
