package postgres

import (
	"context"

	"gorm.io/gorm"

	"yandex-team.ru/bstask/database/model"
	"yandex-team.ru/bstask/domain"
	"yandex-team.ru/bstask/repository/courier"
	timeutil "yandex-team.ru/bstask/util/time"
)

type Repository struct {
	db *gorm.DB
}

var _ courier.Repository = new(Repository)

type RepositoryParams struct {
	Database *gorm.DB
}

func NewRepository(p *RepositoryParams) *Repository {
	return &Repository{db: p.Database}
}

func (r *Repository) CreateCouriers(
	ctx context.Context,
	couriers []*domain.Courier,
) ([]*domain.Courier, error) {
	in := model.PackCouriers(couriers)
	if err := r.db.WithContext(ctx).Create(&in).Error; err != nil {
		return nil, err
	}
	return model.UnpackCouriers(in), nil
}

func (r *Repository) GetCourier(
	ctx context.Context,
	id domain.CourierId,
) (*domain.Courier, error) {
	in := &model.Courier{}
	err := r.db.WithContext(ctx).Where(id).First(in).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, courier.ErrorNotFound
	default:
		return nil, err
	}
	return model.UnpackCourier(in), nil
}

func (r *Repository) GetAvailableCouriers(
	ctx context.Context,
) ([]*domain.Courier, error) {
	in := make([]*model.Courier, 0)
	if err := r.db.WithContext(ctx).Find(&in).Error; err != nil {
		return nil, err
	}
	return model.UnpackCouriers(in), nil
}

func (r *Repository) GetCourierStats(
	ctx context.Context,
	id domain.CourierId,
	dates timeutil.Range,
) (*domain.CourierStats, error) {
	in := &model.Courier{}
	err := r.db.WithContext(ctx).Where(id).First(in).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		return nil, courier.ErrorNotFound
	default:
		return nil, err
	}

	var stats struct{ Sum, Count int }
	query := "select sum(cost), count(*) from orders where courier_id = ? and completed_at between ? and ?"
	if err = r.db.WithContext(ctx).Raw(query, id, dates.Start, dates.End).Scan(&stats).Error; err != nil {
		return nil, err
	}

	return &domain.CourierStats{
		Courier: model.UnpackCourier(in),
		Orders:  stats.Count,
		Income:  stats.Sum,
	}, nil
}

func (r *Repository) ListCouriers(
	ctx context.Context,
	limit, offset int,
) ([]*domain.Courier, error) {
	in := make([]*model.Courier, 0, limit)
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&in).Error
	if err != nil {
		return nil, err
	}
	return model.UnpackCouriers(in), nil
}
