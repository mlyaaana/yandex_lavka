package gen

type CouriersGroupOrders struct {
	CourierId int64         `json:"courier_id"`
	Orders    []GroupOrders `json:"orders"`
}
