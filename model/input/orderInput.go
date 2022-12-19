package input

type CreateOrderInput struct {
	DeliverySourceAddress string `json:"source_address" binding:"required"`
	DeliveryDestAddress   string `json:"destination_address" binding:"required"`
	Items                 uint   `json:"item_id" binding:"required"`
	Quantity              int    `json:"quantity" binding:"required"`
}
