package order

import (
	"errors"

	"github.com/shopspring/decimal"

	"yandex-team.ru/bstask/domain"
)

func validateOrders(in []*domain.Order) error {
	for _, o := range in {
		if o.Weight.LessThan(decimal.Zero) {
			return errors.New("negative order weight")
		}
		if o.Region <= 0 {
			return errors.New("negative order region")
		}
		if o.Cost < 0 {
			return errors.New("negative order cost")
		}
		if o.DeliveryHours.Start.After(o.DeliveryHours.End) {
			return errors.New("delivery hours end before starting")
		}
	}
	return nil
}
