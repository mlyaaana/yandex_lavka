package domain

import (
	"sort"
	"time"

	timeutil "yandex-team.ru/bstask/util/time"
)

type OrderAssigner interface {
	Assign(orders []*Order, couriers []*Courier, at time.Time) map[CourierId][]*OrderGroup
}

type DefaultOrderAssigner struct{}

// Assign algorithm:
// 1) Sort orders by delivery hours
// 2) For each order pick the best courier by this criterion:
//   2.1) Region must be the same, total count of regions must not exceed the constraint
//   2.2) Working hours must be the earliest possible and intersect with the delivery hours
//   2.3) Weight of a group must not exceed the constraint
//   2.4) Orders count must not exceed the constraint
//   2.5) Expected delivery time must not be too early or too late
// 3) If it's possible, add current order to an existing group.
//    If not, then add a new group
func (a *DefaultOrderAssigner) Assign(
	orders []*Order,
	couriers []*Courier,
	assignedAt time.Time,
) map[CourierId][]*OrderGroup {
	assignedAt = assignedAt.UTC().Truncate(timeutil.Day)
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].DeliveryHours.Less(orders[j].DeliveryHours)
	})

	couriersByRegion := make(map[Region][]*Courier)
	for _, courier := range couriers {
		couriersByWorkingHours := make([]*Courier, 0)
		for _, wh := range courier.WorkingHours {
			couriersByWorkingHours = append(couriersByWorkingHours, &Courier{
				Id:           courier.Id,
				Type:         courier.Type,
				Regions:      courier.Regions,
				WorkingHours: []timeutil.Range{wh},
			})
		}
		for _, r := range courier.Regions {
			if _, ok := couriersByRegion[r]; !ok {
				couriersByRegion[r] = make([]*Courier, 0)
			}
			couriersByRegion[r] = append(couriersByRegion[r], couriersByWorkingHours...)
		}
	}
	for _, c := range couriersByRegion {
		sort.Slice(c, func(i, j int) bool {
			return c[i].WorkingHours[0].Less(c[j].WorkingHours[0])
		})
	}

	orderGroups := make(map[CourierId][]*OrderGroup)
	for _, o := range orders {
		if o.CourierId != NilCourierId {
			continue
		}

		dh := o.DeliveryHours
		c := couriersByRegion[o.Region]
		if c == nil {
			continue
		}

		for _, courier := range c {
			wh := courier.WorkingHours[0]

			if !wh.Intersects(dh) {
				continue
			}

			maxRegionCount := CourierRegionsCount[courier.Type]
			maxWeight := CourierGroupWeight[courier.Type]
			maxGroupSize := CourierGroupSize[courier.Type]

			firstOrderTravelTime := CourierFirstOrderTravelTime[courier.Type]
			nextOrderTravelTime := CourierGroupOrdersTravelTime[courier.Type]

			groups, hasGroups := orderGroups[courier.Id]
			lastGroup := &OrderGroup{
				Range: timeutil.Range{
					End: timeutil.Max(wh.Start, dh.Start),
				},
			}
			if hasGroups {
				lastGroup = groups[len(groups)-1]
			} else {
				orderGroups[courier.Id] = make([]*OrderGroup, 0)
			}
			if lastGroup.Range.End.After(wh.End) {
				continue
			}
			_, hasRegion := lastGroup.Regions[o.Region]

			noGroups := !hasGroups
			regionsExceeded := !hasRegion && len(lastGroup.Regions) == maxRegionCount
			countExceeded := len(lastGroup.Orders) == maxGroupSize
			timeExceeded := lastGroup.Range.End.After(dh.End) ||
				lastGroup.Range.End.Add(nextOrderTravelTime).Before(dh.Start)
			weightExceeded := lastGroup.Weight.Add(o.Weight).GreaterThan(maxWeight)

			if noGroups || regionsExceeded || countExceeded || timeExceeded || weightExceeded {
				start := timeutil.Max(lastGroup.Range.End, dh.Start)
				lastGroup = &OrderGroup{
					CourierId: courier.Id,
					GroupId:   lastGroup.GroupId + 1,
					Orders:    []*Order{o},
					Regions: map[Region]struct{}{
						o.Region: {},
					},
					Range: timeutil.Range{
						Start: start,
						End:   start.Add(firstOrderTravelTime),
					},
					Weight: o.Weight,
				}
				orderGroups[courier.Id] = append(orderGroups[courier.Id], lastGroup)
			} else {
				duration := nextOrderTravelTime
				if o.Region != lastGroup.Orders[len(lastGroup.Orders)-1].Region {
					duration = firstOrderTravelTime
				}
				lastGroup.Range.End = lastGroup.Range.End.Add(duration)
				lastGroup.Regions[o.Region] = struct{}{}
				lastGroup.Weight = lastGroup.Weight.Add(o.Weight)
				lastGroup.Orders = append(lastGroup.Orders, o)
			}

			o.AssignedAt = assignedAt
			o.CourierId = courier.Id
			o.GroupId = lastGroup.GroupId
			break
		}
	}

	return orderGroups
}
