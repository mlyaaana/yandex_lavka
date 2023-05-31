package gen

type GroupOrders struct {
	GroupOrderId int64      `json:"group_order_id"`
	Orders       []OrderDto `json:"orders"`
}
