package postgres

import (
	"context"
	"time"

	"gorm.io/gorm"

	"yandex-team.ru/bstask/database/model"
	"yandex-team.ru/bstask/domain"
	"yandex-team.ru/bstask/repository/order"
	"yandex-team.ru/bstask/util/sql"
)

type Repository struct {
	db *gorm.DB
}

var _ order.Repository = new(Repository)

type RepositoryParams struct {
	Database *gorm.DB
}

func NewRepository(p *RepositoryParams) *Repository {
	return &Repository{db: p.Database}
}

func (r *Repository) Atomic(
	ctx context.Context,
	do func(repo order.Repository) error,
) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return do(&Repository{db: tx})
	})
}

func (r *Repository) CreateOrders(
	ctx context.Context,
	orders []*domain.Order,
) ([]*domain.Order, error) {
	in := model.PackOrders(orders)
	if err := r.db.WithContext(ctx).Create(&in).Error; err != nil {
		return nil, err
	}
	return model.UnpackOrders(in), nil
}

func (r *Repository) GetOrder(
	ctx context.Context,
	id domain.OrderId,
) (*domain.Order, error) {
	in := &model.Order{}
	err := r.db.WithContext(ctx).Where(id).First(in).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, order.ErrorNotFound
	default:
		return nil, err
	}
	return model.UnpackOrder(in), nil
}

func (r *Repository) GetOrders(
	ctx context.Context,
	ids []domain.OrderId,
) ([]*domain.Order, error) {
	in := make([]*model.Order, 0, len(ids))
	if err := r.db.WithContext(ctx).Find(&in, ids).Error; err != nil {
		return nil, err
	}
	return model.UnpackOrders(in), nil
}

func (r *Repository) ListOrders(
	ctx context.Context,
	offset, limit int,
) ([]*domain.Order, error) {
	in := make([]*model.Order, 0, limit)
	err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&in).Error
	if err != nil {
		return nil, err
	}
	return model.UnpackOrders(in), nil
}

func (r *Repository) UpdateOrders(
	ctx context.Context,
	orders []*domain.Order,
) error {
	return r.db.WithContext(ctx).Save(model.PackOrders(orders)).Error
}

func (r *Repository) GetAssignedOrders(
	ctx context.Context,
	courierId domain.CourierId,
	date time.Time,
) ([]*domain.Order, error) {
	selector := &model.Order{AssignedAt: sql.NullTime(date)}
	if courierId != domain.NilCourierId {
		selector.CourierId = int64(courierId)
	}

	in := make([]*model.Order, 0)
	if err := r.db.WithContext(ctx).Where(selector).Find(&in).Error; err != nil {
		return nil, err
	}
	return model.UnpackOrders(in), nil
}

func (r *Repository) GetUnassignedOrders(
	ctx context.Context,
	date time.Time,
) ([]*domain.Order, error) {
	in := make([]*model.Order, 0)
	err := r.db.WithContext(ctx).
		Where(&model.Order{CreatedAt: date}).
		Where("assigned_at is null").
		Find(&in).Error
	if err != nil {
		return nil, err
	}
	return model.UnpackOrders(in), nil
}
