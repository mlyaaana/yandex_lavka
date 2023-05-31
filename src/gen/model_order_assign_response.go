package gen

type OrderAssignResponse struct {
	Date     string                `json:"date"`
	Couriers []CouriersGroupOrders `json:"couriers"`
}
