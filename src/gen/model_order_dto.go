package gen

import (
	"time"
)

type OrderDto struct {
	OrderId       int64      `json:"order_id"`
	Weight        float32    `json:"weight"`
	Regions       int32      `json:"regions"`
	DeliveryHours []string   `json:"delivery_hours"`
	Cost          int32      `json:"cost"`
	CompletedTime *time.Time `json:"completed_time,omitempty"`
}
