package gen

type GetCourierMetaInfoResponse struct {
	CourierId    int64    `json:"courier_id"`
	CourierType  string   `json:"courier_type"`
	Regions      []int32  `json:"regions"`
	WorkingHours []string `json:"working_hours"`
	Rating       int32    `json:"rating,omitempty"`
	Earnings     int32    `json:"earnings,omitempty"`
}
