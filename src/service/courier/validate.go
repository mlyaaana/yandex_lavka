package courier

import (
	"errors"
	"fmt"

	"yandex-team.ru/bstask/domain"
)

func validateCouriers(in []*domain.Courier) error {
	for _, c := range in {
		if !domain.CourierTypes[c.Type] {
			return fmt.Errorf("unknown courier type (val: %s)", c.Type)
		}
		if len(c.Regions) > domain.CourierRegionsCount[c.Type] {
			return errors.New("invalid regions count for courier")
		}
		for _, r := range c.Regions {
			if r <= 0 {
				return errors.New("negative courier region")
			}
		}
		for _, h := range c.WorkingHours {
			if h.Start.After(h.End) {
				return errors.New("delivery hours end before starting")
			}
		}
	}
	return nil
}
